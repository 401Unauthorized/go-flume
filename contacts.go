package main

import (
	"context"
	"fmt"
	"net/url"
)

type Contact struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
}

type ContactsResponse struct {
	APIResponseEnvelope
	Data []Contact `json:"data"`
}

type GetContactsParams struct {
	Limit         *int32  // Max number of contacts to return (Defaults to 50)
	Offset        *int32  // Offset of contacts to return (Defaults to 0)
	SortField     *string // Field to sort contacts on (Defaults to id)
	SortDirection *string // Sort direction (Defaults to ASC)
	Type          *string // Filter by this type of contact information
	Category      *string // Filter by this category of contact information
}

func (c *Client) GetContacts(ctx context.Context, params *GetContactsParams) (*ContactsResponse, error) {
	req := fmt.Sprintf("%s/users/%d/contacts", c.BaseURL, c.JWT.UserID)
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
		if params.Type != nil {
			query.Set("type", *params.Type)
		}
		if params.Category != nil {
			query.Set("category", *params.Category)
		}
	}
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var resp ContactsResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
