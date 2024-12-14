package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterUser creates a new user and saves it to the database
func RegisterUser(db *gorm.DB, username, email, password string) bool {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return false
	}

	user := &User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	db.Create(user)

	fmt.Println("USER: ", user)

	// Read
	var userFromDB User
	db.First(&userFromDB, 1) // Find user with id 1
	fmt.Println(userFromDB)
	return true
}

// LoginUser checks if a user exists and verifies the password
func LoginUser(db *gorm.DB, username, password string) (*User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		// User not found
		return nil, fmt.Errorf("user not found")
	}

	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Password doesn't match
		return nil, fmt.Errorf("incorrect password")
	}

	println("Success")
	return &user, nil
}
