package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	auth "diawise/internal/middleware"
	"diawise/internal/models"

	"diawise/internal/shared"
)

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

func GlucoseTrackerEndPointHandler(w http.ResponseWriter, r *http.Request) {
	// Capture glucose level and date from the request query parameters
	glucoseLevel := r.URL.Query().Get("glucose")
	glucoseDate := r.URL.Query().Get("date")

	glucoseParam := map[string]string{glucoseLevel: glucoseDate}

	// Set response header and JSON encode the glucose level and date
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(glucoseParam)
}

func BloodSugarHandler(tmpl *template.Template) http.HandlerFunc {
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

		if err := tmpl.ExecuteTemplate(w, "bloodsugar.html", UserProfileDetails); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func DietAndNutritionHandler(tmpl *template.Template) http.HandlerFunc {
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
		if err := tmpl.ExecuteTemplate(w, "DietAndNutrition.html", UserProfileDetails); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}
