package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"diawise/internal/api"
	handlers "diawise/internal/api"
	database "diawise/internal/database"
	support "diawise/internal/services/support"
	utils "diawise/pkg"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB // since sqlite is an internal database that is file based, we need to  have a single handler to the database. Use mutexes to prevent race conditions
	tmpl *template.Template
	err  error
)

func init() {
	db = database.InitializeDatabase("data/diawise.db")
	// parse all html files in the frontend and its subdirectories beforehand // optimization
	tmpl, err = template.ParseGlob("../../frontend/**/*.html")
	if err != nil {
		log.Fatal(err)
	}

	support.Init()
}

func main() {
	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index(db, tmpl)).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../frontend/src"))))
	router.HandleFunc("/auth/register", handlers.RegisterUser(db)).Methods("POST")
	router.HandleFunc("/auth/login", handlers.LoginUser(db)).Methods("POST")
	router.HandleFunc("/nutrition/meal/log", api.LogMealHandler(db)).Methods("POST")
	router.HandleFunc("/nutrition/mealplan", api.GetMealPlan).Methods("POST")
	// router.HandleFunc("/nutrition/editplan", api.EditPlan).Methods("POST")
	// router.HandleFunc("/nutrition/suggestions", api.GetMealSuggestions).Methods("POST")
	router.HandleFunc("/support", handlers.Support(db, tmpl)).Methods("GET")
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
