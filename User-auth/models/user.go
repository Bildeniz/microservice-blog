package models

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"` // Denormalizing for checking email is unique wihout asking user service
	Username string `json:"username" binding:"required" gorm:"index;unique"`
	Password string `json:"password" binding:"required"`

	Name      string    `json:"name" binding:"required" gorm:"-"`
	Surname   string    `json:"surname" binding:"required" gorm:"-"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
