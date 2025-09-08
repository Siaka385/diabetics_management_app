package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "diawise/internal/middleware"
	"diawise/internal/repository"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func Signup(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var signupData struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&signupData); err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request"})
			return
		}

		if signupData.Username == "" || signupData.Email == "" || signupData.Password == "" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "All fields are required",
			})
			return
		}

		success := repository.RegisterUser(db, signupData.Username, signupData.Username, signupData.Email, signupData.Password)
		if !success {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "Failed to register user",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "User registered successfully",
		})
	}
}

func Signin(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request"})
			return
		}

		user, err := repository.LoginUser(db, loginData.Username, loginData.Password)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "Invalid credentials",
			})
			return
		}

		// Create JWT token
		token, err := auth.CreateToken(user)
		if err != nil {
			fmt.Printf("Token creation failed: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"error": "Token creation failed"})
			return
		}

		// Set JWT token as HTTP-only cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "authToken",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":   "success",
			"message":  "Login successful",
			"redirect": "/dashboard",
		})
	}
}

func Signout(w http.ResponseWriter, r *http.Request) {
		// Clear JWT cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "authToken",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(map[string]string{
			"status":   "success",
			"message":  "Signed out successfully",
			"redirect": "/",
		}); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
}

// AuthStatusHandler returns the current authentication status
func AuthStatus(w http.ResponseWriter, r *http.Request) {
		// Try to get auth cookie
		cookie, err := r.Cookie("authToken")
		if err != nil || cookie == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"authenticated": false,
			})
			return
		}

		// Validate token
		user, err := auth.ParseToken(cookie.Value)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"authenticated": false,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"authenticated": true,
			"user": map[string]interface{}{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
		})
}
