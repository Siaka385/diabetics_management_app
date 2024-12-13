package services

import (
	"errors"
	"time"
)

type Users struct {
	user_id  string `json:"user_id"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}

type Medication struct {
	medication_id    string    `json:"medication_id"`
	user_id          string    `json:"user_id"`
	medication_name  string    `json:"medication_name"`
	dose             string    `json:"dose"`
	dosage_time      time.Time `json:"time"`
	dosage_frequency string    `json:"frequency"`
	notes            string    `json:"notes"`
}

type MedicationService struct{}

// new medication creates a new medication
func NewMedicationService() *MedicationService {
	return &MedicationService{}
}

// Adding medication to the database
func (s *MedicationService) AddMedication(medication Medication) (Medication, error) {
	if medication.medication_name == "" || medication.dose == "" || medication.dosage_time.IsZero() || medication.dosage_frequency == "" {
		return Medication{}, errors.New("Missing required fields")
	}
	// Add medication to the database
	meds := Medication{
		medication_id:    medication.medication_id,
		user_id:          medication.user_id,
		medication_name:  medication.medication_name,
		dose:             medication.dose,
		dosage_time:      medication.dosage_time,
		dosage_frequency: medication.dosage_frequency,
		notes:            medication.notes,
	}

	return meds, nil
}
