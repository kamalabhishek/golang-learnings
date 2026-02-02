package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"movie-api/models"
	"movie-api/storage"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Title       string `json:"title"`
		Director    string `json:"director"`
		ReleaseDate string `json:"releaseDate"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.Title == "" || input.Director == "" || input.ReleaseDate == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	storage.Mu.Lock()
	movie := models.Movie{
		ID:          storage.NextID,
		Title:       input.Title,
		Director:    input.Director,
		ReleaseDate: input.ReleaseDate,
		Status:      "available",
	}
	storage.NextID++
	storage.Movies = append(storage.Movies, movie)
	storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Movies)
}

func MovieByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/movie/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getMovie(w, id)
	case http.MethodPut:
		updateMovie(w, r, id)
	case http.MethodDelete:
		deleteMovie(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getMovie(w http.ResponseWriter, id int) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for _, m := range storage.Movies {
		if m.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(m)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func updateMovie(w http.ResponseWriter, r *http.Request, id int) {
	var input models.Movie
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for i, m := range storage.Movies {
		if m.ID == id {
			if input.Title != "" {
				m.Title = input.Title
			}
			if input.Director != "" {
				m.Director = input.Director
			}
			if input.ReleaseDate != "" {
				m.ReleaseDate = input.ReleaseDate
			}
			if input.Status != "" {
				m.Status = input.Status
			}

			storage.Movies[i] = m
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(m)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, id int) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for i, m := range storage.Movies {
		if m.ID == id {
			storage.Movies = append(storage.Movies[:i], storage.Movies[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func MovieRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/movie/")
	parts := strings.Split(path, "/")

	if len(parts) == 1 {
		MovieByID(w, r)
		return
	}

	if len(parts) == 2 && parts[1] == "review" {
		CreateReview(w, r)
		return
	}

	if len(parts) == 2 && parts[1] == "reviews" {
		GetReviewsByMovie(w, r)
		return
	}

	http.NotFound(w, r)
}
