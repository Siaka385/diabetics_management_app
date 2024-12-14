package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func LogMealHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var foodLog FoodLog
		err := json.NewDecoder(r.Body).Decode(&foodLog)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		nutrientInfo, err := CalculateMealNutrition(foodLog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(nutrientInfo)

		err = SaveMealLog(db, foodLog, nutrientInfo)
		if err != nil {
			http.Error(w, "Failed to save meal log", http.StatusInternalServerError)
			return
		}

		mealInsights := GenerateMealInsights(nutrientInfo)

		response := NutritionResponse{
			Message:      "Meal logged successfully!",
			MealInsights: mealInsights,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func stringifyVitaminsAndMinerals(minrals, vitamins map[string]float64, info NutrientInfo) error {
	vitaminsJSON, err := json.Marshal(vitamins)
	if err != nil {
		return err
	}
	mineralsJSON, err := json.Marshal(minrals)
	if err != nil {
		return err
	}

	info.Vitamins = string(vitaminsJSON)
	info.Minerals = string(mineralsJSON)

	return nil
}

func CalculateMealNutrition(foodLog FoodLog) (NutrientInfo, error) {
	var totalCalories, totalCarbs, totalProtein, totalFat, totalFiber float64
	totalVitamins := make(map[string]float64)
	totalMinerals := make(map[string]float64)

	for _, mealItem := range foodLog.MealItems {
		food, found := foodDatabase[mealItem.FoodItem]
		if !found {
			return NutrientInfo{}, fmt.Errorf("food item '%s' not found in the database", mealItem.FoodItem)
		}

		scaleFactor := mealItem.Weight / food.ServingSize * mealItem.Proportion

		totalCalories += food.Calories * scaleFactor
		totalCarbs += food.Carbs * scaleFactor
		totalProtein += food.Protein * scaleFactor
		totalFat += food.Fat * scaleFactor
		totalFiber += food.Fiber * scaleFactor

		for vitamin, value := range food.Vitamins {
			totalVitamins[vitamin] += value * scaleFactor
		}
		for mineral, value := range food.Minerals {
			totalMinerals[mineral] += value * scaleFactor
		}
	}
	info := NutrientInfo{
		UserID:   foodLog.UserID,
		Calories: totalCalories,
		Carbs:    totalCarbs,
		Protein:  totalProtein,
		Fat:      totalFat,
		Fiber:    totalFiber,
	}
	stringifyVitaminsAndMinerals(totalMinerals, totalVitamins, info)

	return info, nil
}

func SaveMealLog(db *gorm.DB, foodLog FoodLog, nutrientInfo NutrientInfo) error {
	tx := db.Begin()
	if err := tx.Create(&foodLog).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&nutrientInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
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
