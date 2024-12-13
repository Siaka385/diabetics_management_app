package main

import (
	"fmt"
	"net/http"

	utils "diawise/pkg"

	"github.com/gorilla/mux"
)

func main() {
	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	router := mux.NewRouter()

	router.HandleFunc("/", Index).Methods("GET")

	http.ListenAndServe(portStr, router)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
