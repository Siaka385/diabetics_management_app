package api

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func Dashboard(db *gorm.DB, tmpl *template.Template, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
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

		// Serve dashboard page if authenticated
		fmt.Fprintf(w, "Welcome to your dashboard, %s!", username)
		err = tmpl.ExecuteTemplate(w, "dashboard.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}