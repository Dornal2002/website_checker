package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateData(t *testing.T) {
	req, err := http.NewRequest("POST", "/create", bytes.NewBufferString(`{"websites":["example.com", "example.org"]}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Validate the response body
	expected := "Data posted Succesfully"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Ensure the map is populated correctly
	if mp == nil {
		t.Error("Map is not initialized")
	}
}

func TestCheckQuery(t *testing.T) {
	// Set up initial data for testing
	mp = map[string]string{
		"example.com": "UP",
		"example.org": "DOWN",
	}

	req, err := http.NewRequest("GET", "/query?websites=example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckQuery)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Validate the response body
	expected := "example.com is :UP"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCheckStatus(t *testing.T) {
	// Set up initial data for testing
	mp = map[string]string{
		"example.com": "UP",
	}

	// Start the CheckStatus goroutine
	go CheckStatus("example.com")

	// Wait for a sufficient time to allow CheckStatus to update the map
	time.Sleep(10 * time.Second)

	// Validate the updated status in the map
	if mp["example.com"] != "DOWN" {
		t.Errorf("CheckStatus did not update status correctly: got %v want %v", mp["example.com"], "DOWN")
	}
}
