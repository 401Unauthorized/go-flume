package goflume

import (
	"context"
	"fmt"
	"net/url"
)

type User struct {
	ID           int    `json:"id"`
	EmailAddress string `json:"email_address"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
	Status       string `json:"status"`
	Type         string `json:"type"`
}

type UserResponse struct {
	APIResponseEnvelope
	Data []User `json:"data"`
}

func (c *Client) GetUser(ctx context.Context) (*UserResponse, error) {
	req := fmt.Sprintf("%s/users/%d", c.BaseURL, c.JWT.UserID)
	u, err := url.Parse(req)
	if err != nil {
		return nil, err
	}
	var resp UserResponse
	if err := c.apiRequest(ctx, "GET", u, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
