package api

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"gorm.io/gorm"

	auth "diawise/src/auth"
)

type UserProfile struct {
	Abbrev string
	Name   string
}

func Dashboard(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Log user details (optional)
		// fmt.Printf("Authenticated user: %+v\n", user.Name)
		// fmt.Printf("Authenticated user ID: %+v\n", user.ID)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		UserProfileDetails := UserProfile{
			Abbrev: GenerateShortName(user.Name),
			Name:   user.Name,
		}

		// Serve the dashboard page
		if err := tmpl.ExecuteTemplate(w, "dashboard.html", UserProfileDetails); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func GenerateShortName(fullName string) string {
	name := strings.TrimSpace(fullName)
	words := strings.Fields(name)
	if len(words) == 0 {
		return ""
	}
	if len(words) == 1 {
		if len(words[0]) > 0 {
			return string(words[0][0])
		}
		return ""
	}
	shortName := ""
	for _, word := range words {
		if len(word) > 0 {
			shortName += string(word[0])
		}
	}
	return shortName
}
