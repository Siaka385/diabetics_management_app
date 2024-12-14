package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	handlers "diawise/internal/api"
	database "diawise/internal/database"
	support "diawise/internal/services/support"
	utils "diawise/pkg"

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
	tmpl, err = template.ParseGlob("../../frontend/**/*.html")
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
	router.HandleFunc("/education", handlers.Education(db, tmpl)).Methods("GET")
	router.HandleFunc("/track", handlers.BloodSugar(db, tmpl)).Methods("GET")
	router.HandleFunc("/nutrition", handlers.DietAndNutrient(db, tmpl)).Methods("GET")
	router.HandleFunc("/addmedication", handlers.DietAndNutrient(db, tmpl)).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../frontend/src"))))
	router.HandleFunc("/signup", handlers.Signup(db, tmpl, sessionStore)).Methods("GET")
	router.HandleFunc("/auth/signup", handlers.SignupUser(db, sessionStore)).Methods("POST")
	router.HandleFunc("/auth/login", handlers.LoginUser(db, sessionStore)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(db, tmpl, sessionStore)).Methods("GET")
	router.HandleFunc("/logout", handlers.Logout(sessionStore)).Methods("GET")
	router.HandleFunc("/dashboard", handlers.Dashboard(db, tmpl, sessionStore)).Methods("GET")
	router.HandleFunc("/support", handlers.Support(db, tmpl)).Methods("GET")
	router.HandleFunc("/addmed", handlers.AddMedication(db)).Methods("POST")
	router.HandleFunc("/updatemed/{id}", handlers.UpdateMedication(db)).Methods("PUT")
	router.HandleFunc("/deletemed/{id}", handlers.DeleteMedication(db)).Methods("DELETE")
	router.HandleFunc("/listmed", handlers.ListMedications(db)).Methods("GET")
	router.HandleFunc("/api/support/message", handlers.Message(db)).Methods("POST")
	router.HandleFunc("/api/support/events", handlers.SSEvents(db)).Methods("GET")

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
