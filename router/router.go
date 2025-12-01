package router

import (
	"net/http"
	"movie-api/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/movie", handlers.MovieBase)
	mux.HandleFunc("/movie/", handlers.MovieByID)
	mux.HandleFunc("/movies", handlers.MovieList)

	return mux
}
