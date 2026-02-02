package storage

import (
	"sync"
	"movie-api/models"
)

var Reviews = []models.Review{}
var NextReviewID = 1
var ReviewMu sync.Mutex
