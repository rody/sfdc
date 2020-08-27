package sfdc

import (
	"context"
	"net/http"
)

type SObjectsService struct {
	Service
}

type DescribeGlobalResponse struct {
	Encoding     string     `json:"encoding"`
	MaxBatchSize int        `json:"maxBatchSize"`
	SObjects     []Describe `json:"sobjects"`
}

type Describe struct {
	Activateable        bool   `json:"activateable"`
	Custom              bool   `json:"custom"`
	CustomSetting       bool   `json:"customSetting"`
	Createable          bool   `json:"createable"`
	Deletable           bool   `json:"deletable"`
	DeprecatedAndHidden bool   `json:"deprecatedAndHidden"`
	FeedEnabled         bool   `json:"feedEnabled"`
	KeyPrefix           string `json:"keyPrefix"`
	Label               string `json:"label"`
	LabelPlural         string `json:"labelPlural"`
	Layoutable          bool   `json:"layoutable"`
	Mergeable           bool   `json:"mergeable"`
	MRUEnabled          bool   `json:"mruEnabled"`
	Name                string `json:"name"`
	Queryable           bool   `json:"queryable"`
	Replicateable       bool   `json:"replicateable"`
	Searchable          bool   `json:"searchable"`
	Triggerable         bool   `json:"triggerable"`
	Undeletable         bool   `json:"undeletable"`
	Updateable          bool   `json:"updateable"`

	URLs map[string]string `json:"urls"`
}

// DescribeGlobal returns the description of the SObjects on the org
func (s *SObjectsService) DescribeGlobal(ctx context.Context) (*DescribeGlobalResponse, error) {
	req, err := s.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}

	var result DescribeGlobalResponse
	if err = s.client.Do(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
