package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onurhan1337/movie-app-backend/configs"
	"github.com/onurhan1337/movie-app-backend/routes"
)

func main() {
	app := fiber.New()

	// run database
	configs.ConnectDB()

	// routes
	routes.Route(app)

	app.Listen(":8080")
}
