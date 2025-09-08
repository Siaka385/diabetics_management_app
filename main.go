package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"gorm.io/gorm"

	"diawise/internal/handlers"
	"diawise/internal/middleware"
	"diawise/internal/repository"
	"diawise/internal/shared"
)

var (
	db           *gorm.DB
	err          error
	sessionStore *sessions.CookieStore
)

func init() {
	db = repository.InitializeDatabase("data/diawise.db")
	sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))
}

func main() {
	port := shared.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index()).Methods("GET")
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

	router.HandleFunc("/auth/signup", handlers.Signup(db)).Methods("POST")
	router.HandleFunc("/auth/signin", handlers.Signin(db, sessionStore)).Methods("POST")
	router.HandleFunc("/auth/signout", handlers.Signout).Methods("POST")
	router.HandleFunc("/auth/status", handlers.AuthStatus).Methods("GET")

	router.HandleFunc("/nutrition/logmeal", handlers.LogMealHandler(db)).Methods("POST")
	router.Handle("/medication", http.HandlerFunc(middleware.AuthMiddleware(handlers.MedicationHandler()))).Methods("GET")
	router.HandleFunc("/updatemed/{id}", handlers.UpdateMedication(db)).Methods("PUT")
	router.HandleFunc("/deletemed/{id}", handlers.DeleteMedication(db)).Methods("DELETE")
	router.HandleFunc("/listmed", handlers.ListMedications(db)).Methods("GET")
	router.Handle("/education", http.HandlerFunc(middleware.AuthMiddleware(handlers.EducationHandler()))).Methods("GET")
	router.Handle("/glucose-tracker", http.HandlerFunc(middleware.AuthMiddleware(handlers.GlucoseTrackerHandler()))).Methods("GET")
	router.Handle("/post/{id}", http.HandlerFunc(middleware.AuthMiddleware(handlers.PostHandler()))).Methods("GET")

	router.HandleFunc("/createroom", handlers.CreateRoom(db)).Methods("POST")
	router.HandleFunc("/listrooms", handlers.ListRooms(db)).Methods("GET")
	router.HandleFunc("/joinroom", handlers.JoinRoom(db))
	router.HandleFunc("/sendmessage", handlers.SendMessage)
	router.HandleFunc("/deleteroom", handlers.DeleteRoom(db))

	// Restricted routes
	router.Handle("/dashboard", http.HandlerFunc(middleware.AuthMiddleware(handlers.Dashboard(db)))).Methods("GET")
	router.Handle("/support", http.HandlerFunc(middleware.AuthMiddleware(handlers.Support()))).Methods("GET")
	router.Handle("/nutrition", http.HandlerFunc(middleware.AuthMiddleware(handlers.DietAndNutritionHandler()))).Methods("GET")
	router.Handle("/bloodsugar", http.HandlerFunc(middleware.AuthMiddleware(handlers.BloodSugarHandler()))).Methods("GET")
	router.Handle("/blog", http.HandlerFunc(middleware.AuthMiddleware(handlers.BlogHomeHandler()))).Methods("GET")
	router.HandleFunc("/addmedication", handlers.AddMedication(db, sessionStore)).Methods("POST")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handlerWithCORS := corsHandler.Handler(router)

	http.ListenAndServe(portStr, handlerWithCORS)
}
