package services

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Medication struct {
	*gorm.Model
	Medication_id    string    `json:"medication_id"`
	User_id          string    `json:"user_id"`
	Medication_name  string    `json:"medication_name"`
	Dose             string    `json:"dose"`
	Dosage_time      time.Time `json:"time"`
	Dosage_frequency string    `json:"frequency"`
	Notes            string    `json:"notes"`
}

type PageData struct {
	Medications []Medication
	Reminders   []Medication
	Error       string
}

// Adding medication to the database
func AddMedication(db *gorm.DB, medication Medication) (*Medication, error) {
	if medication.Medication_id == "" || medication.User_id == "" || medication.Medication_name == "" || medication.Dose == "" || medication.Dosage_time.IsZero() || medication.Dosage_frequency == "" || medication.Notes == "" {
		return nil, errors.New("missing required fields")
	}
	// Add medication to the database
	meds := &Medication{
		Medication_id:    medication.Medication_id,
		User_id:          medication.User_id,
		Medication_name:  medication.Medication_name,
		Dose:             medication.Dose,
		Dosage_time:      medication.Dosage_time,
		Dosage_frequency: medication.Dosage_frequency,
		Notes:            medication.Notes,
	}

	db.Create(meds)

	return meds, nil
}

// Getting medications by user_id
func GetMedicationsByUserId(db *gorm.DB, userID string) ([]Medication, error) {
	var medications []Medication

	if err := db.Where("user_id = ?", userID).Find(&medications).Error; err != nil {
		return nil, fmt.Errorf("failed to get medications: %v", err)
	}

	return medications, nil
}

// Get all medications
func GetMedications(db *gorm.DB) ([]Medication, error) {
	var medications []Medication
	if err := db.Find(&medications).Error; err != nil {
		return nil, err // Return error if database query fails
	}
	return medications, nil
}

// update medications by medication_id
func UpdateMedication(db *gorm.DB, medication Medication) (Medication, error) {
	// validate required fields
	if medication.Medication_id == "" || medication.User_id == "" || medication.Medication_name == "" || medication.Dose == "" || medication.Dosage_time.IsZero() || medication.Dosage_frequency == "" || medication.Notes == "" {
		return Medication{}, errors.New("missing required fields")
	}

	// update medication in the database
	if err := db.Model(&Medication{}).Where("medication_id = ? AND user_id = ?", medication.Medication_id, medication.User_id).Updates(Medication{
		Medication_name:  medication.Medication_name,
		Dose:             medication.Dose,
		Dosage_time:      medication.Dosage_time,
		Dosage_frequency: medication.Dosage_frequency,
		Notes:            medication.Notes,
	}).Error; err != nil {
		return Medication{}, fmt.Errorf("failed to update medication: %v", err)
	}

	// get the updated medication from the database to return it
	var updatedMed Medication
	if err := db.Where("medication_id =? AND user_id=?", medication.Medication_id, medication.User_id).First(&updatedMed).Error; err != nil {
		return Medication{}, fmt.Errorf("failed to get updated medication: %v", err)
	}
	return updatedMed, nil
}

// Delete medications by medication_id
func DeleteMedication(db *gorm.DB, medication Medication) error {
	// validate required fields
	if medication.Medication_id == "" || medication.User_id == "" {
		return errors.New("missing required fields")
	}

	// delete medication from the database
	if err := db.Where("medication_id =? AND user_id =?", medication.Medication_id, medication.User_id).Delete(&Medication{}).Error; err != nil {
		return fmt.Errorf("failed to delete medication: %v", err)
	}

	return nil
}

// list medication
func ListMedicationsByUserId(db *gorm.DB, userID string) ([]Medication, error) {
	var medications []Medication
	if err := db.Where("user_id = ?", userID).Find(&medications).Error; err != nil {
		return nil, fmt.Errorf("failed to get medications: %v", err)
	}
	return medications, nil
}

// send medication reminders
func SendMedicationReminders(db *gorm.DB, userID int64) error {
	var medications []Medication
	if err := db.Where("user_id = ?", userID).Find(&medications).Error; err != nil {
		return fmt.Errorf("failed to get medications: %v", err)
	}
	for _, medication := range medications {
		// Check if medication reminder time has passed
		if medication.Dosage_time.Before(time.Now()) {
			// Send reminder email or push notification
			fmt.Printf("Sending reminder for medication %s to user %s\n", medication.Medication_name, medication.User_id)
		}
	}
	return nil
}
