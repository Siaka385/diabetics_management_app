package main

import (
	"fmt"
	"html/template"
	"log"
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
	tmpl         *template.Template
	err          error
	sessionStore *sessions.CookieStore
)

func init() {
	db = repository.InitializeDatabase("data/diawise.db")
	tmpl, err = template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	// Parse partials
	tmpl, err = tmpl.ParseGlob("web/templates/partials/*.html")
	if err != nil {
		log.Fatal(err)
	}
	sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))
}

func main() {
	port := shared.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index(db, tmpl)).Methods("GET")
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

	router.HandleFunc("/nutrition/logmeal", handlers.LogMealHandler(db, tmpl)).Methods("POST")
	router.Handle("/medication", http.HandlerFunc(middleware.AuthMiddleware(handlers.MedicationHandler(tmpl)))).Methods("GET")
	router.HandleFunc("/updatemed/{id}", handlers.UpdateMedication(db)).Methods("PUT")
	router.HandleFunc("/deletemed/{id}", handlers.DeleteMedication(db)).Methods("DELETE")
	router.HandleFunc("/listmed", handlers.ListMedications(db)).Methods("GET")
	router.HandleFunc("/education", handlers.EducationHandler(tmpl)).Methods("GET")
	router.HandleFunc("/glucose-tracker", handlers.GlucoseTrackerEndPointHandler).Methods("GET")
	router.HandleFunc("/post/{id}", handlers.PostHandler(tmpl)).Methods("GET")

	router.HandleFunc("/createroom", handlers.CreateRoom(db)).Methods("POST")
	router.HandleFunc("/listrooms", handlers.ListRooms(db)).Methods("GET")
	router.HandleFunc("/joinroom", handlers.JoinRoom(db))
	router.HandleFunc("/sendmessage", handlers.SendMessage)
	router.HandleFunc("/deleteroom", handlers.DeleteRoom(db))

	// Restricted routes
	router.Handle("/dashboard", http.HandlerFunc(middleware.AuthMiddleware(handlers.Dashboard(db)))).Methods("GET")
	router.Handle("/support", http.HandlerFunc(middleware.AuthMiddleware(handlers.Support(tmpl)))).Methods("GET")
	router.Handle("/nutrition", http.HandlerFunc(middleware.AuthMiddleware(handlers.DietAndNutritionHandler(tmpl)))).Methods("GET")
	router.Handle("/bloodsugar", http.HandlerFunc(middleware.AuthMiddleware(handlers.BloodSugarHandler(tmpl)))).Methods("GET")
	router.Handle("/blog", http.HandlerFunc(middleware.AuthMiddleware(handlers.BlogHomeHandler(tmpl)))).Methods("GET")
	router.Handle("/addmedication", http.HandlerFunc(middleware.AuthMiddleware(handlers.AddMedicationHandler(db, tmpl)))).Methods("GET", "POST")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handlerWithCORS := corsHandler.Handler(router)

	http.ListenAndServe(portStr, handlerWithCORS)
}
