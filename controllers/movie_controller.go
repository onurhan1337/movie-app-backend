package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/onurhan1337/movie-app-backend/configs"
	"github.com/onurhan1337/movie-app-backend/models"
	"github.com/onurhan1337/movie-app-backend/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var movieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movies")
var validate = validator.New()

// Add a new movie
func AddMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var movie models.Movie
	defer cancel()

	// validate the request body
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&movie); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newMovie := models.Movie{
		Id:         primitive.NewObjectID(),
		Title:      movie.Title,
		Image:      movie.Image,
		ImdbRating: movie.ImdbRating,
		Year:       movie.Year,
		Director:   movie.Director,
	}

	result, err := movieCollection.InsertOne(ctx, newMovie)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Data: &fiber.Map{"data": result}})
}

// Get all movies
func GetAllMovies(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var movies []models.Movie
	defer cancel()

	results, err := movieCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Data: &fiber.Map{"data": err.Error()}})
	}

	// reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleMovie models.Movie
		if err = results.Decode(&singleMovie); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Data: &fiber.Map{"data": err.Error()}})
		}

		movies = append(movies, singleMovie)
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Data: &fiber.Map{"data": movies}})
}

// Get a movie by id
func GetMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movieId := c.Params("movieId")
	var movie models.Movie
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(movieId)

	err := movieCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&movie)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Data: &fiber.Map{"data": movie}})
}

// Delete a movie by id
func DeleteMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movieId := c.Params("movieId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(movieId)

	result, err := movieCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.Response{Status: http.StatusNotFound, Data: &fiber.Map{"data": "Movie with specified ID not found!"}})
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Data: &fiber.Map{"data": "Movie deleted successfully!"}})
}
