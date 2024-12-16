package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"diawise/src/services"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"gorm.io/gorm"
)

func GetMealSuggestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting meal suggestions...")
}

func LogMealHandler(db *gorm.DB, tmpl *template.Template, ss *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Starting meal logging...")
		// Get user id to use as foreignkey
		vars := mux.Vars(r)
		userID := vars["id"]
		// session, err := ss.Get(r, "session-name")
        // if err!= nil {
		// 	log.Printf("Error retrieving session: %+v\n", err)
        //     http.Error(w, "Error retrieving session", http.StatusInternalServerError)
        //     return
        // }
        // userID, ok := session.Values["user_id"].(string)
		// fmt.Println(userID)
        // if !ok {
		// 	log.Printf("User not authenticated: %+v\n", err)
        //     http.Error(w, "User not authenticated", http.StatusUnauthorized)
        //     return
        // }
		var mealEntry services.MealLogEntry
		mealEntry.UserID = userID
		// Decode the request body into the new struct
		err := json.NewDecoder(r.Body).Decode(&mealEntry)
		if err != nil {
			log.Printf("Invalid input: %+v\n", err)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		// debug meal entry
		log.Printf("Created meal entry: %+v\n", mealEntry)

		// Find or create daily meal log for the current day
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0,0,0,0, time.UTC)

		var dailyMealLog services.DailyMealLog
		result := db.Where("user_id = ? AND date = ?", userID, today).FirstOrCreate(&dailyMealLog, services.DailyMealLog{UserID: userID, Date: today})
		if result.Error!= nil {
			log.Printf("Failed to create/find daily meal log: %+v\n", err)
            http.Error(w, "Failed to create/find daily meal log", http.StatusInternalServerError)
            return
        }

		log.Printf("Logged daily meal log: %+v\n", dailyMealLog)

		// Add meal entry to daily meal log
		analyser, err := NewAIHealthAnalyser()
		if err!= nil {
            log.Printf("Failed to create AI health analyser: %v", err)
            http.Error(w, "Failed to create AI health analyser", http.StatusInternalServerError)
            return
        }
		dietProfile, err := analyser.DietProfile(&mealEntry)
		if err!= nil {
            log.Printf("Failed to generate diet profile: %v", err)
            http.Error(w, "Failed to generate diet profile", http.StatusInternalServerError)
            return
        }
		defer analyser.Close()

		log.Printf("Got diet profile from genai client: %+v\n", dietProfile)

		dietProfile.UserID = userID
		err = services.SaveDietLog(db, *dietProfile)
		if err!= nil {
            log.Printf("Failed to save diet profile: %v", err)
            http.Error(w, "Failed to save diet profile", http.StatusInternalServerError)
            return
        }

		log.Printf("Saved diet profile log: %+v\n", dietProfile)
		mealEntry.DailyMealLogID = dailyMealLog.ID
		mealEntry.DietProfileID = dietProfile.ID
		err = services.SaveMealLog(db, mealEntry)
		if err!= nil {
            http.Error(w, "Failed to save meal log", http.StatusInternalServerError)
            return
        }
		dailyMealLog.Entries = append(dailyMealLog.Entries, mealEntry)
		db.Save(&dailyMealLog)

		// Send JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct{Message string; MealInsights string}{
			Message:      "Meal logged successfully!",
			MealInsights: "Looks good",
		})
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

func BloodSugarHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "bloodsugar.html", Data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func EducationHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "Education.html", Data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func DietAndNutritionHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "DietAndNutrition.html", Data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func CommuniyAndSupportHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "CommunityAndSupport.html", Data); err != nil {
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
