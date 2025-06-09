package goflume

import (
	"context"
	"fmt"
	"net/url"
)

type Subscription struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	AlertType         string `json:"alert_type"`
	AlertInfo         string `json:"alert_info"`
	DeviceID          string `json:"device_id"`
	NotificationTypes int    `json:"notification_types"`
	CreatedDatetime   string `json:"created_datetime"`
	UpdatedDatetime   string `json:"updated_datetime"`
}

type SubscriptionsResponse struct {
	APIResponseEnvelopePagination
	Data []Subscription `json:"data"`
}

type GetSubscriptionsParams struct {
	Limit             *int32  // How many subscriptions to return (Defaults to 50)
	Offset            *int32  // Offset of subscriptions to return (Defaults to 0)
	SortField         *string // Which field to sort the subscriptions on (Defaults to id)
	SortDirection     *string // The direction to sort the subscriptions on (Defaults to ASC)
	AlertType         *string // Only return subscriptions with this alert type
	NotificationTypes *int32  // Only return subscriptions that subscribe to the exact bitmask of notification types
	NotificationType  *int32  // Return all subscriptions that contain the bit
	DeviceID          *string // Only return subscriptions that are for this device
	DeviceType        *int32  // Only return subscriptions for devices of this type
	LocationID        *int32  // Only return subscriptions that are associated with a device at this location
}

func (c *Client) GetSubscriptions(ctx context.Context, params *GetSubscriptionsParams) (*SubscriptionsResponse, error) {
	req := fmt.Sprintf("%s/users/%d/subscriptions", c.BaseURL, c.JWT.UserID)
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
		if params.AlertType != nil {
			query.Set("alert_type", *params.AlertType)
		}
		if params.NotificationTypes != nil {
			query.Set("notification_types", fmt.Sprintf("%d", *params.NotificationTypes))
		}
		if params.NotificationType != nil {
			query.Set("notification_type", fmt.Sprintf("%d", *params.NotificationType))
		}
		if params.DeviceID != nil {
			query.Set("device_id", *params.DeviceID)
		}
		if params.DeviceType != nil {
			query.Set("device_type", fmt.Sprintf("%d", *params.DeviceType))
		}
		if params.LocationID != nil {
			query.Set("location_id", fmt.Sprintf("%d", *params.LocationID))
		}
	}
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var resp SubscriptionsResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SubscriptionResponse struct {
	APIResponseEnvelope
	Data Subscription `json:"data"`
}

func (c *Client) GetSubscription(ctx context.Context, subscriptionID string) (*SubscriptionResponse, error) {
	if subscriptionID == "" {
		return nil, fmt.Errorf("subscriptionID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/subscriptions/%s", c.BaseURL, c.JWT.UserID, subscriptionID)
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	var resp SubscriptionResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
