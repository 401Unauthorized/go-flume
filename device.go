package main

import (
	"context"
	"fmt"
	"net/url"
)

type DevicesParams struct {
	Limit           *int32  // How many devices to return (Defaults to 50)
	Offset          *int32  // Offset of devices to return (Defaults to 0)
	SortField       *string // Which field to sort devices on (Defaults to id)
	SortDirection   *string // Which direction to sort devices on (Defaults to ASC)
	User            *bool   // Include user data in response (Defaults to false)
	Location        *bool   // Include location data in response (Defaults to false)
	ListShared      *bool   // Include devices with shared access (Defaults to false)
	PrimaryLocation *bool   // Only include devices associated with a primary location if true
	LocationID      *int32  // Find devices associated with a specified location ID
	Type            *int32  // Filter devices by their type
}

type Device struct {
	ID           string `json:"id"`
	Type         int    `json:"type"`
	LocationID   int    `json:"location_id"`
	UserID       int    `json:"user_id"`
	BridgeID     string `json:"bridge_id"`
	Oriented     bool   `json:"oriented"`
	LastSeen     string `json:"last_seen"`
	Connected    bool   `json:"connected"`
	BatteryLevel string `json:"battery_level"`
	Product      string `json:"product"`
}

type DevicesResponse struct {
	APIResponseEnvelope
	Data []Device `json:"data"`
}

func (c *Client) GetDevices(ctx context.Context, params *DevicesParams) (*DevicesResponse, error) {
	req := fmt.Sprintf("%s/users/%d/devices", c.BaseURL, c.JWT.UserID)
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
		if params.User != nil {
			query.Set("user", fmt.Sprintf("%t", *params.User))
		}
		if params.Location != nil {
			query.Set("location", fmt.Sprintf("%t", *params.Location))
		}
		if params.ListShared != nil {
			query.Set("list_shared", fmt.Sprintf("%t", *params.ListShared))
		}
		if params.PrimaryLocation != nil {
			query.Set("primary_location", fmt.Sprintf("%t", *params.PrimaryLocation))
		}
		if params.LocationID != nil {
			query.Set("location_id", fmt.Sprintf("%d", *params.LocationID))
		}
		if params.Type != nil {
			query.Set("type", fmt.Sprintf("%d", *params.Type))
		}
	}
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var resp DevicesResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeviceResponse struct {
	APIResponseEnvelope
	Data []Device `json:"data"`
}

type DeviceParams struct {
	User     *bool // Include user data in response (Defaults to false)
	Location *bool // Include location data in response (Defaults to false)
}

func (c *Client) GetDevice(ctx context.Context, deviceID string, params *DeviceParams) (*DeviceResponse, error) {
	if deviceID == "" {
		return nil, fmt.Errorf("deviceID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/devices/%s", c.BaseURL, c.JWT.UserID, deviceID)
	query := url.Values{}
	if params != nil {
		if params.User != nil {
			query.Set("user", fmt.Sprintf("%t", *params.User))
		}
		if params.Location != nil {
			query.Set("location", fmt.Sprintf("%t", *params.Location))
		}
	}
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var resp DeviceResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
