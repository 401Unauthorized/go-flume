package goflume

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockAuthRoundTripper struct {
	resp      *http.Response
	err       error
	expectReq func(*http.Request)
}

func (m *mockAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.expectReq != nil {
		m.expectReq(req)
	}
	return m.resp, m.err
}

func newMockAuthClient(resp *http.Response, err error, expectReq func(*http.Request)) *Client {
	return &Client{
		HTTPClient: &http.Client{Transport: &mockAuthRoundTripper{resp: resp, err: err, expectReq: expectReq}},
		BaseURL:    "https://api.example.com",
	}
}

func TestAuthenticate_Success(t *testing.T) {
	jwt := validJWTToken()
	respBody := TokenResponse{
		APIResponseEnvelope: APIResponseEnvelope{Success: true},
		Data:                []Token{{AccessToken: jwt, RefreshToken: "def", ExpiresIn: 3600, TokenType: "Bearer"}},
	}
	b, _ := json.Marshal(respBody)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(b))),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, func(req *http.Request) {
		if req.Method != http.MethodPost {
			t.Errorf("expected POST method")
		}
		if !strings.Contains(req.URL.String(), "/oauth/token") {
			t.Errorf("expected /oauth/token endpoint")
		}
	})
	err := client.Authenticate(context.Background(), "user", "pass")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client.Token.AccessToken != jwt {
		t.Errorf("expected access token to be set")
	}
}

func TestAuthenticate_ErrorStatus(t *testing.T) {
	resp := &http.Response{
		StatusCode: 401,
		Body:       io.NopCloser(strings.NewReader(`{"success":false,"message":"unauthorized"}`)),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	err := client.Authenticate(context.Background(), "user", "pass")
	if err == nil || !strings.Contains(err.Error(), "auth failed: status 401 body: ") {
		t.Errorf("expected auth failed error, got %v", err)
	}
}

func TestAuthenticate_HttpClientError(t *testing.T) {
	client := newMockAuthClient(nil, context.DeadlineExceeded, nil)
	err := client.Authenticate(context.Background(), "user", "pass")
	if err == nil || !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Errorf("expected context error, got %v", err)
	}
}

func TestAuthenticate_InvalidJWTToken(t *testing.T) {
	// Not enough parts
	respBody := TokenResponse{
		APIResponseEnvelope: APIResponseEnvelope{Success: true},
		Data:                []Token{{AccessToken: "notajwt", RefreshToken: "def", ExpiresIn: 3600, TokenType: "Bearer"}},
	}
	b, _ := json.Marshal(respBody)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(b))),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	err := client.Authenticate(context.Background(), "user", "pass")
	if err == nil || !strings.Contains(err.Error(), "invalid JWT token") {
		t.Errorf("expected invalid JWT token error, got %v", err)
	}
}

func TestAuthenticate_BadBase64JWT(t *testing.T) {
	// Payload is not valid base64 ("!!!" is not valid base64)
	jwt := "header.!!!.signature"
	respBody := TokenResponse{
		APIResponseEnvelope: APIResponseEnvelope{Success: true},
		Data:                []Token{{AccessToken: jwt, RefreshToken: "def", ExpiresIn: 3600, TokenType: "Bearer"}},
	}
	b, _ := json.Marshal(respBody)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(b))),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	err := client.Authenticate(context.Background(), "user", "pass")
	if err == nil || !strings.Contains(err.Error(), "illegal base64 data") {
		t.Errorf("expected base64 decode error, got %v", err)
	}
}

func TestAuthenticate_BadJSONJWT(t *testing.T) {
	// Valid base64, but not valid JSON
	payload := base64.RawURLEncoding.EncodeToString([]byte("notjson"))
	jwt := "header." + payload + ".signature"
	respBody := TokenResponse{
		APIResponseEnvelope: APIResponseEnvelope{Success: true},
		Data:                []Token{{AccessToken: jwt, RefreshToken: "def", ExpiresIn: 3600, TokenType: "Bearer"}},
	}
	b, _ := json.Marshal(respBody)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(b))),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	err := client.Authenticate(context.Background(), "user", "pass")
	if err == nil || !strings.Contains(err.Error(), "invalid character") {
		t.Errorf("expected json unmarshal error, got %v", err)
	}
}

