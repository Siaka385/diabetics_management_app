package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func Dashboard(db *gorm.DB, tmpl *template.Template, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			log.Printf("Error retrieving session: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Debug print
		fmt.Printf("Session Values: %+v\n", session.Values)

		// Check if user is authenticated
		auth, ok := session.Values["authenticated"].(bool)
		username, usernameOk := session.Values["username"].(string)

		// Validate authentication
		if !ok || !auth || !usernameOk {
			fmt.Println("Redirecting to login: Not authenticated")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// Serve dashboard page if authenticated
		if err := tmpl.ExecuteTemplate(w, "UserDashboard.html", username); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
