package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NutritionResponse struct {
	Message      string `json:"message"`
	MealInsights string `json:"mealInsights"`
}

type NutrientInfo struct {
	UserID   string `json:"user"`
	Calories float64
	Carbs    float64
	Protein  float64
	Fat      float64
	Fiber    float64
	Vitamins map[string]float64
	Minerals map[string]float64
}

type FoodLog struct {
	UserID    string     `json:"user"`
	MealItems []MealItem `json:"meal_items"`
}

type MealItem struct {
	FoodItem   string  `json:"food_item"`
	Weight     float64 `json:"weight"`     // in grams
	Proportion float64 `json:"proportion"` // portion of the plate (0.25, 0.5, etc.)
}

type FoodItem struct {
	Name        string             `json:"name"`
	ServingSize float64            `json:"serving_size"`
	Calories    float64            `json:"calories"`
	Carbs       float64            `json:"carbs"`
	Protein     float64            `json:"protein"`
	Fat         float64            `json:"fat"`
	Fiber       float64            `json:"fiber"`
	Vitamins    map[string]float64 `json:"vitamins"`
	Minerals    map[string]float64 `json:"minerals"`
}

var foodDatabase = map[string]FoodItem{
	"Ugali": {
		Name:        "Ugali",
		ServingSize: 100,
		Calories:    123,
		Carbs:       26.0,
		Protein:     3.0,
		Fat:         1.5,
		Fiber:       1.4,
		Vitamins:    map[string]float64{"Vitamin A": 0.0, "Vitamin C": 0.0},
		Minerals:    map[string]float64{"Calcium": 8.0, "Iron": 1.1},
	},
	"Kales": {
		Name:        "Kales",
		ServingSize: 100,
		Calories:    50,
		Carbs:       9.0,
		Protein:     3.0,
		Fat:         0.9,
		Fiber:       2.0,
		Vitamins:    map[string]float64{"Vitamin A": 241.0, "Vitamin C": 120.0, "Vitamin K": 500.0},
		Minerals:    map[string]float64{"Calcium": 150.0, "Iron": 2.7, "Magnesium": 47.0},
	},
	"Fish": {
		Name:        "Fish (Tilapia)",
		ServingSize: 100,
		Calories:    128,
		Carbs:       0.0,
		Protein:     26.0,
		Fat:         3.0,
		Fiber:       0.0,
		Vitamins:    map[string]float64{"Vitamin D": 0.6, "Vitamin B12": 2.4},
		Minerals:    map[string]float64{"Selenium": 0.03, "Phosphorus": 0.2, "Magnesium": 25.0},
	},
	"Broccoli": {
		Name:        "Broccoli",
		ServingSize: 100,
		Calories:    55,
		Carbs:       11.2,
		Protein:     3.7,
		Fat:         0.6,
		Fiber:       2.6,
		Vitamins:    map[string]float64{"Vitamin A": 31.0, "Vitamin C": 89.2, "Vitamin K": 101.0},
		Minerals:    map[string]float64{"Calcium": 47.0, "Iron": 0.7, "Magnesium": 21.0},
	},
	"Chicken": {
		Name:        "Chicken (Grilled)",
		ServingSize: 100,
		Calories:    165,
		Carbs:       0.0,
		Protein:     31.0,
		Fat:         3.6,
		Fiber:       0.0,
		Vitamins:    map[string]float64{"Vitamin B6": 0.6, "Niacin": 14.8},
		Minerals:    map[string]float64{"Selenium": 0.025, "Phosphorus": 0.2, "Magnesium": 24.0},
	},
}

var servings = map[string]float64{
	"Ugali":             1,   // 1 serving of Ugali
	"Kales":             1.5, // 1.5 servings of Kales
	"Fish (Tilapia)":    0.5, // 0.5 serving of Fish
	"Broccoli":          1,   // 1 serving of Broccoli
	"Chicken (Grilled)": 2,   // 2 servings of Chicken
}

