package db

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Surname  string
	Email    string
	Username string `gorm:"index"`

	CreatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedAt time.Time `gorm:"autoCreatedTime"`
}

func Migrate(conn *gorm.DB, ctx context.Context) {
	err := conn.AutoMigrate(&User{})

	if err != nil {
		log.Fatal("There is a problem with Database Migration!")
	}
}
