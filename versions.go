package sfdc

import (
	"context"
	"net/http"
	"net/url"
)

// Version describe an API version available
// on Salesforce
type Version struct {
	Version string `json:"version"`
	Label   string `json:"label"`
	URL     string `json:"url"`
}

// Versions calls the Salesforce API and returns all
// the API versions available on the org
func (c *Client) Versions(ctx context.Context) ([]Version, error) {
	u := c.InstanceURL.ResolveReference(&url.URL{Path: "/services/data"})

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	var versions []Version
	err = c.Do(ctx, req, &versions)
	return versions, err
}
