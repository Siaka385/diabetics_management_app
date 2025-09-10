package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"diawise/internal/models"
)

// Medication repository functions
func AddMedication(db *gorm.DB, medication models.Medication) (*models.Medication, error) {
	if medication.Medication_id == "" || medication.User_id == "" || medication.Medication_name == "" || medication.Dose == "" || medication.Dosage_time.IsZero() || medication.Dosage_frequency == "" || medication.Notes == "" {
		return nil, fmt.Errorf("missing required fields")
	}
	meds := &models.Medication{
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

func GetMedicationsByUserId(db *gorm.DB, userID string) ([]models.Medication, error) {
	var medications []models.Medication
	if err := db.Where("user_id = ?", userID).Find(&medications).Error; err != nil {
		return nil, fmt.Errorf("failed to get medications: %v", err)
	}
	return medications, nil
}

func GetMedications(db *gorm.DB) ([]models.Medication, error) {
	var medications []models.Medication
	if err := db.Find(&medications).Error; err != nil {
		return nil, err
	}
	return medications, nil
}

func UpdateMedication(db *gorm.DB, medication models.Medication) (models.Medication, error) {
	if medication.Medication_id == "" || medication.User_id == "" || medication.Medication_name == "" || medication.Dose == "" || medication.Dosage_time.IsZero() || medication.Dosage_frequency == "" || medication.Notes == "" {
		return models.Medication{}, fmt.Errorf("missing required fields")
	}
	if err := db.Model(&models.Medication{}).Where("medication_id = ? AND user_id = ?", medication.Medication_id, medication.User_id).Updates(models.Medication{
		Medication_name:  medication.Medication_name,
		Dose:             medication.Dose,
		Dosage_time:      medication.Dosage_time,
		Dosage_frequency: medication.Dosage_frequency,
		Notes:            medication.Notes,
	}).Error; err != nil {
		return models.Medication{}, fmt.Errorf("failed to update medication: %v", err)
	}
	var updatedMed models.Medication
	if err := db.Where("medication_id =? AND user_id=?", medication.Medication_id, medication.User_id).First(&updatedMed).Error; err != nil {
		return models.Medication{}, fmt.Errorf("failed to get updated medication: %v", err)
	}
	return updatedMed, nil
}

func DeleteMedication(db *gorm.DB, medication models.Medication) error {
	if medication.Medication_id == "" || medication.User_id == "" {
		return fmt.Errorf("missing required fields")
	}
	if err := db.Where("medication_id =? AND user_id =?", medication.Medication_id, medication.User_id).Delete(&models.Medication{}).Error; err != nil {
		return fmt.Errorf("failed to delete medication: %v", err)
	}
	return nil
}

func ListMedicationsByUserId(db *gorm.DB, userID string) ([]models.Medication, error) {
	return GetMedicationsByUserId(db, userID)
}

func SendMedicationReminders(db *gorm.DB, userID int64) error {
	medications, err := GetMedicationsByUserId(db, fmt.Sprintf("%d", userID))
	if err != nil {
		return err
	}
	for _, medication := range medications {
		if medication.Dosage_time.Before(time.Now()) {
			fmt.Printf("Sending reminder for medication %s to user %s\n", medication.Medication_name, medication.User_id)
		}
	}
	return nil
}

// Add similar for other models as needed
