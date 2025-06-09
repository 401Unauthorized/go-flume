package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

type mockRoundTripper struct {
	resp      *http.Response
	err       error
	expectReq func(*http.Request)
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.expectReq != nil {
		m.expectReq(req)
	}
	return m.resp, m.err
}

func newMockClient(resp *http.Response, err error, expectReq func(*http.Request)) *Client {
	return &Client{
		HTTPClient: &http.Client{Transport: &mockRoundTripper{resp: resp, err: err, expectReq: expectReq}},
		Token:      Token{AccessToken: "test-token"},
	}
}

type testResp struct {
	Message string `json:"message"`
}

func TestApiRequest_Success(t *testing.T) {
	body := testResp{Message: "ok"}
	b, _ := json.Marshal(body)
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(b)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, func(req *http.Request) {
		if req.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("missing or incorrect Authorization header")
		}
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("missing or incorrect Content-Type header")
		}
		if req.URL.Query().Get("envelope") != "true" {
			t.Errorf("missing envelope query param")
		}
	})
	var got testResp
	u, _ := url.Parse("https://api.flumewater.com/api")
	err := client.apiRequest(context.Background(), http.MethodGet, u, nil, &got)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, body) {
		t.Errorf("unexpected response: got %+v, want %+v", got, body)
	}
}

func TestApiRequest_ErrorStatus(t *testing.T) {
	resp := &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(strings.NewReader("not found")),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	u, _ := url.Parse("https://api.flumewater.com/api")
	err := client.apiRequest(context.Background(), http.MethodGet, u, nil, nil)
	if err == nil || !strings.Contains(err.Error(), "API error") {
		t.Errorf("expected API error, got %v", err)
	}
}

func TestApiRequest_HttpClientError(t *testing.T) {
	client := newMockClient(nil, errors.New("network error"), nil)
	u, _ := url.Parse("https://api.flumewater.com/api")
	err := client.apiRequest(context.Background(), http.MethodGet, u, nil, nil)
	if err == nil || !strings.Contains(err.Error(), "network error") {
		t.Errorf("expected network error, got %v", err)
	}
}

func TestApiRequest_NilUrl(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	err := client.apiRequest(context.Background(), http.MethodGet, nil, nil, nil)
	if err == nil || !strings.Contains(err.Error(), "endpoint cannot be nil") {
		t.Errorf("expected endpoint error, got %v", err)
	}
}

func TestApiRequest_ReqBody(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"message":"ok"}`)),
		Header:     make(http.Header),
	}
	var gotBody map[string]string
	client := newMockClient(resp, nil, func(req *http.Request) {
		b, _ := io.ReadAll(req.Body)
		err := json.Unmarshal(b, &gotBody)
		if err != nil {
			t.Fail()
		}
	})
	u, _ := url.Parse("https://api.flumewater.com")
	inBody := map[string]string{"foo": "bar"}
	var got testResp
	err := client.apiRequest(context.Background(), http.MethodPost, u, inBody, &got)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotBody["foo"] != "bar" {
		t.Errorf("expected request body to contain foo=bar, got %+v", gotBody)
	}
}

func TestApiRequest_NoRespBody(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("")),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	u, _ := url.Parse("https://api.flumewater.com/api")
	err := client.apiRequest(context.Background(), http.MethodGet, u, nil, nil)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
