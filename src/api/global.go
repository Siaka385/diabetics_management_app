package api

import "gorm.io/gorm"

// Meal Plan
type MealPlanRequest struct {
	Duration           string   `json:"mp_duration"`
	MealTypes          string   `json:"mp_type"`
	DietaryPreferences string   `json:"mp_diet_pref"`
	PreferredFoods     []string `json:"mp_preferred_foods"`
	FoodRestrictions   []string `json:"mp_avoid_foods"`
}

type Meal struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"` // Breakfast, Lunch, etc.
	Ingredients []string `json:"ingredients"`
}

type MealPlanResponse struct {
	Duration string `json:"duration"`
	Meals    []Meal `json:"meals"`
	Message  string `json:"message"`
}

// Meal log
type NutritionResponse struct {
	Message      string `json:"message"`
	MealInsights string `json:"mealInsights"`
}

type NutrientInfo struct {
	gorm.Model
	UserID   string  `json:"user_id"`
	Calories float64 `json:"calories"`
	Carbs    float64 `json:"carbs"`
	Protein  float64 `json:"protein"`
	Fat      float64 `json:"fat"`
	Fiber    float64 `json:"fiber"`
	Vitamins string  `json:"vitamins"`
	Minerals string  `json:"minerals"`
}

type FoodLog struct {
	gorm.Model
	UserID    string     `json:"user_id"`
	MealItems []MealItem `gorm:"foreignKey:FoodLogID" json:"meal_items"`
}

type MealItem struct {
	gorm.Model
	FoodItem   string  `json:"food_item"`
	Weight     float64 `json:"weight"`
	Proportion float64 `json:"proportion"`
	FoodLogID  uint    `json:"-"`
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
		Name:        "Fish",
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
		Name:        "Chicken",
		ServingSize: 100,
		Calories:    165,
		Carbs:       0.0,
		Protein:     31.0,
		Fat:         3.6,
		Fiber:       0.0,
		Vitamins:    map[string]float64{"Vitamin B6": 0.6, "Niacin": 14.8},
		Minerals:    map[string]float64{"Selenium": 0.025, "Phosphorus": 0.2, "Magnesium": 24.0},
	},
	"Apple": {
		Name:        "Apple",
		ServingSize: 100,
		Calories:    52,
		Carbs:       13.8,
		Protein:     0.3,
		Fat:         0.2,
		Fiber:       2.4,
		Vitamins:    map[string]float64{"Vitamin C": 4.6, "Vitamin A": 0.0, "Vitamin K": 2.2},
		Minerals:    map[string]float64{"Potassium": 0.107, "Phosphorus": 0.02, "Magnesium": 0.009},
	},
}

var servings = map[string]float64{
	"Ugali":             1,
	"Kales":             1.5,
	"Fish (Tilapia)":    0.5,
	"Broccoli":          1,
	"Chicken (Grilled)": 2,
}

var defaultMealPlan = FoodLog{
	UserID: "default-user",
	MealItems: []MealItem{
		{
			FoodItem:   "Eggs",
			Weight:     100,
			Proportion: 0.4,
		},
		{
			FoodItem:   "Whole Wheat Bread",
			Weight:     60,
			Proportion: 0.3,
		},
		{
			FoodItem:   "Avocado",
			Weight:     50,
			Proportion: 0.3,
		},
	},
}
