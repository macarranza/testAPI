package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"database/sql"
	"text/template"
    _ "github.com/mattn/go-sqlite3"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("Welcome Template").Parse(doc)
	if err == nil{
		context := Context{"Welcome!", "By Marvin Carranza Romero"}
		tmpl.Execute(w, context)
	}
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
	
	var end bool
	end = false
		
	if(rows.Next()){

		fmt.Fprintln(w, "[")
		
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
	} else {
		fmt.Fprintln(w, "[")
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]")
	}
	
	
	
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

	var end bool
	end = false
	
	if (rows.Next()){
		
		fmt.Fprintln(w, "[")
		
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
	
	} else {
		fmt.Fprintln(w, "[")
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]")
	}
	
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
	
	if (genreIdR.Next()){
		
		err = genreIdR.Scan(&genreId)
		checkErr(err)
		
		rows, err := db.Query("SELECT * FROM songs WHERE genre=?", genreId)
		checkErr(err)
		rows.Next()
		
		var end bool
		end = false
		
		fmt.Fprintln(w, "[")
		
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
	
	} else {
		fmt.Fprintln(w, "[")
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]")
	}
	
    db.Close()
}

/* Function for searching songs by length */
func getSongsByLength(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db")
    checkErr(err)
	
	vars := mux.Vars(r)
    minLength := vars["minLength"]
	maxLength := vars["maxLength"]
	
	rows, err := db.Query("SELECT * FROM songs WHERE length BETWEEN ? AND ?", minLength, maxLength)
    checkErr(err)	
	
	var end bool
	end = false
	
	if (rows.Next()){
		
		fmt.Fprintln(w, "[")
	
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
	
	} else {
		fmt.Fprintln(w, "[")
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]")
	}
	
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
	
	var end bool
	end = false
	
	if (rows.Next()){
		
		fmt.Fprintln(w, "[")
	
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
	
	} else {
		fmt.Fprintln(w, "[")
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]")
	}
	
    db.Close()
}

/* Function that returns all genres, number of songs, and total length of all the songs by genre */
func getAllGenresExtra(w http.ResponseWriter, r *http.Request){
	
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
		
		var numberOfSongs int
		var totalLength int
		numberOfSongsR, err := db.Query("SELECT COUNT(DISTINCT ID) FROM songs WHERE genre=?", gid)
		totalLengthR, err := db.Query("SELECT TOTAL(length) FROM songs WHERE genre=?", gid)
		numberOfSongsR.Next()
		totalLengthR.Next()
		err = numberOfSongsR.Scan(&numberOfSongs)
		checkErr(err)
		err = totalLengthR.Scan(&totalLength)
		checkErr(err)
		
		row := GenreExtra {Id :gid, Name:gname, NumberOfSongs:numberOfSongs, TotalLength:totalLength}
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