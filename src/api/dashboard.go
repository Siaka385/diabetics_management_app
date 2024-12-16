package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"diawise/src/auth"

	"gorm.io/gorm"
)

func Dashboard(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		// Set the content type for the response
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// Serve the dashboard page if the user is authenticated
		fmt.Printf("Authenticated user: %+v\n", user.Name)
		fmt.Printf("Authenticated user: %+v\n", user.ID)

		if err := tmpl.ExecuteTemplate(w, "dashboard.html", user.Name); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
