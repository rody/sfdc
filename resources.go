package sfdc

import (
	"context"
	"fmt"
	"net/http"
)

type RestResources map[string]string

// Services returns all the availble services on Salesforce
func (c *Client) Resources(ctx context.Context) (*RestResources, error) {

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("/services/data/v%s", c.version), nil)
	if err != nil {
		return nil, err
	}

	var resources RestResources
	if err = c.Do(ctx, req, &resources); err != nil {
		return nil, err
	}

	return &resources, nil
}
