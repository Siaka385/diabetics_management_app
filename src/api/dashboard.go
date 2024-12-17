package api

import (
	"html/template"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func Dashboard(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Log user details (optional)
		// fmt.Printf("Authenticated user: %+v\n", user.Name)
		// fmt.Printf("Authenticated user ID: %+v\n", user.ID)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// Serve the dashboard page
		if err := tmpl.ExecuteTemplate(w, "dashboard.html", user.Name); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
