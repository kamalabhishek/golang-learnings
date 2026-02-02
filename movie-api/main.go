package main

import (
	"log"
	"net/http"

	"movie-api/router"
)

func main() {
	r := router.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
