package main

import (
	"fmt"
	"net/http"

	"diawise/internal/api"
	utils "diawise/pkg"

	"github.com/gorilla/mux"
)

func main() {
	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := mux.NewRouter()

	//router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/blog", api.BlogHomeHandler).Methods("GET")
	router.HandleFunc("/glucose-tracker", api.GlucoseTrackerEndPointHandler).Methods("GET")

	http.ListenAndServe(portStr, router)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
