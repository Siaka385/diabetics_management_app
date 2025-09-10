package handlers

import (
	"encoding/json"
	"net/http"

	auth "diawise/internal/api/middleware"
	"diawise/internal/models"
	"diawise/internal/shared"

	"github.com/gorilla/mux"
)

func BlogHomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Use static posts data with proper IDs for navigation
		staticPosts := []models.Post{
			{
				ID:      "1",
				Title:   "Understanding Diabetes: A Comprehensive Guide",
				Author:  "Dr. Sarah Johnson",
				Date:    "2 days ago",
				Excerpt: "Learn the fundamentals of diabetes management, from blood sugar monitoring to lifestyle changes that can make a real difference.",
				Image:   "understanding-diabetes.jpg",
			},
			{
				ID:      "2",
				Title:   "10 Healthy Recipes for Diabetic-Friendly Meals",
				Author:  "Chef Maria Rodriguez",
				Date:    "1 week ago",
				Excerpt: "Delicious and nutritious meal ideas that help maintain stable blood sugar levels while satisfying your taste buds.",
				Image:   "prevent.jpg",
			},
			{
				ID:      "3",
				Title:   "Exercise and Diabetes: Finding the Right Balance",
				Author:  "Fitness Expert Mike Chen",
				Date:    "2 weeks ago",
				Excerpt: "Discover how regular physical activity can improve your diabetes management and overall health.",
				Image:   "workout.jpg",
			},
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
			Posts:       staticPosts,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserProfileDetails)
	}
}

func PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		vars := mux.Vars(r)
		postID := vars["id"]

		post, exists := Posts[postID]
		if !exists {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		response := struct {
			models.UserProfile
			CurrentPage string
			Post        models.Post
		}{
			UserProfile: models.UserProfile{
				Name:   user.Name,
				Abbrev: shared.GenerateShortName(user.Name),
			},
			CurrentPage: "/post/" + postID,
			Post:        post,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
