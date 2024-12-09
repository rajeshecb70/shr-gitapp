package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helper function to test handlers
func testHandler(t *testing.T, handler http.HandlerFunc, expectedStatus int, expectedContentType string) {
	t.Helper() // Mark this function as a helper for better test reporting

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check HTTP status code
	if rr.Code != expectedStatus {
		t.Errorf("Handler returned wrong status code: got %v, want %v", rr.Code, expectedStatus)
	}

	// Check Content-Type header
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned wrong content type: got %v, want %v", contentType, expectedContentType)
	}
}

// TestUptimeHandler tests the uptime handler
func TestUptimeHandler(t *testing.T) {
	testHandler(t, uptimeHandler, http.StatusOK, "application/json")
}

// TestCPUHandler tests the CPU handler
func TestCPUHandler(t *testing.T) {
	testHandler(t, cpuHandler, http.StatusOK, "application/json")
}

// TestLoadHandler tests the load handler
func TestLoadHandler(t *testing.T) {
	testHandler(t, loadHandler, http.StatusOK, "application/json")
}
