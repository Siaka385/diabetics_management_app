package api

import (
	"encoding/json"
	"net/http"
)

func GlucoseTrackerEndPointHandler(w http.ResponseWriter, r *http.Request) {

	// Capture glucose level and date from the request query parameters
	glucoseLevel := r.URL.Query().Get("glucose")
	glucoseDate := r.URL.Query().Get("date")

	glucoseParam := map[string]string{glucoseLevel: glucoseDate}

	// Set response header and JSON encode the glucose level and date
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(glucoseParam)
}
