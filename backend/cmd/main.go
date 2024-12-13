package main

import (
	"fmt"
	"net/http"

	utils "diawise/pkg"
	"diawise/internal/api"

	"github.com/gorilla/mux"
)

func main() {
	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := mux.NewRouter()

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
