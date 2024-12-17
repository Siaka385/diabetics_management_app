package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type MealType string

type MealLogEntry struct {
	gorm.Model
	UserID         uint      `json:"user_id" validate:"required"`
	DailyMealLogID uint        `json:"daily_meal_log_id"`
	MealType       MealType    `json:"mealType" validate:"required,oneof=Breakfast Lunch Dinner Snack"`
	FoodItem       string      `json:"foodItem" validate:"required"`
	Weight         float64     `json:"weight" validate:"required,min=0"`
	Proportion     float64     `json:"proportion" validate:"required,min=0,max=1"`
	DietProfileID  uint        `json:"diet_profile_id"`
	DietProfile    DietProfile `gorm:"foreignKey:DietProfileID"`
}

type DailyMealLog struct {
	gorm.Model
	UserID        uint         `json:"user_id"`
	Entries       []MealLogEntry `json:"entries"`
	Date          time.Time      `json:"date"`
	TotalCalories float64        `json:"totalCalories,omitempty"`
	TotalCarbs    float64        `json:"totalCarbs,omitempty"`
}

type DietProfile struct {
	gorm.Model
	UserID             uint `json:"user_id"`
	FoodName           string
	CarbIntake         float64 // percentage of daily calories from carbs
	ProteinIntake      float64 // percentage of daily calories from protein
	FatIntake          float64 // percentage of daily calories from fat
	SugarConsumption   float64 // grams of added sugar per day
	WaterIntake        float64 // liters per day
	ProcessedFoodRatio float64 // percentage of diet from processed foods
}

func (profile *DietProfile) ParseDietProfileString(data string) error {
	// Check if the input string appears to be in JSON format
	if strings.HasPrefix(strings.TrimSpace(data), "{") {
		err := json.Unmarshal([]byte(data), profile)
		if err == nil {
			return nil
		}
	}
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "\"", "")
		line = strings.ReplaceAll(line, ":", "")
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return fmt.Errorf("error parsing value for %s: %w", key, err)
		}

		switch key {
		case "CarbIntake":
			profile.CarbIntake = value
		case "ProteinIntake":
			profile.ProteinIntake = value
		case "FatIntake":
			profile.FatIntake = value
		case "SugarConsumption":
			profile.SugarConsumption = value
		case "ProcessedFoodRatio":
			profile.ProcessedFoodRatio = value
		}

	}
	return nil
}

func SaveMealLog(db *gorm.DB, mealLog MealLogEntry) error {
	tx := db.Begin()
	if err := tx.Create(&mealLog).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func SaveDietLog(db *gorm.DB, d DietProfile) error {
	tx := db.Begin()
	if err := tx.Create(&d).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
