package storage

import (
	"sync"
	"movie-api/models"
)

var (
	Movies = []models.Movie{}
	NextID = 1
	Mu     sync.Mutex
)
