package models

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Director    string `json:"director"`
	ReleaseDate string `json:"releaseDate"`
	Status      string `json:"status"`
}
