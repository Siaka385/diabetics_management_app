package api

import (
	"encoding/json"
	"net/http"
	"time"

	"diawise/internal/services"
)

type MedicationHandler struct {
	service *services.MedicationService
}

// handler for managing new medications
func NewMedicalHandler(services *services.MedicationService) *MedicationHandler {
	return &MedicationHandler{service: services}
}

// handler for adding a new medication
func (h *MedicationHandler) AddMedication(w http.ResponseWriter, r *http.Request) {
	var med struct {
		Medication_id    string    `json:"medication_id"`
		User_id          string    `json:"user_id"`
		Medication_name  string    `json:"medication_name"`
		Dose             string    `json:"dose"`
		Dosage_time      time.Time `json:"time"`
		Dosage_frequency string    `json:"frequency"`
		Notes            string    `json:"notes"`
	}

	if err := json.NewDecoder(r.Body).Decode(&med); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMed, err := h.service.AddMedication(services.Medication{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMed)
}
