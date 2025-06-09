package main

import (
	"context"
	"fmt"
	"net/url"
)

type Budget struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Value      int    `json:"value"`
	Thresholds []int  `json:"thresholds"`
	Actual     int    `json:"actual"`
}

type BudgetsResponse struct {
	APIResponseEnvelope
	Data []Budget `json:"data"`
}

type GetBudgetsParams struct {
	Limit         *int32  // Max number of budgets to return (Defaults to 50)
	Offset        *int32  // Offset of budgets to return (Defaults to 0)
	SortField     *string // Field to sort budgets on (Defaults to id)
	SortDirection *string // Which direction to sort budgets on (Defaults to ASC)
}

func (c *Client) GetBudgets(ctx context.Context, deviceID string, params *GetBudgetsParams) (*BudgetsResponse, error) {
	if deviceID == "" {
		return nil, fmt.Errorf("deviceID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/devices/%s/budgets", c.BaseURL, c.JWT.UserID, deviceID)
	query := url.Values{}
	if params != nil {
		if params.Limit != nil {
			query.Set("limit", fmt.Sprintf("%d", *params.Limit))
		}
		if params.Offset != nil {
			query.Set("offset", fmt.Sprintf("%d", *params.Offset))
		}
		if params.SortField != nil {
			query.Set("sort_field", *params.SortField)
		}
		if params.SortDirection != nil {
			query.Set("sort_direction", *params.SortDirection)
		}
	}
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var resp BudgetsResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
