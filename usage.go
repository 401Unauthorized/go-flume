package main

import (
	"context"
	"fmt"
	"net/url"
)

type QueryUsageRequestBody struct {
	RequestID       string   `json:"request_id"`
	Bucket          string   `json:"bucket"`
	SinceDatetime   string   `json:"since_datetime,omitempty"`
	UntilDatetime   string   `json:"until_datetime,omitempty"`
	GroupMultiplier string   `json:"group_multiplier,omitempty"`
	Operation       string   `json:"operation,omitempty"`
	SortDirection   string   `json:"sort_direction,omitempty"`
	Units           string   `json:"units,omitempty"`
	Types           []string `json:"types,omitempty"`
}

type UsageQuery struct {
	Value    int    `json:"value"`
	Datetime string `json:"datetime"`
}

type QueryUsageResponse struct {
	APIResponseEnvelope
	Data []UsageQuery `json:"data"`
}

func (c *Client) QueryUsage(ctx context.Context, deviceID string, data QueryUsageRequestBody) (*QueryUsageResponse, error) {
	if deviceID == "" {
		return nil, fmt.Errorf("deviceID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/devices/%s/query", c.BaseURL, c.JWT.UserID, deviceID)
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	var resp QueryUsageResponse
	if err := c.apiRequest(ctx, "POST", u, data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
