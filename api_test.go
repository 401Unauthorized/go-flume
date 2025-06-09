package goflume

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetUser_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"email_address":"a@b.com"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetUser(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].EmailAddress != "a@b.com" {
		t.Errorf("unexpected user data: %+v", got.Data)
	}
}

func TestGetDevices_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"d1"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetDevices(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].ID != "d1" {
		t.Errorf("unexpected device data: %+v", got.Data)
	}
}

func TestGetDevice_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"d1"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetDevice(context.Background(), "d1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].ID != "d1" {
		t.Errorf("unexpected device data: %+v", got.Data)
	}
}

func TestQueryUsage_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"value":42,"datetime":"2024-01-01T00:00:00Z"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.QueryUsage(context.Background(), "d1", QueryUsageRequestBody{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Value != 42 {
		t.Errorf("unexpected usage data: %+v", got.Data)
	}
}

func TestGetCurrentFlow_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"active":true,"gpm":1.23,"datetime":"2024-01-01T00:00:00Z"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetCurrentFlow(context.Background(), "d1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || !got.Data[0].Active {
		t.Errorf("unexpected flow data: %+v", got.Data)
	}
}

func TestGetLocations_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"name":"Home"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetLocations(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "Home" {
		t.Errorf("unexpected locations data: %+v", got.Data)
	}
}

func TestGetLocation_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"name":"Home"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetLocation(context.Background(), "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "Home" {
		t.Errorf("unexpected location data: %+v", got.Data)
	}
}

func TestUpdateLocation_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"success":true}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.UpdateLocation(context.Background(), "1", LocationPatch{AwayMode: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got.Success {
		t.Errorf("expected success true, got %+v", got)
	}
}

func TestGetBudgets_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"name":"Budget1"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetBudgets(context.Background(), "d1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "Budget1" {
		t.Errorf("unexpected budgets data: %+v", got)
	}
}

func TestGetSubscriptions_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"alert_type":"leak"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetSubscriptions(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].AlertType != "leak" {
		t.Errorf("unexpected subscriptions data: %+v", got.Data)
	}
}

func TestGetSubscription_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":{"id":1,"alert_type":"leak"}}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetSubscription(context.Background(), "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Data.AlertType != "leak" {
		t.Errorf("unexpected subscription data: %+v", got.Data)
	}
}

func TestGetNotifications_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"message":"test"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetNotifications(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Message != "test" {
		t.Errorf("unexpected notifications data: %+v", got.Data)
	}
}

func TestGetUsageAlerts_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"device_id":"d1"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetUsageAlerts(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].DeviceID != "d1" {
		t.Errorf("unexpected usage alerts data: %+v", got)
	}
}

func TestGetEventRules_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"r1","name":"Rule1"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetEventRules(context.Background(), "d1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "Rule1" {
		t.Errorf("unexpected event rules data: %+v", got.Data)
	}
}

func TestGetUsageAlertRules_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"r1","name":"AlertRule1"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetUsageAlertRules(context.Background(), "d1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "AlertRule1" {
		t.Errorf("unexpected usage alert rules data: %+v", got.Data)
	}
}

func TestGetUsageAlertRule_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"r1","name":"AlertRule1"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetUsageAlertRule(context.Background(), "d1", "r1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Data[0].Name != "AlertRule1" {
		t.Errorf("unexpected usage alert rule data: %+v", got.Data)
	}
}

func TestGetContacts_withMockClient(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"type":"email","detail":"a@b.com"}]}`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetContacts(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Type != "email" {
		t.Errorf("unexpected contacts data: %+v", got.Data)
	}
}

func TestGetUser_withMockClient_FailureCases(t *testing.T) {
	// HTTP error
	resp := &http.Response{
		StatusCode: 500,
		Body:       io.NopCloser(strings.NewReader(`Internal Server Error`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUser(context.Background())
	if err == nil {
		t.Error("expected error for HTTP 500, got nil")
	}

	// Malformed JSON
	resp = &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"email_address":`)),
		Header:     make(http.Header),
	}
	client = newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err = client.GetUser(context.Background())
	if err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}

	// Missing data field
	resp = &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"foo":123}`)),
		Header:     make(http.Header),
	}
	client = newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetUser(context.Background())
	if err != nil {
		t.Errorf("unexpected error for missing data field: %v", err)
	}
	if len(got.Data) != 0 {
		t.Errorf("expected empty data for missing field, got: %+v", got.Data)
	}
}

func TestGetDevices_withMockClient_FailureCases(t *testing.T) {
	// HTTP error
	resp := &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(strings.NewReader(`Not Found`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetDevices(context.Background(), nil)
	if err == nil {
		t.Error("expected error for HTTP 404, got nil")
	}

	// Malformed JSON
	resp = &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":`)),
		Header:     make(http.Header),
	}
	client = newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err = client.GetDevices(context.Background(), nil)
	if err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}
}

