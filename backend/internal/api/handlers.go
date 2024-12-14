package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"diawise/internal/services"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

func GetMealSuggestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting meal suggestions...")
}

func LogMealHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mealEntry services.MealLogEntry
		fmt.Println("Here")
		// Decode the request body into the new struct
		err := json.NewDecoder(r.Body).Decode(&mealEntry)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		err = services.SaveMealLog(db, mealEntry)
		if err != nil {
			http.Error(w, "Failed to save meal log", http.StatusInternalServerError)
			return
		}
		fmt.Println("here", mealEntry)

		// Prepare response
		response := NutritionResponse{
			Message:      "Meal logged successfully!",
			MealInsights: "Looks good",
		}

		// Send JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func EditPlan(w http.ResponseWriter, r *http.Request) {
	var updates FoodLog
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	defaultMealPlan = updates

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Meal plan updated successfully"})
}

func Index(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templateName := "index.html"

		err := tmpl.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			InternalServerErrorHandler(w)
		}
	}
}

func GlucoseTrackerEndPointHandler(w http.ResponseWriter, r *http.Request) {
	// Capture glucose level and date from the request query parameters
	glucoseLevel := r.URL.Query().Get("glucose")
	glucoseDate := r.URL.Query().Get("date")

	glucoseParam := map[string]string{glucoseLevel: glucoseDate}

	// Set response header and JSON encode the glucose level and date
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(glucoseParam)
}

func PostHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["id"]

		post, ok := Posts[postID]
		if !ok {
			NotFoundHandler(w)
			return
		}

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
		if err := tmpl.ExecuteTemplate(w, "blog_home.html", Data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func BadRequestHandler(w http.ResponseWriter) {
	tmpl := LoadTemplate()

	Hitch.StatusCode = http.StatusBadRequest
	Hitch.Problem = "Bad Request!"

	err := tmpl.Execute(w, Hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func InternalServerErrorHandler(w http.ResponseWriter) {
	// Check if headers have already been written
	if w.Header().Get("Content-Type") != "" {
		log.Println("Headers already written. Cannot send error page.")
		return
	}

	tmpl := LoadTemplate()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	Hitch.StatusCode = http.StatusInternalServerError
	Hitch.Problem = "Internal Server Error!"

	err := tmpl.Execute(w, Hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}

func NotFoundHandler(w http.ResponseWriter) {
	tmpl := LoadTemplate()

	Hitch.StatusCode = http.StatusNotFound
	Hitch.Problem = "Not Found!"

	err := tmpl.Execute(w, Hitch)
	if err != nil {
		http.Error(w, "Could not execute error template, error page unavailable", http.StatusInternalServerError)
		log.Println("Error executing template: ", err)
	}
}
