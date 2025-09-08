package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	auth "diawise/internal/middleware"
	"diawise/internal/models"
	"diawise/internal/shared"
)

func MedicationHandler(tmpl *template.Template) http.HandlerFunc {
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
