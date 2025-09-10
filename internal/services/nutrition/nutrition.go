package services

import (
	"diawise/internal/models"

	"gorm.io/gorm"
)

type MealType string

func SaveMealLog(db *gorm.DB, mealLog models.MealLogEntry) error {
	tx := db.Begin()
	if err := tx.Create(&mealLog).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func SaveDietLog(db *gorm.DB, d models.DietProfile) error {
	tx := db.Begin()
	if err := tx.Create(&d).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
