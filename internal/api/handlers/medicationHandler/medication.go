package handlers

import (
	"encoding/json"
	"net/http"

	auth "diawise/internal/api/middleware"
	"diawise/internal/models"
	"diawise/internal/shared"
)

func MedicationHandler() http.HandlerFunc {
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
			CurrentPage: "/medication",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserProfileDetails)
	}
}
