package main

import (
	"context"
	"fmt"
	"net/url"
)

type Flow struct {
	Active   bool    `json:"active"`
	GPM      float64 `json:"gpm"`
	Datetime string  `json:"datetime"`
}
type FlowResponse struct {
	APIResponseEnvelope
	Data []Flow `json:"data"`
}

func (c *Client) GetCurrentFlow(ctx context.Context, deviceID string) (*FlowResponse, error) {
	if deviceID == "" {
		return nil, fmt.Errorf("deviceID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/devices/%s/query/active", c.BaseURL, c.JWT.UserID, deviceID)
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	var resp FlowResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
