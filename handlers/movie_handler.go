package handlers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func MovieBase(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Movie base endpoint"))
}

func MovieByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Movie by ID endpoint"))
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Movie list endpoint"))
}
