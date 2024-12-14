package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	auth "diawise/internal/auth"
	medication "diawise/internal/services"
)

func InitializeDatabase(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&auth.User{})
	db.AutoMigrate(&medication.Medication{})

	return db
}
