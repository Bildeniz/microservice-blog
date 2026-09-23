package main

import (
	"context"
	"log"

	controller "User/controllers"
	"User/db"
	"User/repositories"
	"User/routes"

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

	// Repo initialize
	repo := repositories.NewUserRepository(conn)

	// Controller initialize
	ctrl := controller.NewUserController(repo)

	// Routes initalize
	router := routes.SetupRouter(ctrl)
	if err := router.Run(":80"); err != nil {
		log.Fatal(err)
	}
}
