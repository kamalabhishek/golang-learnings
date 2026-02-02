package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"movie-api/models"
	"movie-api/storage"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	movieID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Rating  int    `json:"rating"`
		Comment string `json:"comment"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.Rating < 1 || input.Rating > 5 {
		http.Error(w, "Rating must be between 1 and 5", http.StatusBadRequest)
		return
	}

	storage.ReviewMu.Lock()
	review := models.Review{
		ID:      storage.NextReviewID,
		MovieID: movieID,
		Rating:  input.Rating,
		Comment: input.Comment,
	}
	storage.NextReviewID++
	storage.Reviews = append(storage.Reviews, review)
	storage.ReviewMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

func GetReviewsByMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	movieID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	storage.ReviewMu.Lock()
	defer storage.ReviewMu.Unlock()

	var result []models.Review
	for _, r := range storage.Reviews {
		if r.MovieID == movieID {
			result = append(result, r)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/movie/review/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	storage.ReviewMu.Lock()
	defer storage.ReviewMu.Unlock()

	for i, r := range storage.Reviews {
		if r.ID == id {
			storage.Reviews = append(storage.Reviews[:i], storage.Reviews[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Review not found", http.StatusNotFound)
}
