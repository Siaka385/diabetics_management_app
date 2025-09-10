package models

import "gorm.io/gorm"

// Room model (assuming from handlers)
type Room struct {
	*gorm.Model
	Name        string
	Description string
	CreatorID   uint
	Members     []User `gorm:"many2many:room_members;"`
}
