package sfdc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// QueryService represents the Query api in Salesforce
type QueryService struct {
	Service
}

// QueryResult is returned by the API for a SOQL request.
// it provides 2 convenience methods:
//  - Next() will retrieve the next page of result (if any)
//  - UnmarshallRecords() will perform the JSON unmarshaling of the record result
type QueryResult struct {
	client         *Client
	Done           bool            `json:"done"`
	TotalSize      int64           `json:"totalSize"`
	Records        json.RawMessage `json:"records"`
	NextRecordsURL string          `json:"nextRecordsUrl"`
}

// ExplainResult is returned by the API after an Explain() call
type ExplainResult struct {
	Plans []QueryPlan `json:"plans"`
}

// QueryPlan holds the plan detaisl for a query (see Explain())
type QueryPlan struct {
	Cardinality          int            `json:"cardinality"`
	Fields               []string       `json:"fields"`
	LeadingOperationType string         `json:"leadingOperationType"`
	Notes                []FeedbackNote `json:"notes"`
	RelativeCost         float32        `json:"relativeCost"`
	SObjectCardinality   int            `json:"sobjectCardinality"`
	SObjectType          string         `json:"sobjectType"`
}

// FeedbackNote is a note in a QueryPlan
type FeedbackNote struct {
	Description   string   `json:"description"`
	Fields        []string `json:"fields"`
	TableEnumOrID string   `json:"tableEnumOrId"`
}

// Query executes a soql query on Salesforce
func (qs *QueryService) Query(ctx context.Context, soql string) (*QueryResult, error) {
	req, err := qs.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", soql)
	req.URL.RawQuery = q.Encode()

	var result QueryResult
	err = qs.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	result.client = qs.client
	return &result, nil
}

func (qs *QueryService) Explain(ctx context.Context, soqlOrID string) (*ExplainResult, error) {
	req, err := qs.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("explain", soqlOrID)
	req.URL.RawQuery = q.Encode()

	var result ExplainResult
	err = qs.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (q *QueryResult) UnmarshalRecords(r interface{}) error {
	return json.Unmarshal(q.Records, r)
}

func (q *QueryResult) Next(ctx context.Context) (*QueryResult, error) {
	if q.NextRecordsURL == "" {
		return nil, errors.New("no more results")
	}

	rel := &url.URL{Path: q.NextRecordsURL}
	u := q.client.InstanceURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	var result QueryResult
	err = q.client.Do(ctx, req, &result)

	if err != nil {
		return nil, err
	}

	result.client = q.client
	return &result, nil
}
