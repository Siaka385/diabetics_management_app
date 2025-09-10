package authhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "diawise/internal/api/middleware"
	repository "diawise/internal/repository/userRegistrationLoginRepository"
	authenticationservice "diawise/internal/services/authenticationService"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type RegistrationHandler struct {
	service *authenticationservice.Signup_service
}

func NewRegistrationHandler(service *authenticationservice.Signup_service) *RegistrationHandler {
	return &RegistrationHandler{service: service}
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
