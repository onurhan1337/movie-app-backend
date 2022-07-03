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

	// Run on Local Server.
	// In any need-cases you can comment line 30 (deployment running),
	// and uncomment 27 (local running).

	// Warning: Code shouldn't be pushed if app is running on local)

	// log.Fatal(http.ListenAndServe(":8080", router))

	app.Listen(":" + port)
}
