package db

import (
	"UserConsumer/models"
	"context"
	"log"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB, ctx context.Context) error {
	err := conn.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("There is a problem with Database Migration!")
		return err
	}

	return nil
}
