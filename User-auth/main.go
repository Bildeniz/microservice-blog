package main

import (
	"UserAuth/db"
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Read env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env files")
	}

	// Database Connection
	conn := db.Connect()
	ctx := context.Background()

	// Migrations
	db.Migrate(conn, ctx)

}
