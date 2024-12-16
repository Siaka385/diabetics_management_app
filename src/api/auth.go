package api

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	auth "diawise/src/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

var mySigningKey = []byte("secret")

func generateJWT(user auth.User) (string, error) {
	// Generate the JWT with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Username,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expiration time (24 hours)
	})

	// Sign the token with the secret key
	return token.SignedString(mySigningKey)
}

func Signup(db *gorm.DB, tmpl *template.Template, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "signup.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func SignupUser(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
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
		success := auth.RegisterUser(db, user.Username, user.Email, user.Password)
		if !success {
			// If there is an error (e.g., email already registered), return error response
			response = map[string]string{"status": "error", "message": "Unable to register user"}
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			// User registered successfully
			response = map[string]string{
				"status":   "success",
				"message":  "User registered successfully",
				"redirect": "/", // redirect the user to dashboard after successful registration
			}
			w.WriteHeader(http.StatusCreated)
		}
		// Send the response back to the frontend
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

func LoginUserSuccess(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "login-success.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func LoginUser(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Define response map
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

		// Generate a JWT token
		token, err := generateJWT(*user)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}

		// Set the token in a secure HttpOnly cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "authToken",
			Value:    token,
			Path:     "/",
			HttpOnly: true,                    // Prevents access to the cookie via JavaScript
			Secure:   true,                    // Ensures the cookie is only sent over HTTPS
			SameSite: http.SameSiteStrictMode, // Restricts cross-site usage
		})

		// Send response with the token
		response = map[string]string{
			"status":   "success",
			"message":  "Login successful",
			"token":    token,
			"redirect": "/dashboard",
		}

		// Send the response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// Logout function to clear session
func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete the JWT cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1, // Expire the cookie
	})

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