func TestQueryUsage_withMockClient_FailureCases(t *testing.T) {
	// HTTP error
	resp := &http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(strings.NewReader(`Bad Request`)),
		Header:     make(http.Header),
	}
	client := newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.QueryUsage(context.Background(), "d1", QueryUsageRequestBody{})
	if err == nil {
		t.Error("expected error for HTTP 400, got nil")
	}

	// Malformed JSON
	resp = &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"value":`)),
		Header:     make(http.Header),
	}
	client = newMockClient(resp, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err = client.QueryUsage(context.Background(), "d1", QueryUsageRequestBody{})
	if err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}
}

func TestGetDevices_withAllParams(t *testing.T) {
	limit := int32(10)
	offset := int32(5)
	sortField := "type"
	sortDirection := "DESC"
	user := true
	location := true
	listShared := true
	primaryLocation := true
	locationID := int32(123)
	typeVal := int32(2)
	params := &DevicesParams{
		Limit:           &limit,
		Offset:          &offset,
		SortField:       &sortField,
		SortDirection:   &sortDirection,
		User:            &user,
		Location:        &location,
		ListShared:      &listShared,
		PrimaryLocation: &primaryLocation,
		LocationID:      &locationID,
		Type:            &typeVal,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"d1"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetDevices(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].ID != "d1" {
		t.Errorf("unexpected device data: %+v", got.Data)
	}
	if !strings.Contains(capturedURL, "limit=10") ||
		!strings.Contains(capturedURL, "offset=5") ||
		!strings.Contains(capturedURL, "sort_field=type") ||
		!strings.Contains(capturedURL, "sort_direction=DESC") ||
		!strings.Contains(capturedURL, "user=true") ||
		!strings.Contains(capturedURL, "location=true") ||
		!strings.Contains(capturedURL, "list_shared=true") ||
		!strings.Contains(capturedURL, "primary_location=true") ||
		!strings.Contains(capturedURL, "location_id=123") ||
		!strings.Contains(capturedURL, "type=2") {
		t.Errorf("not all params present in query: %s", capturedURL)
	}
}

func TestGetLocations_withAllParams(t *testing.T) {
	limit := int32(20)
	offset := int32(2)
	sortField := "name"
	sortDirection := "DESC"
	listShared := true
	params := &GetLocationsParams{
		Limit:         &limit,
		Offset:        &offset,
		SortField:     &sortField,
		SortDirection: &sortDirection,
		ListShared:    &listShared,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"name":"Home"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetLocations(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "Home" {
		t.Errorf("unexpected locations data: %+v", got.Data)
	}
	if !strings.Contains(capturedURL, "limit=20") ||
		!strings.Contains(capturedURL, "offset=2") ||
		!strings.Contains(capturedURL, "sort_field=name") ||
		!strings.Contains(capturedURL, "sort_direction=DESC") ||
		!strings.Contains(capturedURL, "list_shared=true") {
		t.Errorf("not all params present in query: %s", capturedURL)
	}
}

func TestGetBudgets_withAllParams(t *testing.T) {
	limit := int32(15)
	offset := int32(3)
	sortField := "value"
	sortDirection := "DESC"
	params := &GetBudgetsParams{
		Limit:         &limit,
		Offset:        &offset,
		SortField:     &sortField,
		SortDirection: &sortDirection,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"name":"Budget1"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetBudgets(context.Background(), "d1", params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "Budget1" {
		t.Errorf("unexpected budgets data: %+v", got)
	}
	if !strings.Contains(capturedURL, "limit=15") ||
		!strings.Contains(capturedURL, "offset=3") ||
		!strings.Contains(capturedURL, "sort_field=value") ||
		!strings.Contains(capturedURL, "sort_direction=DESC") {
		t.Errorf("not all params present in query: %s", capturedURL)
	}
}

func TestGetSubscriptions_withAllParams(t *testing.T) {
	limit := int32(5)
	offset := int32(1)
	sortField := "alert_type"
	sortDirection := "DESC"
	alertType := "leak"
	notificationTypes := int32(3)
	notificationType := int32(2)
	deviceID := "d1"
	deviceType := int32(4)
	locationID := int32(7)
	params := &GetSubscriptionsParams{
		Limit:             &limit,
		Offset:            &offset,
		SortField:         &sortField,
		SortDirection:     &sortDirection,
		AlertType:         &alertType,
		NotificationTypes: &notificationTypes,
		NotificationType:  &notificationType,
		DeviceID:          &deviceID,
		DeviceType:        &deviceType,
		LocationID:        &locationID,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"alert_type":"leak"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetSubscriptions(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].AlertType != "leak" {
		t.Errorf("unexpected subscriptions data: %+v", got.Data)
	}
	for _, expect := range []string{
		"limit=5", "offset=1", "sort_field=alert_type", "sort_direction=DESC", "alert_type=leak", "notification_types=3", "notification_type=2", "device_id=d1", "device_type=4", "location_id=7",
	} {
		if !strings.Contains(capturedURL, expect) {
			t.Errorf("missing param %s in query: %s", expect, capturedURL)
		}
	}
}

func TestGetNotifications_withAllParams(t *testing.T) {
	limit := int32(7)
	offset := int32(2)
	sortField := "created_datetime"
	sortDirection := "DESC"
	deviceID := "dev123"
	locationID := int32(42)
	typeVal := int32(3)
	typesVal := int32(5)
	read := true
	params := &GetNotificationsParams{
		Limit:         &limit,
		Offset:        &offset,
		SortField:     &sortField,
		SortDirection: &sortDirection,
		DeviceID:      &deviceID,
		LocationID:    &locationID,
		Type:          &typeVal,
		Types:         &typesVal,
		Read:          &read,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"message":"test"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetNotifications(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Message != "test" {
		t.Errorf("unexpected notifications data: %+v", got.Data)
	}
	for _, expect := range []string{
		"limit=7", "offset=2", "sort_field=created_datetime", "sort_direction=DESC", "device_id=dev123", "location_id=42", "type=3", "types=5", "read=true",
	} {
		if !strings.Contains(capturedURL, expect) {
			t.Errorf("missing param %s in query: %s", expect, capturedURL)
		}
	}
}

func TestGetUsageAlerts_withAllParams(t *testing.T) {
	limit := int32(8)
	offset := int32(4)
	sortField := "id"
	sortDirection := "DESC"
	deviceID := "dev456"
	flumeLeak := true
	params := &GetUsageAlertsParams{
		Limit:         &limit,
		Offset:        &offset,
		SortField:     &sortField,
		SortDirection: &sortDirection,
		DeviceID:      &deviceID,
		FlumeLeak:     &flumeLeak,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"device_id":"dev456"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetUsageAlerts(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].DeviceID != "dev456" {
		t.Errorf("unexpected usage alerts data: %+v", got)
	}
	for _, expect := range []string{
		"limit=8", "offset=4", "sort_field=id", "sort_direction=DESC", "device_id=dev456", "flume_leak=true",
	} {
		if !strings.Contains(capturedURL, expect) {
			t.Errorf("missing param %s in query: %s", expect, capturedURL)
		}
	}
}

func TestGetEventRules_withAllParams(t *testing.T) {
	limit := int32(9)
	offset := int32(6)
	sortField := "id"
	sortDirection := "DESC"
	params := &GetEventRulesParams{
		Limit:         &limit,
		Offset:        &offset,
		SortField:     &sortField,
		SortDirection: &sortDirection,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"r1","name":"Rule1"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetEventRules(context.Background(), "d1", params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "Rule1" {
		t.Errorf("unexpected event rules data: %+v", got.Data)
	}
	for _, expect := range []string{
		"limit=9", "offset=6", "sort_field=id", "sort_direction=DESC",
	} {
		if !strings.Contains(capturedURL, expect) {
			t.Errorf("missing param %s in query: %s", expect, capturedURL)
		}
	}
}

func TestGetUsageAlertRules_withAllParams(t *testing.T) {
	limit := int32(11)
	offset := int32(7)
	sortField := "id"
	sortDirection := "DESC"
	params := &GetUsageAlertRulesParams{
		Limit:         &limit,
		Offset:        &offset,
		SortField:     &sortField,
		SortDirection: &sortDirection,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"r1","name":"AlertRule1"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetUsageAlertRules(context.Background(), "d1", params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "AlertRule1" {
		t.Errorf("unexpected usage alert rules data: %+v", got.Data)
	}
	for _, expect := range []string{
		"limit=11", "offset=7", "sort_field=id", "sort_direction=DESC",
	} {
		if !strings.Contains(capturedURL, expect) {
			t.Errorf("missing param %s in query: %s", expect, capturedURL)
		}
	}
}

func TestGetContacts_withAllParams(t *testing.T) {
	limit := int32(12)
	offset := int32(8)
	sortField := "id"
	sortDirection := "DESC"
	typeVal := "email"
	category := "personal"
	params := &GetContactsParams{
		Limit:         &limit,
		Offset:        &offset,
		SortField:     &sortField,
		SortDirection: &sortDirection,
		Type:          &typeVal,
		Category:      &category,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":1,"type":"email","detail":"a@b.com"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetContacts(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Type != "email" {
		t.Errorf("unexpected contacts data: %+v", got.Data)
	}
	for _, expect := range []string{
		"limit=12", "offset=8", "sort_field=id", "sort_direction=DESC", "type=email", "category=personal",
	} {
		if !strings.Contains(capturedURL, expect) {
			t.Errorf("missing param %s in query: %s", expect, capturedURL)
		}
	}
}

func TestGetDevice_withAllParams(t *testing.T) {
	user := true
	location := true
	params := &DeviceParams{
		User:     &user,
		Location: &location,
	}
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"d1"}]}`)),
		Header:     make(http.Header),
	}
	var capturedURL string
	client := newMockClient(resp, nil, func(req *http.Request) { capturedURL = req.URL.String() })
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	got, err := client.GetDevice(context.Background(), "d1", params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].ID != "d1" {
		t.Errorf("unexpected device data: %+v", got.Data)
	}
	if !strings.Contains(capturedURL, "user=true") || !strings.Contains(capturedURL, "location=true") {
		t.Errorf("not all params present in query: %s", capturedURL)
	}
}

