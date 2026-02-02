package models

type Review struct {
	ID      int    `json:"id"`
	MovieID int    `json:"movieId"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
