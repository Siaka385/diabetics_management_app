package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"diawise/internal/services"

	"gorm.io/gorm"
)

type Medication struct {
	Medication   services.Medication
	ReminderTime time.Duration `json:"reminder_time"`
}

// handler for adding a new medication
func AddMedication(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var med services.Medication

		if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newMed, err := services.AddMedication(db, med)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newMed)
	}
}

// handler for deleting a medication
func DeleteMedication(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var med services.Medication

		if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := services.DeleteMedication(db, med)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
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

// handler for listing all medications
func ListMedications(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "user_id is required", http.StatusBadRequest)
			return
		}
		medications, err := services.ListMedicationsByUserId(db, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(medications)
	}
}

// handler for medication reminder
func MedicationReminder(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "user_id is required", http.StatusBadRequest)
			return
		}
		// Fixed: Convert userID to int64 before passing it to SendMedicationReminders
		userIDInt, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusBadRequest)
			return
		}
		err = services.SendMedicationReminders(db, userIDInt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
