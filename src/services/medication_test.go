package services_test

import (
	"testing"
	"time"

	"diawise/src/services"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestAddMedication_ErrorWhenUserIDIsEmpty(t *testing.T) {
	db := new(gorm.DB)
	medication := services.Medication{
		Medication_id:    "1234",
		User_id:          "",
		Medication_name:  "Aspirin",
		Dose:             "500 mg",
		Dosage_time:      time.Now(),
		Dosage_frequency: "daily",
		Notes:            "No side effects",
	}
	_, err := services.AddMedication(db, medication)
	require.Error(t, err)
	require.Equal(t, "missing required fields", err.Error())
}

func TestAddMedication_ErrorWhenDoseIsEmpty(t *testing.T) {
	db := new(gorm.DB)
	medication := services.Medication{
		Medication_id:    "1234",
		User_id:          "user1",
		Medication_name:  "Aspirin",
		Dose:             "",
		Dosage_time:      time.Now(),
		Dosage_frequency: "daily",
		Notes:            "No side effects",
	}
	_, err := services.AddMedication(db, medication)
	require.Error(t, err)
	require.Equal(t, "missing required fields", err.Error())
}
