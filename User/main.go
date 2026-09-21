package main

import (
	"context"
	"fmt"
	"log"

	"User/db"

	"github.com/joho/godotenv"
)

func main() {
	// Read env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env files")
	}

	conn := db.Connect()
	ctx := context.Background()

	db.Migrate(conn, ctx)

}
