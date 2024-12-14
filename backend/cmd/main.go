package main

import (
	"fmt"
	"net/http"

	handlers "diawise/internal/api"
	database "diawise/internal/database"
	utils "diawise/pkg"
	"diawise/internal/api"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

var db *gorm.DB // since sqlite is an internal database that is file based, we need to  have a single handler to the database. Use mutexes to prevent race conditions

func init() {
	db = database.InitializeDatabase("data/diawise.db")
}

func main() {
	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index(db)).Methods("GET")
	router.HandleFunc("/auth/register", handlers.RegisterUser(db)).Methods("POST")
	router.HandleFunc("/auth/login", handlers.LoginUser(db)).Methods("POST")

	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handlerWithCORS := corsHandler.Handler(router) // apply the CORS middleware to the router

	http.ListenAndServe(portStr, handlerWithCORS)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/nutrition/mealplan", api.GetDefaultMealPlan).Methods("GET")
	router.HandleFunc("/nutrition/editplan", api.EditPlan).Methods("POST")
	router.HandleFunc("/api/nutrition/meal/log", api.LogMealHandler).Methods("POST")
	router.HandleFunc("/nutrition/suggestions", api.GetMealSuggestions).Methods("POST")

	http.ListenAndServe(portStr, router)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
