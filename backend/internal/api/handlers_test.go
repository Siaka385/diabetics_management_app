package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGlucoseTrackerEndPointHandler(t *testing.T) {
	tests := []struct {
		queryParams      string
		expectedStatus   int
		expectedResponse map[string]string
	}{
		{
			queryParams:    "glucose=120&date=2024-12-13",
			expectedStatus: http.StatusOK,
			expectedResponse: map[string]string{
				"120": "2024-12-13",
			},
		},
		{
			queryParams:    "glucose=&date=2024-12-13",
			expectedStatus: http.StatusOK,
			expectedResponse: map[string]string{
				"": "2024-12-13",
			},
		},
		{
			queryParams:    "glucose=110",
			expectedStatus: http.StatusOK,
			expectedResponse: map[string]string{
				"110": "",
			},
		},
	}

	for _, test := range tests {
		r := httptest.NewRequest(http.MethodGet, "/?"+test.queryParams, nil)
		w := httptest.NewRecorder()

		GlucoseTrackerEndPointHandler(w, r)

		result := w.Result()
		defer result.Body.Close()

		if result.StatusCode != test.expectedStatus {
			t.Errorf("expected status %d, got %d", test.expectedStatus, result.StatusCode)
		}

		var response map[string]string
		if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		for k, v := range test.expectedResponse {
			if response[k] != v {
				t.Errorf("expected response[%q] = %q, got %q", k, v, response[k])
			}
		}
	}
}

func TestBlogHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	BlogHomeHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if !contains(w.Body.String(), "Blog Posts") {
		t.Fatalf("expected page to contain 'Blog Posts'")
	}
}

func contains(content, substr string) bool {
	return strings.Contains(content, substr)
}