var defaultMealPlan = FoodLog{
	UserID: "default-user", // Placeholder for user ID
	MealItems: []MealItem{
		{
			FoodItem:   "Eggs",
			Weight:     100, // Weight in grams
			Proportion: 0.4, // Represents 40% of the plate
		},
		{
			FoodItem:   "Whole Wheat Bread",
			Weight:     60,  // Weight in grams
			Proportion: 0.3, // Represents 30% of the plate
		},
		{
			FoodItem:   "Avocado",
			Weight:     50,  // Weight in grams
			Proportion: 0.3, // Represents 30% of the plate
		},
	},
}

var foodLogs []FoodLog

// GetMealSuggestions handles the GET /nutrition/suggestions endpoint
func GetMealSuggestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting meal suggestions...")
}

func GetDefaultMealPlan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting default meal plan...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(defaultMealPlan)
}

func EditPlan(w http.ResponseWriter, r *http.Request) {
	var updates FoodLog
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	defaultMealPlan = updates

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Meal plan updated successfully"})
}

// Handler to log a consumed meal
func LogMealHandler(w http.ResponseWriter, r *http.Request) {
	// Parse incoming JSON request
	var foodLog FoodLog
	err := json.NewDecoder(r.Body).Decode(&foodLog)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Calculate nutrition for the meal
	nutrientInfo, err := CalculateMealNutrition(foodLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store the food log and nutrient information (stubbed function, replace with actual implementation)
	err = SaveMealLog(foodLog, nutrientInfo)
	if err != nil {
		http.Error(w, "Failed to save meal log", http.StatusInternalServerError)
		return
	}

	// Generate insights based on the meal
	mealInsights := GenerateMealInsights(nutrientInfo)

	// Respond with the nutrition information and insights
	response := NutritionResponse{
		Message:      "Meal logged successfully!",
		MealInsights: mealInsights,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CalculateMealNutrition(foodLog FoodLog) (NutrientInfo, error) {
	var totalCalories, totalCarbs, totalProtein, totalFat, totalFiber float64
	totalVitamins := make(map[string]float64)
	totalMinerals := make(map[string]float64)

	for _, mealItem := range foodLog.MealItems {
		food, found := foodDatabase[mealItem.FoodItem]
		if !found {
			return NutrientInfo{}, fmt.Errorf("Food item '%s' not found in the database", mealItem.FoodItem)
		}

		// Scale nutrition values by weight and proportion
		scaleFactor := mealItem.Weight / food.ServingSize * mealItem.Proportion

		totalCalories += food.Calories * scaleFactor
		totalCarbs += food.Carbs * scaleFactor
		totalProtein += food.Protein * scaleFactor
		totalFat += food.Fat * scaleFactor
		totalFiber += food.Fiber * scaleFactor

		// Sum vitamins and minerals
		for vitamin, value := range food.Vitamins {
			totalVitamins[vitamin] += value * scaleFactor
		}
		for mineral, value := range food.Minerals {
			totalMinerals[mineral] += value * scaleFactor
		}
	}

	return NutrientInfo{
		UserID:   foodLog.UserID,
		Calories: totalCalories,
		Carbs:    totalCarbs,
		Protein:  totalProtein,
		Fat:      totalFat,
		Fiber:    totalFiber,
		Vitamins: totalVitamins,
		Minerals: totalMinerals,
	}, nil
}

func SaveMealLog(foodLog FoodLog, nutrientInfo NutrientInfo) error {
	return nil
}

func GenerateMealInsights(nutrientInfo NutrientInfo) string {
	if nutrientInfo.Protein > 50 {
		return "Your meal is high in protein, great for muscle building!"
	}
	if nutrientInfo.Fat > 30 {
		return "Your meal contains high fat. Consider balancing it with leaner options."
	}
	if nutrientInfo.Calories < 400 {
		return "Your meal is low in calories. Ensure you are eating enough."
	}
	return "Your meal is well-balanced!"
}
