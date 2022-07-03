package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/onurhan1337/movie-app-backend/configs"
	"github.com/onurhan1337/movie-app-backend/routes"
)

func main() {
	app := fiber.New()

	godotenv.Load()
	port := os.Getenv("PORT")

	// run database
	configs.ConnectDB()

	// routes
	routes.Route(app)

	app.Listen(":" + port)
}
