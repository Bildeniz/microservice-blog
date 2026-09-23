package models

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Username string `json:"username" binding:"required" gorm:"index;unique"`

	Password string `json:"password" binding:"required" gorm:"-"`

	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
