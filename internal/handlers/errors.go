package handlers

import (
	"log"
	"net/http"
)

func BadRequestHandler(w http.ResponseWriter) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func InternalServerErrorHandler(w http.ResponseWriter) {
	log.Println("Internal Server Error occurred")
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func NotFoundHandler(w http.ResponseWriter) {
	http.Error(w, "Not Found", http.StatusNotFound)
}
