package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	handlers "diawise/src/api"
	database "diawise/src/database"
	support "diawise/src/services/support"
	utils "diawise/src/utils"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

var (
	db           *gorm.DB // since sqlite is an internal database that is file based, we need to  have a single handler to the database. Use mutexes to prevent race conditions
	tmpl         *template.Template
	err          error
	sessionStore *sessions.CookieStore
)

func init() {
	db = database.InitializeDatabase("data/diawise.db")
	// parse all html files in the frontend and its subdirectories beforehand // optimization
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	// initialize all the structures needed for the support
	// includes the mutex to help coordinate the clients' map
	support.Init()

	// sessions and cookies
	secret := utils.GenerateRandomString(32)
	sessionStore = sessions.NewCookieStore([]byte(secret))

	sessionStore.Options = &sessions.Options{
		Path:     "/dashboard",
		MaxAge:   3600,  // expiration time in seconds
		HttpOnly: true,  // the cookie should be only accessible by HTTP(S)
		Secure:   false, // set to true in production to use with HTTPS
	}
}

func main() {
	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index(db, tmpl)).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/nutrition/meal/log", handlers.LogMealHandler(db)).Methods("POST")
	router.HandleFunc("/nutrition/mealplan", handlers.GetMealPlan).Methods("POST")
	// router.HandleFunc("/nutrition/editplan", api.EditPlan).Methods("POST")
	// router.HandleFunc("/nutrition/suggestions", api.GetMealSuggestions).Methods("POST")
	router.HandleFunc("/signup", handlers.Signup(db, tmpl, sessionStore)).Methods("GET")
	router.HandleFunc("/auth/signup", handlers.SignupUser(db, sessionStore)).Methods("POST")
	router.HandleFunc("/auth/login", handlers.LoginUser(db, sessionStore)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(db, tmpl, sessionStore)).Methods("GET")
	router.HandleFunc("/logout", handlers.Logout(sessionStore)).Methods("GET")
	router.HandleFunc("/dashboard", handlers.Dashboard(db, tmpl, sessionStore)).Methods("GET")
	router.HandleFunc("/support", handlers.Support(db, tmpl)).Methods("GET")
	router.HandleFunc("/medication", handlers.MedicationPageHandler(db, tmpl)).Methods("GET")
	router.HandleFunc("/addmed", handlers.AddMedication(db)).Methods("POST")
	router.HandleFunc("/updatemed/{id}", handlers.UpdateMedication(db)).Methods("PUT")
	router.HandleFunc("/deletemed/{id}", handlers.DeleteMedication(db)).Methods("DELETE")
	router.HandleFunc("/listmed", handlers.ListMedications(db)).Methods("GET")
	router.HandleFunc("/api/support/message", handlers.Message(db)).Methods("POST")
	router.HandleFunc("/api/support/events", handlers.SSEvents(db)).Methods("GET")
	router.HandleFunc("/blog", handlers.BlogHomeHandler(tmpl)).Methods("GET")
	router.HandleFunc("/glucose-tracker", handlers.GlucoseTrackerEndPointHandler).Methods("GET")
	router.HandleFunc("/post/{id}", handlers.PostHandler(tmpl)).Methods("GET")

	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handlerWithCORS := corsHandler.Handler(router) // apply the CORS middleware to the router

	http.ListenAndServe(portStr, handlerWithCORS)
}