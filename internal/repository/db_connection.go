package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"diawise/internal/models"
)

func InitializeDatabase(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Medication{})
	db.AutoMigrate(&models.MealLogEntry{})
	db.AutoMigrate(&models.DailyMealLog{})
	db.AutoMigrate(&models.DietProfile{})
	db.AutoMigrate(&models.Room{})

	return db
}
