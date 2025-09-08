package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	auth "diawise/internal/middleware"
	"diawise/internal/models"
	"diawise/internal/shared"

	"github.com/gorilla/mux"
)

func PostHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["id"]

		post, ok := Posts[postID]
		if !ok {
			NotFoundHandler(w)
			return
		}

		post.Abbrev = shared.GenerateShortName(post.Author)

		err := tmpl.ExecuteTemplate(w, "blog_display.html", post)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			InternalServerErrorHandler(w)
			return
		}
	}
}

func BlogHomeHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		UserProfileDetails := struct {
			models.UserProfile
			CurrentPage string
			Posts       []models.Post
		}{
			UserProfile: models.UserProfile{
				Name:   user.Name,
				Abbrev: shared.GenerateShortName(user.Name),
			},
			CurrentPage: "/blog",
			Posts:       Data.Posts,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserProfileDetails)
	}
}

func EducationHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "education.html", Data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

