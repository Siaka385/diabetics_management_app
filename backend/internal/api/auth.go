package api

import (
	"encoding/json"
	"net/http"

	auth "diawise/internal/auth"

	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var response map[string]string

		var user auth.User
		// Decode the incoming JSON to a Piece struct
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err := auth.RegisterUser(db, user.Username, user.Email, user.Password)
		if !err {
			response = map[string]string{"status": "error", "message": "unable to register user"}
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			response = map[string]string{"status": "success", "message": "User registered successfully"}
		}

		// Return a 201 Created status
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var response map[string]string

		// Decode the incoming JSON body
		var userInput struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// Parse the JSON request body
		if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Try to find the user in the database
		user, err := auth.LoginUser(db, userInput.Username, userInput.Password)
		if err != nil {
			response = map[string]string{"status": "error", "message": "unable to login user"}
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			response = map[string]string{"status": "success", "message": "Login successful", "username": user.Username}

		}

		json.NewEncoder(w).Encode(response)
	}
}
