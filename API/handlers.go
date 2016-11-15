package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"database/sql"
    _ "github.com/mattn/go-sqlite3"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "")
}

/* Function that returns all songs in the database */
func getAllSongs(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db")
    checkErr(err)
	
	rows, err := db.Query("SELECT * FROM songs")
    checkErr(err)
	
	fmt.Fprintln(w, "[")
	
	var end bool
	end = false
	
	rows.Next()
	
	for (!end) {
        var sid int
        var artist string
        var song string
        var genre int
		var slength int		
        err = rows.Scan(&sid, &artist, &song, &genre, &slength)
        checkErr(err)
		
		var songGenre string
		songGenreR, err := db.Query("SELECT name FROM genres WHERE id=?", genre)
		songGenreR.Next()
		err = songGenreR.Scan(&songGenre)
		checkErr(err)
		
		row := Song {Id :sid, Artist:artist, Song:song, Genre:songGenre, Length:slength}
		rowE, err := json.Marshal(row)
		
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(rowE))
		
		if(rows.Next()){
			fmt.Fprintln(w, ",")
		} else {
			fmt.Fprintln(w, " ")
			end = true
		}
	}
	
	fmt.Fprintln(w, "]")

	
    db.Close()
}

/* Function for searching song by artist name */
func getSongByArtist(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db")
    checkErr(err)
	
	vars := mux.Vars(r)
    artistRequest := vars["songArtist"]
	
	rows, err := db.Query("SELECT * FROM songs WHERE artist=?", artistRequest)
    checkErr(err)
	
	fmt.Fprintln(w, "[")
	
	var end bool
	end = false
	
	rows.Next()
	
	for (!end) {
        var sid int
        var artist string
        var song string
        var genre int
		var slength int		
        err = rows.Scan(&sid, &artist, &song, &genre, &slength)
        checkErr(err)
		
		var songGenre string
		songGenreR, err := db.Query("SELECT name FROM genres WHERE id=?", genre)
		songGenreR.Next()
		err = songGenreR.Scan(&songGenre)
		checkErr(err)
		
		row := Song {Id :sid, Artist:artist, Song:song, Genre:songGenre, Length:slength}
		rowE, err := json.Marshal(row)
		
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(rowE))
		
		if(rows.Next()){
			fmt.Fprintln(w, ",")
		} else {
			fmt.Fprintln(w, " ")
			end = true
		}
	}
	
	fmt.Fprintln(w, "]")
	
    db.Close()
}

/* Function for searching song by song name */
func getSongByName(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db")
    checkErr(err)
	
	vars := mux.Vars(r)
    songRequest := vars["songName"]
	
	rows, err := db.Query("SELECT * FROM songs WHERE song=?", songRequest)
    checkErr(err)
	
	fmt.Fprintln(w, "[")
	
	var end bool
	end = false
	
	rows.Next()
	
	for (!end) {
        var sid int
        var artist string
        var song string
        var genre int
		var slength int		
        err = rows.Scan(&sid, &artist, &song, &genre, &slength)
        checkErr(err)
		
		var songGenre string
		songGenreR, err := db.Query("SELECT name FROM genres WHERE id=?", genre)
		songGenreR.Next()
		err = songGenreR.Scan(&songGenre)
		checkErr(err)
		
		row := Song {Id :sid, Artist:artist, Song:song, Genre:songGenre, Length:slength}
		rowE, err := json.Marshal(row)
		
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(rowE))
		
		if(rows.Next()){
			fmt.Fprintln(w, ",")
		} else {
			fmt.Fprintln(w, " ")
			end = true
		}
	}
	
	fmt.Fprintln(w, "]")
	
    db.Close()
}

/* Function for searching song by genre name */
func getSongByGenre(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db")
    checkErr(err)
	
	vars := mux.Vars(r)
    genreRequest := vars["songGenre"]
	
	var genreId int
	genreIdR, err := db.Query("SELECT id FROM genres WHERE name=?", genreRequest)
	genreIdR.Next()
	err = genreIdR.Scan(&genreId)
	checkErr(err)
	
	rows, err := db.Query("SELECT * FROM songs WHERE genre=?", genreId)
    checkErr(err)
	
	fmt.Fprintln(w, "[")
	
	var end bool
	end = false
	
	rows.Next()
	
	for (!end) {
        var sid int
        var artist string
        var song string
        var genre int
		var slength int		
        err = rows.Scan(&sid, &artist, &song, &genre, &slength)
        checkErr(err)
		
		row := Song {Id :sid, Artist:artist, Song:song, Genre:genreRequest, Length:slength}
		rowE, err := json.Marshal(row)
		
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(rowE))
		
		if(rows.Next()){
			fmt.Fprintln(w, ",")
		} else {
			fmt.Fprintln(w, " ")
			end = true
		}
	}
	
	fmt.Fprintln(w, "]")
	
    db.Close()
}

/* Function that returns all genres in the database */
func getAllGenres(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db")
    checkErr(err)
	
	rows, err := db.Query("SELECT * FROM genres")
    checkErr(err)
	
	fmt.Fprintln(w, "[")
	
	var end bool
	end = false
	
	rows.Next()
	
	for (!end) {
        var gid int
        var gname string	
        err = rows.Scan(&gid, &gname)
        checkErr(err)
		
		row := Genre {Id :gid, Name:gname}
		rowE, err := json.Marshal(row)
		
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(rowE))
		
		if(rows.Next()){
			fmt.Fprintln(w, ",")
		} else {
			fmt.Fprintln(w, " ")
			end = true
		}
	}

	fmt.Fprintln(w, "]")

    db.Close()
	
}

/* Function for searching genre by genre name */
func getGenreByName(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db")
    checkErr(err)
	
	vars := mux.Vars(r)
    genreRequest := vars["genreName"]
	
	rows, err := db.Query("SELECT * FROM genres WHERE name=?", genreRequest)
    checkErr(err)
	
	fmt.Fprintln(w, "[")
	
	var end bool
	end = false
	
	rows.Next()
	
	for (!end) {
        var gid int
        var gname string	
        err = rows.Scan(&gid, &gname)
        checkErr(err)
		
		row := Genre {Id :gid, Name:gname}
		rowE, err := json.Marshal(row)
		
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(rowE))
		
		if(rows.Next()){
			fmt.Fprintln(w, ",")
		} else {
			fmt.Fprintln(w, " ")
			end = true
		}
	}
	
	fmt.Fprintln(w, "]")
	
    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}