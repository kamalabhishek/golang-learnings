package router

import (
	"net/http"

	"movie-api/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/movie", handlers.CreateMovie)
	mux.HandleFunc("/movies", handlers.GetAllMovies)
	mux.HandleFunc("/movie/", handlers.MovieRouter)

	mux.HandleFunc("/movie/review/", handlers.DeleteReview)

	return mux
}
