package api

import (
	"encoding/json"
	"fmt"
)

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
