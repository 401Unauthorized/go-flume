package main

import (
	"net/http"
	"time"
)

const BaseApiUrl = "https://api.flumewater.com"

type Client struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
	HTTPClient   *http.Client
	Token        Token
	JWT          JWTPayload
}

func NewClient(clientID, clientSecret string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}
	return &Client{
		BaseURL:      BaseApiUrl,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}
}
