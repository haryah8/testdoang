package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var movies = []Movie{
	{ID: 1, Title: "Titanic", Year: 2001},
	{ID: 2, Title: "Marvel Deadpool", Year: 2024},
	{ID: 3, Title: "Cek Toko Sebelah", Year: 2020},
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie Movie

	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// "http:localhost:8080/movies/1"
	idStr := strings.TrimPrefix(r.URL.Path, "/movies/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func Routing(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMovie(w, r)
	case http.MethodPost:
		createMovie(w, r)
	case http.MethodDelete:
		deleteMovie(w, r)
	}
}
func main() {
	http.HandleFunc("/movies", Routing)
	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
}
