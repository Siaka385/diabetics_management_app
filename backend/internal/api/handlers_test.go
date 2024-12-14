package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetMealSuggestions(t *testing.T) {
	// // Create test input data
	// input := map[string]interface{}{
	// 	"name":        "Grilled Chicken Salad",
	// 	"calories":    350,
	// 	"protein":     30.0,
	// 	"carbs":       20.0,
	// 	"fats":        10.0,
	// 	"servingSize": "1 bowl",
	// }

	// data, err := json.Marshal(input)
	// if err != nil {
	// 	t.Fatalf("Failed to serialize input data: %v", err)
	// }

	// req, err := http.NewRequest("GET", "/nutrition/suggestions", bytes.NewReader(data))
	// if err != nil {
	// 	t.Fatalf("Failed to create request: %v", err)
	// }

	// handler := http.HandlerFunc(GetMealSuggestions)
	// rr := httptest.NewRecorder()

	// handler.ServeHTTP(rr, req)

	// if err := rr.Code; err != http.StatusOK {
	// 	t.Error("Handlre returned wrong status")
	// }

	// expectedResponse := map[string]string{
	// 	"message":      "Meal suggestions received",
	// 	"mealInsights": "The meal 'Grilled Chicken Salad' has 350 calories, 30.00 grams of protein, 20.00 grams of carbs, and 10.00 grams of fats.",
	// }
	// var gotResponse = map[string]string{}
	// err = json.Unmarshal(rr.Body.Bytes(), &gotResponse)
	// if err!= nil {
	//     t.Fatalf("Failed to deserialize response: %v", err)
	// }
	// for key, value := range gotResponse {
	// 	if actualValue, ok := expectedResponse[key]; !ok || actualValue != value {
	//         t.Errorf("Expected '%s' got '%s'", expectedResponse[key], value)
	//     }
	// }
}

func TestGetDefaultMealPlan(t *testing.T) {
	req, err := http.NewRequest("GET", "/nutrition/mealplan", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	handler := http.HandlerFunc(GetMealPlan)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if err := rr.Code; err != http.StatusOK {
		t.Error("Handler returned wrong status")
	}

	expectedResponse := defaultMealPlan

	var gotResponse FoodLog
	err = json.Unmarshal(rr.Body.Bytes(), &gotResponse)
	if err != nil {
		t.Fatalf("Failed to deserialize response: %v", err)
	}
	if !reflect.DeepEqual(gotResponse, expectedResponse) {
		t.Errorf("Expected '%v' got '%v'", expectedResponse, gotResponse)
	}
}

func TestEditPlan(t *testing.T) {
	edit := defaultMealPlan
	data, err := json.Marshal(edit)
	if err != nil {
		t.Fatalf("Failed to serialize input data: %v", err)
	}
	req, err := http.NewRequest("POST", "/nutrition/editplan", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	handler := http.HandlerFunc(EditPlan)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if err := rr.Code; err != http.StatusOK {
		t.Error("Handler returned wrong status")
	}

	expectedResp := map[string]string{"message": "Meal plan updated successfully"}
	var gotResp map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &gotResp)
	if err != nil {
		t.Fatalf("Failed to deserialize response: %v", err)
	}
	if !reflect.DeepEqual(gotResp, expectedResp) {
		t.Errorf("Expected '%v' got '%v'", expectedResp, gotResp)
	}
}

func TestLogMealHandler(t *testing.T) {
	mealLog := defaultMealPlan
	data, err := json.Marshal(mealLog)
	if err != nil {
		t.Fatalf("Failed to serialize input data: %v", err)
	}
	req, err := http.NewRequest("POST", "/nutrition/meal/log", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	handler := http.HandlerFunc(EditPlan)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if err := rr.Code; err != http.StatusOK {
		t.Error("Handler returned wrong status")
	}

	expectedResp := map[string]string{"message": "Meal plan updated successfully"}
	var gotResp map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &gotResp)
	if err != nil {
		t.Fatalf("Failed to deserialize response: %v", err)
	}
	if !reflect.DeepEqual(gotResp, expectedResp) {
		t.Errorf("Expected '%v' got '%v'", expectedResp, gotResp)
	}
}

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
