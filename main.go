package main

import (
    "net/http"
    "errors"
    "github.com/gin-gonic/gin"
)

type movie struct {
    ID          string  `json:"id"`
    Title       string  `json:"title"`
    Year       int  `json:year`
    ImdbRating  float64 `json:imdbRating`
}

var movies = []movie{
    {ID: "1", Title: "The Shawshank Redemption", Year: 1994, ImdbRating: 9.3},
    {ID: "2", Title: "The Godfather", Year: 1972, ImdbRating: 9.2},
    {ID: "3", Title: "The Godfather: Part II", Year: 1974, ImdbRating: 9.0},
}

func main() {
    router := gin.Default()
    router.GET("/movies", getMovies)
    router.GET("/movies/:id", getMovie)
    router.PATCH("/movies/:id", getMovie)
    router.POST("/movies", addMovie)
    router.Run(":8080")
}

// This function get the all movies
func getMovies(context *gin.Context) {
    context.JSON(http.StatusOK, movies)
}

// This function add a new movie
func addMovie(context *gin.Context) {
    var newMovie movie

    if err := context.BindJSON(&newMovie); err != nil {
        return
    }

    movies = append(movies, newMovie)
    context.IndentedJSON(http.StatusCreated, newMovie)
}

func getMovie(context *gin.Context) {
    id := context.Param("id")
    movie, err := getMovieById(id)

    if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    context.IndentedJSON(http.StatusOK, movie)
}

// This function get a movie by ID
func getMovieById(id string) (*movie, error) {
    for i, m := range movies {
        if m.ID == id {
            return &movies[i], nil
        }
    }
    return nil, errors.New("Movie not found")
}