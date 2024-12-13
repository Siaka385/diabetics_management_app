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

	handler := http.HandlerFunc(GetDefaultMealPlan)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if err := rr.Code; err != http.StatusOK {
		t.Error("Handler returned wrong status")
	}

	expectedResponse := MealPlan{
		Name:        "Diabetic-Friendly Breakfast",
		Calories:    300,
		Protein:     25.0,
		Carbs:       20.0,
		Fats:        10.0,
		Ingredients: []string{"Eggs", "Whole Wheat Bread", "Avocado"},
	}

	var gotResponse MealPlan
	err = json.Unmarshal(rr.Body.Bytes(), &gotResponse)
	if err != nil {
		t.Fatalf("Failed to deserialize response: %v", err)
	}
	if !reflect.DeepEqual(gotResponse, expectedResponse) {
		t.Errorf("Expected '%v' got '%v'", expectedResponse, gotResponse)
	}
}

func TestEditPlan(t *testing.T) {
	edit := MealPlan{
		Name:        "Diabetic-Friendly Breakfast",
		Calories:    300,
		Protein:     25.0,
		Carbs:       20.0,
		Fats:        10.0,
		Ingredients: []string{"Milk", "Whole Wheat Bread", "Molly"},
	}
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
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogMealHandler(tt.args.w, tt.args.r)
		})
	}
}