func TestGetDevice_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetDevice(context.Background(), "d1", nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetDevice_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetDevice(context.Background(), "d1", nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetDevices_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetDevices(context.Background(), nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetDevices_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetDevices(context.Background(), nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetLocations_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetLocations(context.Background(), nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetLocations_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetLocations(context.Background(), nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetBudgets_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetBudgets(context.Background(), "d1", nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetBudgets_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetBudgets(context.Background(), "d1", nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetSubscriptions_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetSubscriptions(context.Background(), nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetSubscriptions_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetSubscriptions(context.Background(), nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetNotifications_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetNotifications(context.Background(), nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetNotifications_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetNotifications(context.Background(), nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetUsageAlerts_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlerts(context.Background(), nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetUsageAlerts_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlerts(context.Background(), nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetEventRules_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetEventRules(context.Background(), "d1", nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetEventRules_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetEventRules(context.Background(), "d1", nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetUsageAlertRules_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlertRules(context.Background(), "d1", nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetUsageAlertRules_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlertRules(context.Background(), "d1", nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetContacts_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetContacts(context.Background(), nil)
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetContacts_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetContacts(context.Background(), nil)
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetUser_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUser(context.Background())
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetUser_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUser(context.Background())
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetCurrentFlow_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetCurrentFlow(context.Background(), "d1")
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetCurrentFlow_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetCurrentFlow(context.Background(), "d1")
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetDevice_emptyDeviceID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetDevice(context.Background(), "", nil)
	if err == nil || !strings.Contains(err.Error(), "deviceID cannot be empty") {
		t.Error("expected error for empty deviceID, got nil or wrong error")
	}
}

func TestQueryUsage_emptyDeviceID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.QueryUsage(context.Background(), "", QueryUsageRequestBody{})
	if err == nil || !strings.Contains(err.Error(), "deviceID cannot be empty") {
		t.Error("expected error for empty deviceID, got nil or wrong error")
	}
}

func TestGetCurrentFlow_emptyDeviceID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetCurrentFlow(context.Background(), "")
	if err == nil || !strings.Contains(err.Error(), "deviceID cannot be empty") {
		t.Error("expected error for empty deviceID, got nil or wrong error")
	}
}

func TestGetLocation_emptyLocationID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetLocation(context.Background(), "")
	if err == nil || !strings.Contains(err.Error(), "locationID cannot be empty") {
		t.Error("expected error for empty locationID, got nil or wrong error")
	}
}

func TestUpdateLocation_emptyLocationID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.UpdateLocation(context.Background(), "", LocationPatch{AwayMode: true})
	if err == nil || !strings.Contains(err.Error(), "locationID cannot be empty") {
		t.Error("expected error for empty locationID, got nil or wrong error")
	}
}

func TestGetBudgets_emptyDeviceID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetBudgets(context.Background(), "", nil)
	if err == nil || !strings.Contains(err.Error(), "deviceID cannot be empty") {
		t.Error("expected error for empty deviceID, got nil or wrong error")
	}
}

func TestGetEventRules_emptyDeviceID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetEventRules(context.Background(), "", nil)
	if err == nil || !strings.Contains(err.Error(), "deviceID cannot be empty") {
		t.Error("expected error for empty deviceID, got nil or wrong error")
	}
}

func TestGetUsageAlertRules_emptyDeviceID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlertRules(context.Background(), "", nil)
	if err == nil || !strings.Contains(err.Error(), "deviceID cannot be empty") {
		t.Error("expected error for empty deviceID, got nil or wrong error")
	}
}

func TestGetUsageAlertRule_emptyDeviceID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlertRule(context.Background(), "", "r1")
	if err == nil || !strings.Contains(err.Error(), "deviceID cannot be empty") {
		t.Error("expected error for empty deviceID, got nil or wrong error")
	}
}

func TestGetUsageAlertRule_emptyRuleID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlertRule(context.Background(), "d1", "")
	if err == nil || !strings.Contains(err.Error(), "ruleID cannot be empty") {
		t.Error("expected error for empty ruleID, got nil or wrong error")
	}
}

func TestGetSubscription_emptySubscriptionID(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetSubscription(context.Background(), "")
	if err == nil || !strings.Contains(err.Error(), "subscriptionID cannot be empty") {
		t.Error("expected error for empty subscriptionID, got nil or wrong error")
	}
}

func TestGetLocation_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetLocation(context.Background(), "1")
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetLocation_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":[]}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetLocation(context.Background(), "1")
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestUpdateLocation_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.UpdateLocation(context.Background(), "1", LocationPatch{AwayMode: true})
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestUpdateLocation_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"success":false}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.UpdateLocation(context.Background(), "1", LocationPatch{AwayMode: true})
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetSubscription_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad url"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetSubscription(context.Background(), "subid")
	if err == nil {
		t.Error("expected error for url.Parse, got nil")
	}
}

func TestGetSubscription_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":{}}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetSubscription(context.Background(), "subid")
	if err == nil {
		t.Error("expected error from apiRequest, got nil")
	}
}

func TestGetUsageAlertRule_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad-url://"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlertRule(context.Background(), "dev1", "rule1")
	if err == nil || !strings.Contains(err.Error(), "parse") {
		t.Errorf("expected url.Parse error, got: %v", err)
	}
}

func TestGetUsageAlertRule_apiRequestError(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"data":{}}`)), Header: make(http.Header)}
	client := newMockClient(resp, fmt.Errorf("apiRequest fail"), nil)
	client.BaseURL = "http://x"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.GetUsageAlertRule(context.Background(), "dev1", "rule1")
	if err == nil || !strings.Contains(err.Error(), "apiRequest fail") {
		t.Errorf("expected apiRequest error, got: %v", err)
	}
}

func TestQueryUsage_urlParseError(t *testing.T) {
	client := newMockClient(nil, nil, nil)
	client.BaseURL = ":bad-url://"
	client.JWT = JWTPayload{UserID: 1}
	_, err := client.QueryUsage(context.Background(), "dev1", QueryUsageRequestBody{})
	if err == nil || !strings.Contains(err.Error(), "parse") {
		t.Errorf("expected url.Parse error, got: %v", err)
	}
}
