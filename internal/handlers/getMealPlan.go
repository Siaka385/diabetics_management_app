package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"diawise/internal/models"
)

// Generate meal plans
func GetMealPlan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting default meal plan...")
	var req models.MealPlanRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Duration != "Single Day" && req.Duration != "Whole Week" {
		http.Error(w, "Invalid duration. Use 'Single Day' or 'Whole Week'.", http.StatusBadRequest)
		return
	}

	if len(req.MealTypes) == 0 {
		http.Error(w, "At least one meal type must be selected.", http.StatusBadRequest)
		return
	}

	mealPlan := generateMealPlan(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mealPlan)
}

func generateMealPlan(mp models.MealPlanRequest) models.MealPlanResponse {
	availableMeals := []models.Meal{
		{Name: "Ugali with Sukuma Wiki", Type: "Lunch", Ingredients: []string{"Ugali", "Kale", "Tomato"}},
		{Name: "Githeri", Type: "Lunch", Ingredients: []string{"Maize", "Beans"}},
		{Name: "Matoke with Beef", Type: "Dinner", Ingredients: []string{"Plantains", "Beef", "Onion"}},
		{Name: "Chai with Bread", Type: "Breakfast", Ingredients: []string{"Tea", "Bread", "Butter"}},
		{Name: "Nyama Choma with Kachumbari", Type: "Dinner", Ingredients: []string{"Grilled Meat", "Tomato", "Onion"}},
	}

	// Filter meals by type and preferences
	filteredMeals := []models.Meal{}
	for _, meal := range availableMeals {
		if mp.MealTypes == meal.Type {
			// Apply dietary preferences filtering (simplified)
			if mp.DietaryPreferences == "Vegetarian" && Contains(meal.Ingredients, "Beef") {
				continue // Skip non-vegetarian meals
			}
			filteredMeals = append(filteredMeals, meal)
		}
	}

	var mealPlan []models.Meal
	if mp.Duration == "Single Day" {
		// Randomly pick meals for the day (one for each type)
		mealPlan = pickOneMealPerType(filteredMeals, mp.MealTypes)
	} else {
		// For whole week, repeat daily meal selection for 7 days
		dayPlan := pickOneMealPerType(filteredMeals, mp.MealTypes)
		for i := 0; i < 7; i++ {
			mealPlan = append(mealPlan, dayPlan...)
		}
	}
	return models.MealPlanResponse{
		Duration: mp.Duration,
		Meals:    mealPlan,
		Message:  "Meal plan generated successfully.",
	}
}

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func pickOneMealPerType(meals []models.Meal, mealType string) []models.Meal {
	result := []models.Meal{}
	for _, meal := range meals {
		if meal.Type == mealType {
			result = append(result, meal)
			break // Pick the first matching meal
		}
	}
	return result
}
