package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "2468",
		Title: "Vacanze Romane",
		Director: &Director{
			Firstname: "Federico",
			Lastname:  "Fellini",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "4680",
		Title: "I soliti ignoti",
		Director: &Director{
			Firstname: "Vittorio",
			Lastname:  "Gassman",
		},
	})
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal()
	}

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, m := range movies {
		if m.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, m := range movies {
		if params["id"] == m.ID {
			json.NewEncoder(w).Encode(m)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	_ = json.NewDecoder(r.Body).Decode(&m)
	m.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, m)
	json.NewEncoder(w).Encode(m)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, m := range movies {
		if params["id"] == m.ID {
			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(m)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
