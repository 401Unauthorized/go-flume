package main

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient_DefaultHTTPClient(t *testing.T) {
	client := NewClient("id", "secret", nil)
	if client.BaseURL != "https://api.flumewater.com" {
		t.Errorf("BaseURL not set correctly: got %s", client.BaseURL)
	}
	if client.ClientID != "id" {
		t.Errorf("ClientID not set correctly: got %s", client.ClientID)
	}
	if client.ClientSecret != "secret" {
		t.Errorf("ClientSecret not set correctly: got %s", client.ClientSecret)
	}
	if client.HTTPClient == nil {
		t.Fatal("HTTPClient should not be nil")
	}
	if client.HTTPClient.Timeout != 10*time.Second {
		t.Errorf("Default HTTPClient timeout should be 10s, got %v", client.HTTPClient.Timeout)
	}
}

func TestNewClient_CustomHTTPClient(t *testing.T) {
	custom := &http.Client{Timeout: 5 * time.Second}
	client := NewClient("id", "secret", custom)
	if client.HTTPClient != custom {
		t.Error("Custom HTTPClient was not set correctly")
	}
}
