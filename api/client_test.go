package api_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/atom-sgt/ctg-client-go/api"
)

const testdataPath = "../testdata"

func TestGetStudies(t *testing.T) {
	// Load the mock response data from the JSON file
	mockResponsePath := filepath.Join(testdataPath, "studies-small.json")
	mockResponseBody, err := os.ReadFile(mockResponsePath)
	if err != nil {
		t.Fatalf("Failed to read mock response file '%s': %v", mockResponsePath, err)
	}

	// Create a mock server using httptest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request path and method
		if r.URL.Path != "/studies" {
			t.Errorf("Expected to request '/studies', got: %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got: %s", r.Method)
		}

		// Set the response header
		w.Header().Set("Content-Type", "application/json")

		// Write the loaded mock data to the response writer
		w.WriteHeader(http.StatusOK) // Explicitly set status OK
		if _, err := w.Write(mockResponseBody); err != nil {
			t.Fatalf("Failed to write mock response body: %v", err)
		}
	}))
	defer server.Close() // Close the server when the test finishes

	// Create a new API client using the mock server's URL
	client := api.New(server.URL)

	// Call the method being tested
	studies, err := client.GetStudies()
	if err != nil {
		t.Fatalf("GetStudies returned an error: %v", err)
	}

	// TODO: Better tests
	expectNCTID := "NCT00072579"
	nctID := studies.Studies[0].ProtocolSection.IdentificationModule.NctID
	if expectNCTID != nctID {
		t.Errorf("Expected %s but found %s", expectNCTID, nctID)
	}
}

// Keep the API error test as it doesn't rely on a specific data file
func TestGetStudies_APIError(t *testing.T) {
	// Create a mock server that returns a 500 Internal Server Error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	client := api.New(server.URL)

	// Call the method being tested
	_, err := client.GetStudies()

	// Check that an error was returned
	if err == nil {
		t.Error("Expected an error from GetAllStudies when API returns non-OK status, but got nil")
	}
}
