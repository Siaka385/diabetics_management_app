package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	auth "diawise/internal/middleware"
	"diawise/internal/models"
	"diawise/internal/repository"
	"diawise/internal/services"
	"diawise/internal/shared"
)

func LogMealHandler(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Starting meal logging...")
		// Get user id to use as foreignkey
		// Retrieve the JWT token from cookies
		cookie, err := r.Cookie("authToken")
		if err != nil || cookie == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Parse and validate the JWT token
		tokenString := cookie.Value
		user, err := auth.ParseToken(tokenString)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		userID := user.ID

		var mealEntry services.MealLogEntry
		mealEntry.UserID = userID
		// Decode the request body into the new struct
		err = json.NewDecoder(r.Body).Decode(&mealEntry)
		if err != nil {
			log.Printf("Invalid input: %+v\n", err)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Find or create daily meal log for the current day
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

		var dailyMealLog models.DailyMealLog
		result := db.Where("user_id = ? AND date = ?", userID, today).FirstOrCreate(&dailyMealLog, models.DailyMealLog{UserID: userID, Date: today})
		if result.Error != nil {
			log.Printf("Failed to create/find daily meal log: %+v\n", err)
			http.Error(w, "Failed to create/find daily meal log", http.StatusInternalServerError)
			return
		}

		// Add meal entry to daily meal log
		analyser, err := NewAIHealthAnalyser()
		if err != nil {
			log.Printf("Failed to create AI health analyser: %v", err)
			http.Error(w, "Failed to create AI health analyser", http.StatusInternalServerError)
			return
		}
		dietProfileModel, err := analyser.DietProfile(&mealEntry)
		if err != nil {
			log.Printf("Failed to generate diet profile: %v", err)
			http.Error(w, "Failed to generate diet profile", http.StatusInternalServerError)
			return
		}
		defer analyser.Close()

		dietProfileModel.UserID = userID

		// Convert *models.DietProfile to services.DietProfile
		dietProfile := services.DietProfile{
			UserID:            dietProfileModel.UserID,
			FoodName:          dietProfileModel.FoodName,
			CaloriesIntake:    dietProfileModel.Calories,
			CarbIntake:        dietProfileModel.Carbs,
			ProteinIntake:     dietProfileModel.Protein,
			FatIntake:         dietProfileModel.Fat,
			SugarConsumption:  0,
			WaterIntake:       0,
			ProcessedFoodRatio: 0,
		}

		err = services.SaveDietLog(db, dietProfile)
		if err != nil {
			log.Printf("Failed to save diet profile: %v", err)
			http.Error(w, "Failed to save diet profile", http.StatusInternalServerError)
			return
		}

		mealEntry.DailyMealLogID = dailyMealLog.ID
		mealEntry.DietProfileID = dietProfile.ID

		err = services.SaveMealLog(db, mealEntry)
		if err != nil {
			http.Error(w, "Failed to save meal log", http.StatusInternalServerError)
			return
		}

		// Convert services.MealLogEntry to models.MealLogEntry for dailyMealLog
		mealEntryModel := models.MealLogEntry{
			UserID:         mealEntry.UserID,
			DailyMealLogID: mealEntry.DailyMealLogID,
			DietProfileID:  mealEntry.DietProfileID,
			FoodName:       mealEntry.FoodItem,
			Quantity:       mealEntry.Weight,
			Calories:       dietProfileModel.Calories,
			Carbs:          dietProfileModel.Carbs,
			Protein:        dietProfileModel.Protein,
			Fat:            dietProfileModel.Fat,
			Date:           time.Now(),
		}

		dailyMealLog.Entries = append(dailyMealLog.Entries, mealEntryModel)
		db.Save(&dailyMealLog)

		// Send JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			DietProfile models.DietProfile
		}{
			DietProfile: *dietProfileModel,
		})
	}
}

func EditPlan(w http.ResponseWriter, r *http.Request) {
	var updates models.FoodLog
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
			return
		}

		UserProfileDetails := models.UserProfile{
			Name:   user.Name,
			Abbrev: shared.GenerateShortName(user.Name),
		}
		Data.Profile = UserProfileDetails
		if err := tmpl.ExecuteTemplate(w, "blog_home.html", Data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func BloodSugarHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		UserProfileDetails := models.UserProfile{
			Name:   user.Name,
			Abbrev: shared.GenerateShortName(user.Name),
		}

		if err := tmpl.ExecuteTemplate(w, "bloodsugar.html", UserProfileDetails); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func MedicationHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "medication.html", Data); err != nil {
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
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		UserProfileDetails := models.UserProfile{
			Name:   user.Name,
			Abbrev: shared.GenerateShortName(user.Name),
		}
		if err := tmpl.ExecuteTemplate(w, "DietAndNutrition.html", UserProfileDetails); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func CommuniyAndSupportHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve user from context
		user, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		UserProfileDetails := models.UserProfile{
			Name:   user.Name,
			Abbrev: shared.GenerateShortName(user.Name),
		}

		if err := tmpl.ExecuteTemplate(w, "CommunityAndSupport.html", UserProfileDetails); err != nil {
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
				"status": "error",
				"message": "All fields are required",
			})
			return
		}
		
		success := repository.RegisterUser(db, signupData.Username, signupData.Username, signupData.Email, signupData.Password)
		if !success {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status": "error",
				"message": "Failed to register user",
			})
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "success",
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
			"status": "error",
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
			"status": "success",
			"message": "Login successful",
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
				"id":   user.ID,
				"name": user.Name,
				"email": user.Email,
			},
		})
	}
}