func TestAuthenticate_MissingJSONJWT(t *testing.T) {
	respBody := TokenResponse{
		APIResponseEnvelope: APIResponseEnvelope{Success: true},
		Data:                []Token{},
	}
	b, _ := json.Marshal(respBody)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(b))),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	err := client.Authenticate(context.Background(), "user", "pass")
	if err == nil || !strings.Contains(err.Error(), "no token data received") {
		t.Errorf("expected no token data received, got %v", err)
	}
}

func TestRefreshAccessToken_Success(t *testing.T) {
	respBody := TokenResponse{
		APIResponseEnvelope: APIResponseEnvelope{Success: true},
		Data:                []Token{{AccessToken: validJWTToken(), RefreshToken: "refresh", ExpiresIn: 3600, TokenType: "Bearer"}},
	}
	b, _ := json.Marshal(respBody)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(b))),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	client.ClientID = "id"
	client.ClientSecret = "secret"
	client.Token.RefreshToken = "refresh"
	err := client.RefreshAccessToken(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client.Token.AccessToken == "" {
		t.Errorf("expected access token to be set")
	}
}

func TestRefreshAccessToken_ErrorStatus(t *testing.T) {
	resp := &http.Response{
		StatusCode: 401,
		Body:       io.NopCloser(strings.NewReader(`{"success":false,"message":"unauthorized"}`)),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	client.ClientID = "id"
	client.ClientSecret = "secret"
	client.Token.RefreshToken = "refresh"
	err := client.RefreshAccessToken(context.Background())
	if err == nil || !strings.Contains(err.Error(), "auth failed") {
		t.Errorf("expected auth failed error, got %v", err)
	}
}

func TestRefreshAccessToken_HttpClientError(t *testing.T) {
	client := newMockAuthClient(nil, context.DeadlineExceeded, nil)
	client.ClientID = "id"
	client.ClientSecret = "secret"
	client.Token.RefreshToken = "refresh"
	err := client.RefreshAccessToken(context.Background())
	if err == nil || !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Errorf("expected context error, got %v", err)
	}
}

func TestGetToken_StatusNot200(t *testing.T) {
	resp := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(strings.NewReader("forbidden")),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	err := client.getToken(context.Background(), "https://api.example.com/oauth/token", map[string]string{"foo": "bar"})
	if err == nil || !strings.Contains(err.Error(), "auth failed") {
		t.Errorf("expected auth failed error, got %v", err)
	}
}

func TestGetToken_InvalidJSON(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("not json")),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	err := client.getToken(context.Background(), "https://api.example.com/oauth/token", map[string]string{"foo": "bar"})
	if err == nil || !strings.Contains(err.Error(), "invalid character") {
		t.Errorf("expected json error, got %v", err)
	}
}

func TestGetToken_ExtractJWTPayloadError(t *testing.T) {
	respBody := TokenResponse{
		APIResponseEnvelope: APIResponseEnvelope{Success: true},
		Data:                []Token{{AccessToken: "notajwt", RefreshToken: "def", ExpiresIn: 3600, TokenType: "Bearer"}},
	}
	b, _ := json.Marshal(respBody)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(b))),
		Header:     make(http.Header),
	}
	client := newMockAuthClient(resp, nil, nil)
	err := client.getToken(context.Background(), "https://api.example.com/oauth/token", map[string]string{"foo": "bar"})
	if err == nil || !strings.Contains(err.Error(), "invalid JWT token") {
		t.Errorf("expected invalid JWT token error, got %v", err)
	}
}

func validJWTToken() string {
	// header: {"alg":"HS256","typ":"JWT"}
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	// payload: {"exp":9999999999,"iat":1111111111,"iss":"issuer","scope":["a"],"sub":"sub","type":"type","user_id":1}
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"exp":9999999999,"iat":1111111111,"iss":"issuer","scope":["a"],"sub":"sub","type":"type","user_id":1}`))
	return header + "." + payload + ".sig"
}
