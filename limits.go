package sfdc

import (
	"context"
	"net/http"
)

// LimitsService interacts with the Limits API
type LimitsService struct {
	Service
}

// Limits represents the limits response
type Limits map[string]LimitMetric

// LimitMetric contains the Max and Remaining values for a limit
type LimitMetric struct {
	Max       int `json:"Max"`
	Remaining int `json:"Remaining"`
}

// Get retrieves all the limit values for the current org
func (l *LimitsService) Get(ctx context.Context) (*Limits, error) {
	req, err := l.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}

	var limits Limits
	if err = l.client.Do(ctx, req, &limits); err != nil {
		return nil, err
	}

	return &limits, err
}
