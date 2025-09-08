package models

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Username string `json:"username" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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

// MealLogEntry model
type MealLogEntry struct {
	*gorm.Model
	UserID         uint
	DailyMealLogID uint
	DietProfileID  uint
	FoodName       string
	Quantity       float64
	Calories       float64
	Carbs          float64
	Protein        float64
	Fat            float64
	Date           time.Time
}

// DailyMealLog model
type DailyMealLog struct {
	*gorm.Model
	UserID uint
	Date   time.Time
	Entries []MealLogEntry
}

// DietProfile model
type DietProfile struct {
	*gorm.Model
	UserID       uint
	FoodName     string
	Quantity     float64
	Calories     float64
	Carbs        float64
	Protein      float64
	Fat          float64
	HealthScore  int
	Suggestions  string
	Date         time.Time
}

func (profile *DietProfile) ParseDietProfileString(data string) error {
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
		case "CaloriesIntake":
			profile.Calories = value
		case "CarbIntake":
			profile.Carbs = value
		case "ProteinIntake":
			profile.Protein = value
		case "FatIntake":
			profile.Fat = value
		case "SugarConsumption":
			// ignore or set to something
		case "ProcessedFoodRatio":
			// ignore
		}
	}
	return nil
}

// BlogPost model
type BlogPost struct {
	Title   string
	Author  string
	Date    string
	Content string
}

// Post model
type Post struct {
	ID      string
	Title   string
	Excerpt template.HTML
	Author  string
	Date    string
	Content template.HTML
	Image   string
	Abbrev  string
}

// Issue model
type Issue struct {
	StatusCode int
	Problem    string
}

// UserProfile for templates
type UserProfile struct {
	Name   string
	Abbrev string
}

// Room model (assuming from handlers)
type Room struct {
	*gorm.Model
	Name        string
	Description string
	CreatorID   uint
	Members     []User `gorm:"many2many:room_members;"`
}




// Meal Plan
type MealPlanRequest struct {
	Duration           string   `json:"mp_duration"`
	MealTypes          string   `json:"mp_type"`
	DietaryPreferences string   `json:"mp_diet_pref"`
	PreferredFoods     []string `json:"mp_preferred_foods"`
	FoodRestrictions   []string `json:"mp_avoid_foods"`
}

type Meal struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"` // Breakfast, Lunch, etc.
	Ingredients []string `json:"ingredients"`
}

type MealPlanResponse struct {
	Duration string `json:"duration"`
	Meals    []Meal `json:"meals"`
	Message  string `json:"message"`
}

// Meal log
type NutritionResponse struct {
	Message      string `json:"message"`
	MealInsights string `json:"mealInsights"`
}

type NutrientInfo struct {
	gorm.Model
	UserID   string  `json:"user_id"`
	Calories float64 `json:"calories"`
	Carbs    float64 `json:"carbs"`
	Protein  float64 `json:"protein"`
	Fat      float64 `json:"fat"`
	Fiber    float64 `json:"fiber"`
	Vitamins string  `json:"vitamins"`
	Minerals string  `json:"minerals"`
}

type FoodLog struct {
	gorm.Model
	UserID    string     `json:"user_id"`
	MealItems []MealItem `gorm:"foreignKey:FoodLogID" json:"meal_items"`
}

type MealItem struct {
	gorm.Model
	FoodItem   string  `json:"food_item"`
	Weight     float64 `json:"weight"`
	Proportion float64 `json:"proportion"`
	FoodLogID  uint    `json:"-"`
}

type FoodItem struct {
	Name        string             `json:"name"`
	ServingSize float64            `json:"serving_size"`
	Calories    float64            `json:"calories"`
	Carbs       float64            `json:"carbs"`
	Protein     float64            `json:"protein"`
	Fat         float64            `json:"fat"`
	Fiber       float64            `json:"fiber"`
	Vitamins    map[string]float64 `json:"vitamins"`
	Minerals    map[string]float64 `json:"minerals"`
}
