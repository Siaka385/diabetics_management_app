package api

import (
	"encoding/json"
	"net/http"
	"html/template"

	auth "diawise/internal/auth"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

// func RegisterUser(db *gorm.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		var response map[string]string

// 		var user auth.User
// 		// Decode the incoming JSON to a Piece struct
// 		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 			http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 			return
// 		}

// 		err := auth.RegisterUser(db, user.Username, user.Email, user.Password)
// 		if !err {
// 			response = map[string]string{"status": "error", "message": "unable to register user"}
// 			w.WriteHeader(http.StatusInternalServerError)
// 		} else {
// 			response = map[string]string{"status": "success", "message": "User registered successfully"}
// 		}

// 		// Return a 201 Created status
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(response)
// 	}
// }

func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var response map[string]string
		var user auth.User

		// Decode the incoming JSON
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Register user
		err := auth.RegisterUser(db, user.Username, user.Email, user.Password)
		if !err {
			response = map[string]string{"status": "error", "message": "unable to register user"}
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			response = map[string]string{"status": "success", "message": "User registered successfully"}
			w.WriteHeader(http.StatusCreated)
		}

		json.NewEncoder(w).Encode(response)
	}
}

func Login(db *gorm.DB, tmpl *template.Template, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Check if user is already authenticated
		if auth, ok := session.Values["authenticated"].(bool); ok && auth {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}

		err = tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func LoginUser(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var response map[string]string

		// Decode the incoming JSON body
		var userInput struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Try to find the user in the database
		user, err := auth.LoginUser(db, userInput.Username, userInput.Password)
		if err != nil {
			response = map[string]string{"status": "error", "message": "unable to login user"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Set session values
		session.Values["authenticated"] = true
		session.Values["username"] = user.Username
		session.Values["user_id"] = user.ID

		// Save the session
		if err := session.Save(r, w); err != nil {
			http.Error(w, "Error saving session", http.StatusInternalServerError)
			return
		}

		// Send JSON response with the redirect URL
		response = map[string]string{
			"status":   "success",
			"message":  "Login successful",
			"redirect": "/dashboard",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// Logout function to clear session
func Logout(sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Clear session values
		session.Values["authenticated"] = false
		session.Values["username"] = ""
		session.Values["user_id"] = nil

		// Save the session to clear it
		if err := session.Save(r, w); err != nil {
			http.Error(w, "Error clearing session", http.StatusInternalServerError)
			return
		}

		// Redirect to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
