package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	movies = append(movies, Movie{Id: "1", Isbn: "12345", Title: "Black Adam",
		Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "54321", Title: "Spiderman",
		Director: &Director{FirstName: "Steven", LastName: "Smith"}})

	router := mux.NewRouter()
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies/{id}", updateMovieById).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovieById).Methods("DELETE")

	fmt.Printf("Starting Go HTTP server on port: 8000.\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func createMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(response).Encode(movie)
}

func getMovies(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(movies)
}

func getMovieById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, movie := range movies {
		if movie.Id == params["id"] {
			json.NewEncoder(response).Encode(movie)
			return
		}
	}
}

func updateMovieById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(response).Encode(movie)
			return
		}
	}
}

func deleteMovieById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(movies)
}
