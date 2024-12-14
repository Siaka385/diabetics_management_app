package api

import (
	"fmt"
	"net/http"

	"diawise/internal/services"

	"gorm.io/gorm"
)

func Message(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20) // 10MB maximum size
		if err != nil {
			http.Error(w, "Error parsing multipart form data", http.StatusBadRequest)
			return
		}

		// extract the message from the form
		message := r.FormValue("message")
		if message == "" {
			http.Error(w, "Message cannot be empty", http.StatusBadRequest)
			return
		}

		fmt.Println("Message: ", message)
		if message == "" {
			http.Error(w, "Message cannot be empty", http.StatusBadRequest)
			return
		}

		// broadcast the message to all connected clients via SSE
		services.Broadcast(message)

		// send a response to the client without refreshing the page
		w.Write([]byte("Message sent"))
	}
}

func SSEvents(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Create a channel for this client
		clientChan := make(chan services.SSEvent)
		services.Register(clientChan)

		// Flusher to push data to client immediately
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		// Serve events to this client
		go func() {
			for event := range clientChan {
				// Send event data
				fmt.Fprintf(w, "data: %s\n\n", event.Data)
				flusher.Flush()
			}
		}()

		// Keep the connection open
		select {}
	}
}
