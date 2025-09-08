package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"diawise/internal/repository"
	auth "diawise/internal/middleware"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func Signup(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "signup.html", nil)
		if err != nil {
			InternalServerErrorHandler(w)
		}
	}
}

func SignupUser(db *gorm.DB) http.HandlerFunc {
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

func Login(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			InternalServerErrorHandler(w)
		}
	}
}

func LoginUser(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
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
		fmt.Printf("Creating token for user: %s (ID: %d)\n", user.Name, user.ID)
		token, err := auth.CreateToken(user)
		if err != nil {
			fmt.Printf("Token creation failed: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"error": "Token creation failed"})
			return
		}

		fmt.Printf("Token created, setting cookie\n")
		// Set JWT token as HTTP-only cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "authToken",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})
		fmt.Printf("Cookie set successfully\n")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":   "success",
			"message":  "Login successful",
			"redirect": "/dashboard",
		})
	}
}

func LoginUserSuccess(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "login-success.html", nil)
		if err != nil {
			InternalServerErrorHandler(w)
		}
	}
}

func Logout(sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Clear JWT cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "authToken",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		})

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// AuthStatusHandler returns the current authentication status
func AuthStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
