package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"diawise/src/services"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type Medication struct {
	Medications  services.Medication
	ReminderTime time.Duration `json:"reminder_time"`
}

func MedicationPageHandler(db *gorm.DB, tmpl *template.Template, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Get the user_id from the session
		userID, ok := session.Values["user_id"].(string)
		fmt.Println("IDENTIFIER: ", userID)
		if !ok {
			session.AddFlash("Please log in to access the medication page.")
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if r.Method == http.MethodPost {
			// Handle adding new medication
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
				User_id:          userID,
				Medication_id:    "med_" + time.Now().Format("20060102150405"),
				Notes:            r.FormValue("notes"),
			}

			// Call the AddMedication function from the services package
			_, err = services.AddMedication(db, newMedication)
			if err != nil {
				http.Error(w, "Failed to add medication: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// Redirect to the medication page after successful addition
			http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
			return
		}

		// Fetch medications for display (GET request)
		medications, err := services.ListMedicationsByUserId(db, userID)
		if err != nil {
			http.Error(w, "Failed to fetch medications", http.StatusInternalServerError)
			return
		}

		data := struct {
			Medications []services.Medication
			Username    string
		}{
			Medications: medications,
			Username:    session.Values["username"].(string),
		}

		if err := tmpl.ExecuteTemplate(w, "medication.html", data); err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}

// handler for adding a new medication
func AddMedication(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Get the user_id from the session
		userID, ok := session.Values["user_id"].(string)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}

		var med services.Medication
		if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		med.User_id = userID
		result, err := services.AddMedication(db, med)
		if err != nil {
			http.Error(w, "Failed to add medication", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteMedication(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Get the user_id from the session
		userID, ok := session.Values["user_id"].(string)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		med := services.Medication{
			Medication_id: id,
			User_id:       userID,
		}

		if err := services.DeleteMedication(db, med); err != nil {
			http.Error(w, "Failed to delete medication", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateMedication(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Get the user_id from the session
		userID, ok := session.Values["user_id"].(string)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		var med services.Medication
		if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		med.Medication_id = id
		med.User_id = userID
		result, err := services.UpdateMedication(db, med)
		if err != nil {
			http.Error(w, "Failed to update medication", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func ListMedications(db *gorm.DB, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Get the user_id from the session
		userID, ok := session.Values["user_id"].(string)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}

		medications, err := services.ListMedicationsByUserId(db, userID)
		if err != nil {
			http.Error(w, "Failed to fetch medications", http.StatusInternalServerError)
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

func AddMedicationHandler(db *gorm.DB, tmpl *template.Template, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		// Check if user is authenticated
		auth, ok := session.Values["authenticated"].(bool)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		userID, userIDOk := session.Values["user_id"].(string)

		// Validate authentication
		if !ok || !auth || !userIDOk {
			fmt.Println("Redirecting to login: Not authenticated")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if r.Method == http.MethodGet {
			// Fetch existing medications for the user
			medications, err := services.ListMedicationsByUserId(db, userID)
			if err != nil {
				http.Error(w, "Failed to fetch medications", http.StatusInternalServerError)
				return
			}

			// Get the username from the session
			username, ok := session.Values["username"].(string)
			if !ok {
				username = "Unknown User" // Default in case username isn't found
			}

			// Prepare data for the template
			data := struct {
				Medications []services.Medication
				Username    string
			}{
				Medications: medications,
				Username:    username,
			}

			// Render the template
			if err := tmpl.ExecuteTemplate(w, "medication.html", data); err != nil {
				http.Error(w, "Failed to render template", http.StatusInternalServerError)
			}
			return
		}

		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			medicationName := r.FormValue("medication_name")
			dose := r.FormValue("dose")
			frequency := r.FormValue("frequency")
			reminderTime := r.FormValue("reminder_time")
			notes := r.FormValue("notes")

			// Ensure all required fields are filled out
			if medicationName == "" || dose == "" || frequency == "" || reminderTime == "" {
				http.Error(w, "All fields are required", http.StatusBadRequest)
				return
			}

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
				User_id:          userID,
				Medication_id:    "med_" + time.Now().Format("20060102150405"),
				Notes:            notes,
			}

			// Call the AddMedication function from the services package
			_, err = services.AddMedication(db, newMedication)
			if err != nil {
				http.Error(w, "Failed to add medication: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// Redirect to the medication page after successful addition
			http.Redirect(w, r, "/addmedication", http.StatusSeeOther)
			return
		}

		// If the method is neither GET nor POST, return an error
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
