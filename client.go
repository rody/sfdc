package sfdc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultVersion = "48.0"
)

type Client struct {
	client      *http.Client
	version     string
	InstanceURL *url.URL
	UserAgent   string

	Query *QueryService
}

type ClientOption func(*Client) error

// NewClient creates an sfdc client. If given, httpClient
// is used for all HTTP operations, if not, the default http
// client will be used instead (http.DefaultClient)
//
// To use a OAuth2 flow with the client, you need to pass an
// OAuth2-aware client
//
// instanceURL is the URL containing the URL of youe salesforce org.
//
// example:
// TODO
//
func NewClient(httpClient *http.Client, instanceURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(instanceURL)
	if err != nil {
		return nil, err
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client:      httpClient,
		version:     defaultVersion,
		InstanceURL: u,
	}

	for _, opt := range opts {
		if err = opt(c); err != nil {
			return c, err
		}
	}

	c.Query = &QueryService{Service{client: c, basePath: fmt.Sprintf("/services/data/v%s/query/", c.version)}}

	return c, nil
}

func WithVersion(version string) ClientOption {
	return func(c *Client) error {
		// TODO: add checks and handle v prefix
		c.version = version
		return nil
	}
}

// Version returns the API version set for this client
func (c *Client) Version() string {
	return c.version
}

// Do sends an HTTP request and tries to unmarshall the response in the given
// interface.
// if v is an io.Writer, the content of the Body os written to the writer, otherwise
// the body is unmarshal (using the json package) into the given data structure
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)

	if c.UserAgent != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		// Drain and close the body to let the Transport reuse the connection
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	err = checkResponse(resp)
	if err != nil {
		return err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.InstanceURL.ResolveReference(rel)

	var content io.Reader
	if body != nil {
		if r, ok := body.(io.Reader); ok {
			content = r
		} else {
			data, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			content = bytes.NewReader(data)
		}
	}

	return http.NewRequest(method, u.String(), content)
}

func checkResponse(res *http.Response) error {
	if c := res.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := ErrorResponse{Response: res}
	data, err := ioutil.ReadAll(res.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &errorResponse.Errors)
		if err != nil {
			return err
		}
	}
	return errorResponse
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Errors   []struct {
		Message   string `json:"message"`
		Errorcode string `json:"errorCode"`
	}
}

func (r ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Errors[0].Message)
}
