package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) apiRequest(ctx context.Context, method string, u *url.URL, reqBody, respBody any) error {
	if u == nil {
		return fmt.Errorf("endpoint cannot be nil")
	}

	var body io.Reader
	if reqBody != nil {
		b, _ := json.Marshal(reqBody)
		body = bytes.NewBuffer(b)
	}
	req, _ := http.NewRequestWithContext(ctx, method, u.String(), body)
	if c.Token.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token.AccessToken)
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Set("envelope", "true")
	req.URL.RawQuery = q.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode >= 400 {
		dat, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s (%d) %s", u, resp.StatusCode, dat)
	}
	if respBody != nil {
		return json.NewDecoder(resp.Body).Decode(respBody)
	}
	return nil
}
