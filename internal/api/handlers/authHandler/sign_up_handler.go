package authhandler

import (
	"encoding/json"
	"net/http"

	auth "diawise/internal/api/middleware"
	repository "diawise/internal/repository/userRegistrationLoginRepository"
	authenticationservice "diawise/internal/services/authenticationService"

	"gorm.io/gorm"
)

type SignupHandler struct {
	service authenticationservice.Signup_service
}

func NewSignupHandler(service authenticationservice.Signup_service) *SignupHandler {
	return &SignupHandler{service: service}
}

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
