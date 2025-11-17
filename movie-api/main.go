package main

import (
	"fmt"
	"log"
	"net/http"
	"movie-api/handlers"
)

func main() {
	http.HandleFunc("/movie", handlers.CreateMovie)
	http.HandleFunc("/movies", handlers.GetAllMovies)
	http.HandleFunc("/movie/", handlers.MovieByID)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
