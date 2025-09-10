package api

import (
	"net/http"

	"diawise/internal/api/middleware"

	// Handler imports
	authhandler "diawise/internal/api/handlers/authHandler"
	bloghandler "diawise/internal/api/handlers/blogHandler"
	bloodsugarhandler "diawise/internal/api/handlers/bloodSugarHandler"
	dashboardhandler "diawise/internal/api/handlers/dashboardHandler"
	educationhandler "diawise/internal/api/handlers/educationHandler"
	indexhandler "diawise/internal/api/handlers/indexPageHandler"
	medicationhandler "diawise/internal/api/handlers/medicationHandler"
	nutritionhandler "diawise/internal/api/handlers/nutritionHandler"
	supporthandler "diawise/internal/api/handlers/supportGroupHandler"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) http.Handler {
	router := mux.NewRouter()

	// Initialize session store
	sessionStore := sessions.NewCookieStore([]byte("your-secret-key"))

	router.HandleFunc("/", indexhandler.Index()).Methods("GET")
	// Custom static file handler with proper MIME types
	staticHandler := http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path[len(path)-4:] == ".css" {
			w.Header().Set("Content-Type", "text/css")
		} else if path[len(path)-3:] == ".js" {
			w.Header().Set("Content-Type", "application/javascript")
		}
		http.FileServer(http.Dir("web/static")).ServeHTTP(w, r)
	}))

	router.PathPrefix("/static/").Handler(staticHandler)

	router.HandleFunc("/auth/signup", authhandler.Signup(db)).Methods("POST")
	router.HandleFunc("/auth/signin", authhandler.Signin(db, sessionStore)).Methods("POST")
	router.HandleFunc("/auth/signout", authhandler.Signout).Methods("POST")
	router.HandleFunc("/auth/status", authhandler.AuthStatus).Methods("GET")

	router.HandleFunc("/nutrition/logmeal", nutritionhandler.LogMealHandler(db)).Methods("POST")
	router.Handle("/medication", http.HandlerFunc(middleware.AuthMiddleware(medicationhandler.MedicationHandler()))).Methods("GET")
	router.HandleFunc("/updatemed/{id}", medicationhandler.UpdateMedication(db)).Methods("PUT")
	router.HandleFunc("/deletemed/{id}", medicationhandler.DeleteMedication(db)).Methods("DELETE")
	router.HandleFunc("/listmed", medicationhandler.ListMedications(db)).Methods("GET")
	router.Handle("/education", http.HandlerFunc(middleware.AuthMiddleware(educationhandler.EducationHandler()))).Methods("GET")
	router.Handle("/glucose-tracker", http.HandlerFunc(middleware.AuthMiddleware(bloodsugarhandler.GlucoseTrackerHandler()))).Methods("GET")
	router.Handle("/post/{id}", http.HandlerFunc(middleware.AuthMiddleware(bloghandler.PostHandler()))).Methods("GET")

	router.HandleFunc("/createroom", supporthandler.CreateRoom(db)).Methods("POST")
	router.HandleFunc("/listrooms", supporthandler.ListRooms(db)).Methods("GET")
	router.HandleFunc("/joinroom", supporthandler.JoinRoom(db))
	router.HandleFunc("/sendmessage", supporthandler.SendMessage)
	router.HandleFunc("/deleteroom", supporthandler.DeleteRoom(db))

	// Restricted routes
	router.Handle("/dashboard", http.HandlerFunc(middleware.AuthMiddleware(dashboardhandler.Dashboard(db)))).Methods("GET")
	router.Handle("/support", http.HandlerFunc(middleware.AuthMiddleware(supporthandler.Support()))).Methods("GET")
	router.Handle("/nutrition", http.HandlerFunc(middleware.AuthMiddleware(bloodsugarhandler.DietAndNutritionHandler()))).Methods("GET")
	router.Handle("/bloodsugar", http.HandlerFunc(middleware.AuthMiddleware(bloodsugarhandler.BloodSugarHandler()))).Methods("GET")
	router.Handle("/blog", http.HandlerFunc(middleware.AuthMiddleware(bloghandler.BlogHomeHandler()))).Methods("GET")
	router.HandleFunc("/addmedication", medicationhandler.AddMedication(db, sessionStore)).Methods("POST")

	return router
}
