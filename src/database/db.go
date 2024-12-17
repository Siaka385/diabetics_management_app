package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	auth "diawise/src/auth"
	api "diawise/src/api"
	medication "diawise/src/services"
)

func InitializeDatabase(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&auth.User{})
	// db.AutoMigrate(&api.FoodLog{})
	// db.AutoMigrate(&api.NutrientInfo{})
	// db.AutoMigrate(&api.MealItem{})
	db.AutoMigrate(&medication.Medication{})
	db.AutoMigrate(&medication.MealLogEntry{})
	db.AutoMigrate(&medication.DailyMealLog{})
	db.AutoMigrate(&api.Room{})

	return db
}
