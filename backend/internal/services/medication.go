package services

import (
	"errors"
	"fmt"
	"time"

	"diawise/internal/database"
)

type Users struct {
	User_id  string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Medication struct {
	Medication_id    string    `json:"medication_id"`
	User_id          string    `json:"user_id"`
	Medication_name  string    `json:"medication_name"`
	Dose             string    `json:"dose"`
	Dosage_time      time.Time `json:"time"`
	Dosage_frequency string    `json:"frequency"`
	Notes            string    `json:"notes"`
}

type MedicationService struct{}

// new medication creates a new medication
func NewMedicationService() *MedicationService {
	return &MedicationService{}
}

// Adding medication to the database
func (s *MedicationService) AddMedication(medication Medication) (Medication, error) {
	if medication.Medication_id == "" || medication.User_id == "" || medication.Medication_name == "" || medication.Dose == "" || medication.Dosage_time.IsZero() || medication.Dosage_frequency == "" || medication.Notes == "" {
		return Medication{}, errors.New("missing required fields")
	}
	// Add medication to the database
	meds := Medication{
		Medication_id:    medication.Medication_id,
		User_id:          medication.User_id,
		Medication_name:  medication.Medication_name,
		Dose:             medication.Dose,
		Dosage_time:      medication.Dosage_time,
		Dosage_frequency: medication.Dosage_frequency,
		Notes:            medication.Notes,
	}

	// Handle any errors during database operations
	_, err := database.DB.Exec( // SQL query to insert a new medication into the database
		"INSERT INTO medications (medication_id, user_id, medication_name, medication_dose, dosage_time, dosage_frequency, notes) VALUES (?,?,?,?,?,?,?)",
		meds.Medication_id, meds.User_id, meds.Medication_name, meds.Dose, meds.Dosage_time, meds.Dosage_frequency, meds.Notes,
	)
	if err != nil {
		return Medication{}, fmt.Errorf("failed to add medication: %v", err)
	}

	return meds, nil
}

// Getting medications by user_id
func (s *MedicationService) GetMedicationsByUserId(userID string) ([]Medication, error) {
	rows, err := database.DB.Query("SELECT * FROM medications WHERE user_id=?", userID) // SQL query to select all medications for a given user
	if err != nil {
		return nil, fmt.Errorf("failed to get medications: %v", err)
	}
	defer rows.Close()

	var medications []Medication
	for rows.Next() {
		var med Medication
		err := rows.Scan(&med.Medication_id, &med.User_id, &med.Medication_name, &med.Dose, &med.Dosage_time, &med.Dosage_frequency, &med.Notes)
		if err != nil {
			return nil, fmt.Errorf("failed to scan medication row: %v", err)
		}
		medications = append(medications, med)
	}

	return medications, nil
}

// update medications by medication_id
func (s *MedicationService) UpdateMedication(medication Medication) (Medication, error) {
	if medication.Medication_id == "" || medication.User_id == "" || medication.Medication_name == "" || medication.Dose == "" || medication.Dosage_time.IsZero() || medication.Dosage_frequency == "" || medication.Notes == "" {
		return Medication{}, errors.New("missing required fields")
	}

	_, err := database.DB.Exec( // SQL query to update a medication in the database
		"UPDATE medications SET medication_name=?, medication_dose=?, dosage_time=?, dosage_frequency=?, notes=? WHERE medication_id=? AND user_id=?",
		medication.Medication_name, medication.Dose, medication.Dosage_time, medication.Dosage_frequency, medication.Notes,
	)
	if err != nil {
		return Medication{}, fmt.Errorf("failed to update medication: %v", err)
	}

	var med Medication
	err = database.DB.QueryRow(
		"SELECT * FROM medications WHERE medication_id=? AND user_id=?", medication.Medication_id, medication.User_id,
	).Scan(&med.Medication_id, &med.User_id, &med.Medication_name, &med.Dose, &med.Dosage_time, &med.Dosage_frequency, &med.Notes)
	if err != nil {
		return Medication{}, fmt.Errorf("failed to get updated medication: %v", err)
	}
	return med, nil
}
