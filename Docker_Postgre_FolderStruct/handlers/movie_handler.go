package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"Docker_postgre_Folderstruct/database"
	"Docker_postgre_Folderstruct/models"
)


func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	query := `
	SELECT id, title,  isbn, director_firstname, director_lastname
	FROM movies WHERE id=$1
	`

	var movie models.Movie
	var id int

	movie.Director = &models.Director{}

	err := database.DB.QueryRow(query, params["id"]).Scan(
		&id,
		&movie.Title,
		
		&movie.Isbn,
		&movie.Director.Firstname,
		&movie.Director.Lastname,
	)

	if err != nil {
		http.Error(w, "Movie not found", 404)
		return
	}

	movie.ID = strconv.Itoa(id)

	json.NewEncoder(w).Encode(movie)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := database.DB.Query(`
	SELECT id, title,  isbn, director_firstname, director_lastname FROM movies
	`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		var id int

		movie.Director = &models.Director{}

		err := rows.Scan(
			&id,
			&movie.Title,
			
			&movie.Isbn,
			&movie.Director.Firstname,
			&movie.Director.Lastname,
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		movie.ID = strconv.Itoa(id)
		movies = append(movies, movie)
	}

	json.NewEncoder(w).Encode(movies)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if movie.Director == nil {
		http.Error(w, "Director required", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO movies (title,  isbn, director_firstname, director_lastname)
	VALUES ($1, $2, $3, $4) RETURNING id
	`

	var id int

	err = database.DB.QueryRow(
		query,
		movie.Title,
		
		movie.Isbn,
		movie.Director.Firstname,
		movie.Director.Lastname,
	).Scan(&id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	movie.ID = strconv.Itoa(id)

	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	query := `DELETE FROM movies WHERE id=$1`

	result, err := database.DB.Exec(query, params["id"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Movie not found", 404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Movie deleted",
	})
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var movie models.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if movie.Director == nil {
		http.Error(w, "Director required", http.StatusBadRequest)
		return
	}

	query := `
	UPDATE movies
	SET title=$1 , isbn=$2, director_firstname=$3, director_lastname=$4
	WHERE id=$5
	`

	result, err := database.DB.Exec(
		query,
		movie.Title,
		
		movie.Isbn,
		movie.Director.Firstname,
		movie.Director.Lastname,
		params["id"],
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Movie not found", 404)
		return
	}

	movie.ID = params["id"]
	json.NewEncoder(w).Encode(movie)
}