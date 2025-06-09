package goflume

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type JWTPayload struct {
	Exp    int      `json:"exp"`
	Iat    int      `json:"iat"`
	Iss    string   `json:"iss"`
	Scope  []string `json:"scope"`
	Sub    string   `json:"sub"`
	Type   string   `json:"type"`
	UserID int      `json:"user_id"`
}

type TokenResponse struct {
	APIResponseEnvelope
	Data []Token `json:"data"`
}

func (c *Client) Authenticate(ctx context.Context, username, password string) error {
	url := c.BaseURL + "/oauth/token"
	reqBody := map[string]string{
		"grant_type":    "password",
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
		"username":      username,
		"password":      password,
	}
	return c.getToken(ctx, url, reqBody)
}

func (c *Client) RefreshAccessToken(ctx context.Context) error {
	url := c.BaseURL + "/oauth/token"
	reqBody := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
		"refresh_token": c.Token.RefreshToken,
	}
	return c.getToken(ctx, url, reqBody)
}

func (c *Client) getToken(ctx context.Context, url string, reqBody map[string]string) error {
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != 200 {
		dat, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("auth failed: status %d body: %s", resp.StatusCode, dat)
	}
	var tr TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return err
	}

	if len(tr.Data) > 0 {
		c.Token = tr.Data[0]

		jwt, err := extractJWTPayload(c.Token.AccessToken)
		if err != nil {
			return err
		}

		if jwt != nil {
			c.JWT = *jwt
		}
	} else {
		return errors.New("no token data received")
	}

	return nil
}

func extractJWTPayload(token string) (*JWTPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil, errors.New("invalid JWT token")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	var claims JWTPayload
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}
	return &claims, nil
}
