package handlers

import (
	"diawise/internal/models"
)
	
var foodDatabase = map[string]models.FoodItem{
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

var defaultMealPlan = models.FoodLog{
	UserID: "default-user",
	MealItems: []models.MealItem{
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
