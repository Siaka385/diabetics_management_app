package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"diawise/internal/services"

	"gorm.io/gorm"
)

// handler for adding a new medication
func AddMedication(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var med services.Medication

		if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("MED: ", med)

		newMed, err := services.AddMedication(db, med)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newMed)
	}
}

// handler for updating an existing medication
func UpdateMedication(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var med services.Medication

		if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedMed, err := services.UpdateMedication(db, med)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedMed)
	}
}
