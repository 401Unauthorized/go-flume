package main

import (
	"context"
	"fmt"
	"net/url"
)

type UsageAlertQuery struct {
	RequestID     string   `json:"request_id"`
	SinceDatetime string   `json:"since_datetime"`
	UntilDatetime string   `json:"until_datetime"`
	TZ            string   `json:"tz"`
	Bucket        string   `json:"bucket"`
	DeviceID      []string `json:"device_id"`
}

type UsageAlert struct {
	ID                int             `json:"id"`
	DeviceID          string          `json:"device_id"`
	TriggeredDatetime string          `json:"triggered_datetime"`
	FlumeLeak         bool            `json:"flume_leak"`
	Query             UsageAlertQuery `json:"query"`
	EventRuleName     string          `json:"event_rule_name"`
}

type UsageAlertsResponse struct {
	APIResponseEnvelope
	Data []UsageAlert `json:"data"`
}

type GetUsageAlertsParams struct {
	Limit         *int32  // How many usage alerts to return (Defaults to 50)
	Offset        *int32  // Offset of usage alerts to return (Defaults to 0)
	SortField     *string // Which field to sort usage alerts on (Defaults to id)
	SortDirection *string // Which direction to sort usage alerts on (Defaults to ASC)
	DeviceID      *string // Return usage alerts for this device_id
	FlumeLeak     *bool   // Returns usage alerts determined to be leak s
}

func (c *Client) GetUsageAlerts(ctx context.Context, params *GetUsageAlertsParams) (*UsageAlertsResponse, error) {
	req := fmt.Sprintf("%s/users/%d/usage-alerts", c.BaseURL, c.JWT.UserID)
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
		if params.DeviceID != nil {
			query.Set("device_id", *params.DeviceID)
		}
		if params.FlumeLeak != nil {
			query.Set("flume_leak", fmt.Sprintf("%t", *params.FlumeLeak))
		}
	}
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var resp UsageAlertsResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
