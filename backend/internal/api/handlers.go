package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MealPlan struct {
	Name       string  `json:"name"`
	Calories   int     `json:"calories"`
	Protein    float64 `json:"protein"`   // in grams
	Carbs      float64 `json:"carbs"`     // in grams
	Fats       float64 `json:"fats"`      // in grams
	ServingSize string `json:"servingSize"`
	Ingredients []string `json:"ingredients"`
}

type NutritionResponse struct {
	Message      string `json:"message"`
	MealInsights string `json:"mealInsights"`
}

var defaultMealPlan = MealPlan{
	Name:        "Diabetic-Friendly Breakfast",
	Calories:    300,
	Protein:     25.0,
	Carbs:       20.0,
	Fats:        10.0,
	Ingredients: []string{"Eggs", "Whole Wheat Bread", "Avocado"},
}

// GetMealSuggestions handles the GET /nutrition/suggestions endpoint
func GetMealSuggestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting meal suggestions...")
	// var meal MealPlan
	// err := json.NewDecoder(r.Body).Decode(&meal)
	// if err != nil {
	// 	http.Error(w, "invalid request from nutrition suggestions", http.StatusBadRequest)
	// 	return
	// }
	// insights := fmt.Sprintf(
	// 	"The meal '%s' has %d calories, %.2f grams of protein, %.2f grams of carbs, and %.2f grams of fats.",
	// 	meal.Name, meal.Calories, meal.Protein, meal.Carbs, meal.Fats,
	// )
	// response := NutritionResponse{
	// 	Message:      "Meal suggestions received",
	// 	MealInsights: insights,
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)
}


func GetDefaultMealPlan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting default meal plan...")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(defaultMealPlan)
}


// curl -v -X GET http://localhost:9000/nutrition/mealplan \
// -H "Content-Type: application/json" \
// -d {
// 	"name": "Diabetic-Friendly Breakfast",
// 	"calories": 300,
// 	"protein": 25.0,
// 	"carbs": 20.0,
// 	"fats": 10.0,
// 	"ingredients": ["Eggs", "Whole Wheat Bread", "Avocado"]
//   }
  