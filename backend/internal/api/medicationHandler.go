package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"diawise/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Medication struct {
	Medication   services.Medication
	ReminderTime time.Duration `json:"reminder_time"`
}

func MedicationPageHandler(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement logic to fetch medications and render the medication page
		if err := tmpl.ExecuteTemplate(w, "medication.html", nil); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newMed)
	}
}

// handler for deleting a medication
func DeleteMedication(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		med := services.Medication{Medication_id: id}
		if err := services.DeleteMedication(db, med); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
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
		w.Header().Set("Content-Type", "application/json")
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

func AddMedicationHandler(db *gorm.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		medicationName := r.FormValue("medication_name")
		dose := r.FormValue("dose")
		frequency := r.FormValue("frequency")
		reminderTime := r.FormValue("reminder_time")

		dosageTime, err := time.Parse("15:04", reminderTime)
		if err != nil {
			http.Error(w, "Invalid reminder time format (HH:MM)", http.StatusBadRequest)
			return
		}

		// Create a new Medication struct
		newMedication := services.Medication{
			Medication_name:  medicationName,
			Dose:             dose,
			Dosage_frequency: frequency,
			Dosage_time:      dosageTime,
			// Set User_id and Medication_id as needed
			User_id:       "user123", // Replace with actual user ID from session
			Medication_id: "med_" + time.Now().Format("20060102150405"),
			Notes:         r.FormValue("notes"), // Add notes if available in the form
		}

		// Call the AddMedication function from the services package
		addedMedication, err := services.AddMedication(db, newMedication)
		if err != nil {
			http.Error(w, "Failed to add medication: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the added medication as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(addedMedication)
	}
}
