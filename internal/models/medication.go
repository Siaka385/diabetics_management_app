package models

import (
	"time"

	"gorm.io/gorm"
)

// Medication model
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

// PageData for medication
type PageData struct {
	Medications []Medication
	Reminders   []Medication
	Error       string
}
