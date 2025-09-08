package handlers

import (
	"html/template"
	"log"
	"net/http"

	"gorm.io/gorm"

	"diawise/internal/middleware"
	"diawise/internal/models"
	"diawise/internal/shared"
)

func Dashboard(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := middleware.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		UserProfileDetails := models.UserProfile{
			Abbrev: shared.GenerateShortName(user.Name),
			Name:   user.Name,
		}

		// Serve the dashboard page
		if err := tmpl.ExecuteTemplate(w, "dashboard.html", UserProfileDetails); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
