package authhandler

import (
	"encoding/json"
	"net/http"
)

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
