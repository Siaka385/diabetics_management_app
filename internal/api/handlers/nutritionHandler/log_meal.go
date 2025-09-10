package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	auth "diawise/internal/api/middleware"
	"diawise/internal/models"
	genai "diawise/internal/services/genAi"
	services "diawise/internal/services/nutrition"

	"gorm.io/gorm"
)

func stringifyVitaminsAndMinerals(minrals, vitamins map[string]float64, info models.NutrientInfo) error {
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

func CalculateMealNutrition(foodLog models.FoodLog) (models.NutrientInfo, error) {
	var totalCalories, totalCarbs, totalProtein, totalFat, totalFiber float64
	totalVitamins := make(map[string]float64)
	totalMinerals := make(map[string]float64)

	for _, mealItem := range foodLog.MealItems {
		food, found := foodDatabase[mealItem.FoodItem]
		if !found {
			return models.NutrientInfo{}, fmt.Errorf("food item '%s' not found in the database", mealItem.FoodItem)
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
	info := models.NutrientInfo{
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

func GenerateMealInsights(nutrientInfo models.NutrientInfo) string {
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

func LogMealHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Starting meal logging...")
		// Get user id to use as foreignkey
		// Retrieve the JWT token from cookies
		cookie, err := r.Cookie("authToken")
		if err != nil || cookie == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Parse and validate the JWT token
		tokenString := cookie.Value
		user, err := auth.ParseToken(tokenString)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		userID := user.ID

		var mealEntry models.MealLogEntry
		mealEntry.UserID = userID
		// Decode the request body into the new struct
		err = json.NewDecoder(r.Body).Decode(&mealEntry)
		if err != nil {
			log.Printf("Invalid input: %+v\n", err)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Find or create daily meal log for the current day
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

		var dailyMealLog models.DailyMealLog
		result := db.Where("user_id = ? AND date = ?", userID, today).FirstOrCreate(&dailyMealLog, models.DailyMealLog{UserID: userID, Date: today})
		if result.Error != nil {
			log.Printf("Failed to create/find daily meal log: %+v\n", err)
			http.Error(w, "Failed to create/find daily meal log", http.StatusInternalServerError)
			return
		}

		// Add meal entry to daily meal log
		analyser, err := genai.NewAIHealthAnalyser()
		if err != nil {
			log.Printf("Failed to create AI health analyser: %v", err)
			http.Error(w, "Failed to create AI health analyser", http.StatusInternalServerError)
			return
		}
		dietProfileModel, err := analyser.DietProfile(&mealEntry)
		if err != nil {
			log.Printf("Failed to generate diet profile: %v", err)
			http.Error(w, "Failed to generate diet profile", http.StatusInternalServerError)
			return
		}
		defer analyser.Close()

		dietProfileModel.UserID = userID

		// Convert *models.DietProfile to models.DietProfile
		dietProfile := models.DietProfile{
			UserID:      dietProfileModel.UserID,
			FoodName:    dietProfileModel.FoodName,
			Quantity:    dietProfileModel.Quantity,
			Calories:    dietProfileModel.Calories,
			Carbs:       dietProfileModel.Carbs,
			Protein:     dietProfileModel.Protein,
			Fat:         dietProfileModel.Fat,
			HealthScore: 0,
			Suggestions: "",
			Date:        time.Now(),
		}

		err = services.SaveDietLog(db, dietProfile)
		if err != nil {
			log.Printf("Failed to save diet profile: %v", err)
			http.Error(w, "Failed to save diet profile", http.StatusInternalServerError)
			return
		}

		mealEntry.DailyMealLogID = dailyMealLog.ID
		mealEntry.DietProfileID = dietProfile.ID

		err = services.SaveMealLog(db, mealEntry)
		if err != nil {
			http.Error(w, "Failed to save meal log", http.StatusInternalServerError)
			return
		}

		// Convert services.MealLogEntry to models.MealLogEntry for dailyMealLog
		mealEntryModel := models.MealLogEntry{
			UserID:         mealEntry.UserID,
			DailyMealLogID: mealEntry.DailyMealLogID,
			DietProfileID:  mealEntry.DietProfileID,
			FoodName:       mealEntry.FoodName,
			Quantity:       mealEntry.Quantity,
			Calories:       dietProfileModel.Calories,
			Carbs:          dietProfileModel.Carbs,
			Protein:        dietProfileModel.Protein,
			Fat:            dietProfileModel.Fat,
			Date:           time.Now(),
		}

		dailyMealLog.Entries = append(dailyMealLog.Entries, mealEntryModel)
		db.Save(&dailyMealLog)

		// Send JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			DietProfile models.DietProfile
		}{
			DietProfile: *dietProfileModel,
		})
	}
}
