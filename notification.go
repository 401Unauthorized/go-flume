package goflume

import (
	"context"
	"fmt"
	"net/url"
)

type Notification struct {
	ID              int    `json:"id"`
	DeviceID        string `json:"device_id"`
	UserID          int    `json:"user_id"`
	Type            int    `json:"type"`
	Message         string `json:"message"`
	CreatedDatetime string `json:"created_datetime"`
	Title           string `json:"title"`
	Read            bool   `json:"read"`
	Extra           string `json:"extra"`
}

type NotificationsResponse struct {
	APIResponseEnvelope
	Data []Notification `json:"data"`
}

type GetNotificationsParams struct {
	Limit         *int32  // How many notifications to return (Defaults to 50)
	Offset        *int32  // Offset of notifications to return (Defaults to 0)
	SortField     *string // Which field to sort notifications on (Defaults to created_datetime)
	SortDirection *string // Which direction to sort notifications on (Defaults to ASC)
	DeviceID      *string // Return notifications sent from a device with this device_id
	LocationID    *int32  // Returns notifications for this location
	Type          *int32  // Filter notifications by this type
	Types         *int32  // Return notifications of the bitmask of notification types
	Read          *bool   // Filter by notifications that are read or not
}

func (c *Client) GetNotifications(ctx context.Context, params *GetNotificationsParams) (*NotificationsResponse, error) {
	req := fmt.Sprintf("%s/users/%d/notifications", c.BaseURL, c.JWT.UserID)
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
		if params.LocationID != nil {
			query.Set("location_id", fmt.Sprintf("%d", *params.LocationID))
		}
		if params.Type != nil {
			query.Set("type", fmt.Sprintf("%d", *params.Type))
		}
		if params.Types != nil {
			query.Set("types", fmt.Sprintf("%d", *params.Types))
		}
		if params.Read != nil {
			query.Set("read", fmt.Sprintf("%t", *params.Read))
		}
	}
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var resp NotificationsResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
