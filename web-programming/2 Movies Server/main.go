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
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
	//params := mux.Vars(r)
	/*for _, val := range movies {
		if val.Id == params["id"] {
			json.NewEncoder(w).Encode(val)
			return
		}
	}*/
}
func getMoviesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, val := range movies {
		if val.Id == params["id"] {
			err := json.NewEncoder(w).Encode(val)
			if err != nil {
				return
			}
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for i, val := range movies {
		if param["id"] == val.Id {
			movies = append(movies[:i], movies[i+1:]...)
			err := json.NewEncoder(w).Encode(&movies)
			if err != nil {
				return
			}
			break
		}
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for i, val := range movies {
		if val.Id == param["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = param["id"]
			movies = append(movies, movie)
			err := json.NewEncoder(w).Encode(&movie)
			if err != nil {
				return
			}

		}
	}
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Title: "First movie", Director: &Director{Lastname: "Dohn", Firstname: "Jeremy"}})
	movies = append(movies, Movie{Id: "2", Title: "Second movie", Director: &Director{Lastname: "Suza", Firstname: "Liza"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMoviesById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	fmt.Printf("Server is listening")
	log.Fatal(http.ListenAndServe(":8080", r))

}
