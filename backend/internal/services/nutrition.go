package services

import "gorm.io/gorm"

type MealType string

type MealLogEntry struct {
	gorm.Model
	MealType   MealType `json:"mealType" validate:"required,oneof=Breakfast Lunch Dinner Snack"`
	FoodItem   string   `json:"foodItem" validate:"required"`
	Weight     float64  `json:"weight" validate:"required,min=0"`
	Proportion float64  `json:"proportion" validate:"required,min=0,max=1"`
}

type DailyMealLog struct {
	gorm.Model
	Entries       []MealLogEntry `json:"entries"`
	TotalCalories float64        `json:"totalCalories,omitempty"`
	TotalCarbs    float64        `json:"totalCarbs,omitempty"`
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
