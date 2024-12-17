package auth

import "time"

// type User struct {
// 	gorm.Model
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type User struct {
//     ID       uint   `gorm:"primary_key"`
//     Username string `gorm:"unique"`
//     Password string
//     Name     string
//     Email    string
// }

type User struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    Username  string    `gorm:"unique" json:"username"`
    Password  string    `json:"password"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
