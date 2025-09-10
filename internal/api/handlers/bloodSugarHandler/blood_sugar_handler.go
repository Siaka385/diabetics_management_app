package handlers

import (
	"encoding/json"
	"net/http"

	auth "diawise/internal/api/middleware"
	"diawise/internal/models"

	"diawise/internal/shared"
)

var defaultMealPlan models.FoodLog

func EditPlan(w http.ResponseWriter, r *http.Request) {
	var updates models.FoodLog
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	defaultMealPlan = updates

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Meal plan updated successfully"})
}

func BloodSugarHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		UserProfileDetails := struct {
			models.UserProfile
			CurrentPage string
		}{
			UserProfile: models.UserProfile{
				Name:   user.Name,
				Abbrev: shared.GenerateShortName(user.Name),
			},
			CurrentPage: "/bloodsugar",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserProfileDetails)
	}
}

func DietAndNutritionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		UserProfileDetails := struct {
			models.UserProfile
			CurrentPage string
		}{
			UserProfile: models.UserProfile{
				Name:   user.Name,
				Abbrev: shared.GenerateShortName(user.Name),
			},
			CurrentPage: "/nutrition",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserProfileDetails)
	}
}

func GlucoseTrackerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		UserProfileDetails := struct {
			models.UserProfile
			CurrentPage string
		}{
			UserProfile: models.UserProfile{
				Name:   user.Name,
				Abbrev: shared.GenerateShortName(user.Name),
			},
			CurrentPage: "/glucose-tracker",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserProfileDetails)
	}
}
