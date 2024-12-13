package main

import (
	"fmt"
	"net/http"

	utils "diawise/pkg"
	handlers "diawise/internal/api"
	database "diawise/internal/database"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	db *gorm.DB // since sqlite is an internal database that is file based, we need to  have a single handler to the database. Use mutexes to prevent race conditions
)

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

	http.ListenAndServe(portStr, router)
}

