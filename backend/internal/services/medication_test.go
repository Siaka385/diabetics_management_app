package services_test

import (
	"testing"
	"time"

	"diawise/internal/services"
)

func TestAddMedication_EmptyMedicationId(t *testing.T) {
	medicationService := services.NewMedicationService()

	// Create a medication with empty medication_id
	invalidMedication := services.Medication{
		Medication_id:    "",
		User_id:          "user1",
		Medication_name:  "Aspirin",
		Dose:             "500 mg",
		Dosage_time:      time.Now(),
		Dosage_frequency: "daily",
		Notes:            "No side effects",
	}

	_, err := medicationService.AddMedication(invalidMedication)

	// Check if the error matches expected behavior
	if err == nil {
		t.Error("Expected an error when medication_id is empty")
	}
}

func TestAddMedication_MissingUserID(t *testing.T) {
	medicationService := services.NewMedicationService()

	invalidMedication := services.Medication{
		Medication_id:    "123",
		Dose:             "500 mg",
		Dosage_time:      time.Now(),
		Dosage_frequency: "daily",
		Notes:            "No side effects",
	}

	_, err := medicationService.AddMedication(invalidMedication)

	if err == nil {
		t.Error("Expected an error when user_id is empty")
	}
}
