package handlers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"

	"diawise/internal/middleware"
	"diawise/internal/models"
	"diawise/internal/shared"
)

func Dashboard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := middleware.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserProfileDetails := struct {
			models.UserProfile
			CurrentPage string
		}{
			UserProfile: models.UserProfile{
				Abbrev: shared.GenerateShortName(user.Name),
				Name:   user.Name,
			},
			CurrentPage: "/dashboard",
		}

		// Return JSON response
		if err := json.NewEncoder(w).Encode(UserProfileDetails); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
