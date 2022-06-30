package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onurhan1337/movie-app-backend/controllers"
)

func Route(app *fiber.App) {
	// All routes related to comes here
	app.Get("/movie", controllers.GetAllMovies)
	app.Get("/movie/:movieId", controllers.GetMovie)
	app.Post("/movie", controllers.AddMovie)
	app.Delete("/movie/:movieId", controllers.DeleteMovie)
}
