package goflume

import (
	"context"
	"fmt"
	"net/url"
)

type EventRule struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Active           bool    `json:"active"`
	FlowRate         float64 `json:"flow_rate"`
	Duration         int     `json:"duration"`
	NotifyEvery      int     `json:"notify_every"`
	NotificationType string  `json:"notification_type"`
}

type EventRulesResponse struct {
	APIResponseEnvelope
	Data []EventRule `json:"data"`
}

type GetEventRulesParams struct {
	Limit         *int32  // Max number of event rules to return (Defaults to 50)
	Offset        *int32  // Offset of event rules to return (Defaults to 0)
	SortField     *string // Field to sort event rules on (Defaults to id)
	SortDirection *string // Sort direction (Defaults to ASC)
}

func (c *Client) GetEventRules(ctx context.Context, deviceID string, params *GetEventRulesParams) (*EventRulesResponse, error) {
	if deviceID == "" {
		return nil, fmt.Errorf("deviceID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/devices/%s/event_rules", c.BaseURL, c.JWT.UserID, deviceID)
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
	var resp EventRulesResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UsageAlertRule struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Enabled   bool    `json:"enabled"`
	Threshold float64 `json:"threshold"`
	Unit      string  `json:"unit"`
}

type UsageAlertRulesResponse struct {
	APIResponseEnvelope
	Data []UsageAlertRule `json:"data"`
}

type GetUsageAlertRulesParams struct {
	Limit         *int32  // Max number of event rules to return (Defaults to 50)
	Offset        *int32  // Offset of event rules to return (Defaults to 0)
	SortField     *string // Field to sort event rules on (Defaults to id)
	SortDirection *string // Sort direction (Defaults to ASC)
}

func (c *Client) GetUsageAlertRules(ctx context.Context, deviceID string, params *GetUsageAlertRulesParams) (*UsageAlertRulesResponse, error) {
	if deviceID == "" {
		return nil, fmt.Errorf("deviceID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/devices/%s/usage_alert_rules", c.BaseURL, c.JWT.UserID, deviceID)
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
	var resp UsageAlertRulesResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UsageAlertRuleResponse struct {
	APIResponseEnvelope
	Data []UsageAlertRule `json:"data"`
}

func (c *Client) GetUsageAlertRule(ctx context.Context, deviceID, ruleID string) (*UsageAlertRuleResponse, error) {
	if deviceID == "" {
		return nil, fmt.Errorf("deviceID cannot be empty")
	}
	if ruleID == "" {
		return nil, fmt.Errorf("ruleID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/devices/%s/usage_alert_rules/%s", c.BaseURL, c.JWT.UserID, deviceID, ruleID)
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	var resp UsageAlertRuleResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
