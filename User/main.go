package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Read env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env files")
	}

}
