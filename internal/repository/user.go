package repository

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"diawise/internal/models"

)

// RegisterUser creates a new user and saves it to the database
func RegisterUser(db *gorm.DB, username, name, email, password string) bool {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// fmt.Println("Error hashing password:", err)
		return false
	}

	user := &models.User{
		Username: username,
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	db.Create(user)

	// Read
	var userFromDB models.User
	db.First(&userFromDB, 1) // Find user with id 1
	// fmt.Println(userFromDB)
	return true
}

// LoginUser checks if a user exists and verifies the password
func LoginUser(db *gorm.DB, username, password string) (*models.User, error) {
	var user models.User
	result := db.Where("username = ? OR email = ?", username, username).First(&user)
	if result.Error != nil {
		// User not found
		// fmt.Printf("User not found: %v\n", result.Error)
		return nil, fmt.Errorf("user not found")
	}

	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Password doesn't match
		// return nil, fmt.Errorf("incorrect password")
	}

	return &user, nil
}
