package goflume

import (
	"context"
	"fmt"
	"net/url"
)

type Location struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	Name            string `json:"name"`
	PrimaryLocation bool   `json:"primary_location"`
	Address         string `json:"address"`
	Address2        string `json:"address_2"`
	City            string `json:"city"`
	State           string `json:"state"`
	PostalCode      string `json:"postal_code"`
	Country         string `json:"country"`
	TZ              string `json:"tz"`
	Installation    string `json:"installation"`
	InsurerID       int    `json:"insurer_id"`
	BuildingType    string `json:"building_type"`
	AwayMode        bool   `json:"away_mode"`
}

type LocationsResponse struct {
	APIResponseEnvelope
	Data []Location `json:"data"`
}

type GetLocationsParams struct {
	Limit         *int32  // Max number of locations to return (Defaults to 50)
	Offset        *int32  // Offset of locations to return (Defaults to 0)
	SortField     *string // Field to sort locations on (Defaults to id)
	SortDirection *string // Which direction to sort locations on (Defaults to ASC)
	ListShared    *bool   // Include locations with shared access (Defaults to false)
}

func (c *Client) GetLocations(ctx context.Context, params *GetLocationsParams) (*LocationsResponse, error) {
	req := fmt.Sprintf("%s/users/%d/locations", c.BaseURL, c.JWT.UserID)
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
		if params.ListShared != nil {
			query.Set("list_shared", fmt.Sprintf("%t", *params.ListShared))
		}
	}
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var resp LocationsResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type LocationResponse struct {
	APIResponseEnvelope
	Data []Location `json:"data"`
}

func (c *Client) GetLocation(ctx context.Context, locationID string) (*LocationResponse, error) {
	if locationID == "" {
		return nil, fmt.Errorf("locationID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/locations/%s", c.BaseURL, c.JWT.UserID, locationID)
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	var resp LocationResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type LocationPatch struct {
	AwayMode bool `json:"away_mode"`
}

func (c *Client) UpdateLocation(ctx context.Context, locationID string, patch LocationPatch) (*APIResponseEnvelope, error) {
	if locationID == "" {
		return nil, fmt.Errorf("locationID cannot be empty")
	}
	req := fmt.Sprintf("%s/users/%d/locations/%s", c.BaseURL, c.JWT.UserID, locationID)
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	var resp APIResponseEnvelope
	if err := c.apiRequest(ctx, "PATCH", u, patch, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
