package main

import (
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

	db.Connect()
	fmt.Println("Connection is succes!")
}
