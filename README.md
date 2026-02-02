````md
# 📽️ Movie API (Go)

A simple REST API built using Go’s `net/http` package to manage **movies** and their **reviews/ratings**.  
The application uses **in-memory storage**, so all data is lost when the server restarts.  
This project is implemented incrementally as part of a Go API design assignment.

---

## 🚀 Getting Started

### Prerequisites
- Go 1.20 or above

### Setup & Run
```bash
git clone <your-repo-url>
cd movie-api
go mod tidy
go run .
````

Server starts at:

```
http://localhost:8080
```

---

## 🎬 Movie APIs

### Create a Movie

```bash
curl -X POST http://localhost:8080/movie \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Interstellar",
    "director": "Christopher Nolan",
    "releaseDate": "2014-11-07"
  }'
```

---

### Get All Movies

```bash
curl http://localhost:8080/movies
```

---

### Get Movie by ID

```bash
curl http://localhost:8080/movie/1
```

---

### Update Movie by ID

```bash
curl -X PUT http://localhost:8080/movie/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Interstellar (Updated)",
    "director": "Christopher Nolan",
    "releaseDate": "2014-11-07",
    "status": "unavailable"
  }'
```

---

### Delete Movie by ID

```bash
curl -X DELETE http://localhost:8080/movie/1
```

---

## ⭐ Review / Rating APIs

> Reviews are always associated with a movie using `movieId`.

### Add a Review for a Movie

```bash
curl -X POST http://localhost:8080/movie/1/review \
  -H "Content-Type: application/json" \
  -d '{
    "rating": 5,
    "comment": "Amazing movie"
  }'
```

---

### Get All Reviews for a Movie

```bash
curl http://localhost:8080/movie/1/reviews
```

---

### Delete a Review

```bash
curl -X DELETE http://localhost:8080/movie/review/1
```

---

## Error & Validation Notes

* Rating must be between **1 and 5**
* Invalid IDs return `400 Bad Request`
* Missing resources return `404 Not Found`
* Unsupported methods return `405 Method Not Allowed`

---

## Project Structure

```
movie-api/
├── main.go
├── go.mod
├── router/
│   └── router.go
├── handlers/
│   ├── movie_handler.go
│   └── review_handler.go
├── models/
│   ├── movie.go
│   └── review.go
└── storage/
    ├── store.go
    └── review_store.go
```

---

## Notes

* Uses in-memory storage only
* Restarting the app clears all data
* Built using Go standard library (no frameworks)

---

## Future Improvements

* Validate movie existence before adding reviews
* Add average rating per movie
* Add update review API
* Add unit tests
* Add API versioning
* Migrate to Gin/Fiber
