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

/* Route handlers */

/* Function that shows the index/services page */
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html") /* Sets the content type to HTML */
	tmpl, err := template.New("Services").Parse(doc) /* Parses the template */
	if err == nil{
		context := Context{"Welcome!", "By Marvin Carranza Romero"}
		tmpl.Execute(w, context) /* Shows the index/services page */
	}
}

/* Function that returns all songs in the database */
func getAllSongs(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") /* Sets content type to JSON */
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db") /* Opens the database */
    checkErr(err)
	
	rows, err := db.Query("SELECT * FROM songs") /* Query: Selects all songs in database */
    checkErr(err)
	
	fmt.Fprintln(w, "[") /* JSON start */
	
	var end bool /* Var used as the loop end condition  */
	end = false
	
	rows.Next()
	
	/* Result rows loop */
	for (!end) {
        var sid int /* Var for song id */
        var artist string /* Var for song artist */
        var song string /* Var for song name */
        var genre int /* Var for genre id */
		var slength int /* Var for song length */
        err = rows.Scan(&sid, &artist, &song, &genre, &slength) /* Sets the values from the row */
        checkErr(err)
		
		var songGenre string /* Var for genre name */
		songGenreR, err := db.Query("SELECT name FROM genres WHERE id=?", genre) /* Query: Selects the genre name that matches with the genre id */
		songGenreR.Next()
		err = songGenreR.Scan(&songGenre) /* Sets the genre name */
		checkErr(err)
		
		row := Song {Id :sid, Artist:artist, Song:song, Genre:songGenre, Length:slength}
		rowE, err := json.Marshal(row) /* JSON encode of the row */
		
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
	
	fmt.Fprintln(w, "]") /* JSON end */

    db.Close() /* Closes the database */
}

/* Function for searching song by artist name */
func getSongByArtist(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") /* Sets content type to JSON */
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db") /* Opens the database */
    checkErr(err)
	
	vars := mux.Vars(r)
    artistRequest := vars["songArtist"] /* Sets the artist var with the requested one */
	
	rows, err := db.Query("SELECT * FROM songs WHERE artist=?", artistRequest) /* Query: Selects all songs of the selected artist */
    checkErr(err)
	
	var end bool /* Var used as the loop end condition  */
	end = false
		
	if(rows.Next()){

		fmt.Fprintln(w, "[") /* JSON start */
		
		/* Result rows loop */
		for (!end) {
			var sid int /* Var for song id */
			var artist string /* Var for song artist */
			var song string /* Var for song name */
			var genre int /* Var for genre id */
			var slength int	/* Var for song length */
			err = rows.Scan(&sid, &artist, &song, &genre, &slength) /* Sets the values from the row */
			checkErr(err)
			
			var songGenre string /* Var for genre name */
			songGenreR, err := db.Query("SELECT name FROM genres WHERE id=?", genre) /* Query: Selects the genre name that matches with the genre id */
			songGenreR.Next()
			err = songGenreR.Scan(&songGenre) /* Sets the genre name */
			checkErr(err)
			
			row := Song {Id :sid, Artist:artist, Song:song, Genre:songGenre, Length:slength}
			rowE, err := json.Marshal(row) /* JSON encode of the row */
			
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
		
		fmt.Fprintln(w, "]") /* JSON end */
	} else {
		fmt.Fprintln(w, "[") /* JSON start */
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]") /* JSON end */
	}
	
    db.Close() /* Closes the database */
}

/* Function for searching song by song name */
func getSongByName(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") /* Sets content type to JSON */
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db") /* Opens the database */
    checkErr(err)
	
	vars := mux.Vars(r)
    songRequest := vars["songName"] /* Sets the song name var with the requested one */
	
	rows, err := db.Query("SELECT * FROM songs WHERE song=?", songRequest) /* Query: Selects all songs of the selected name */
    checkErr(err)

	var end bool /* Var used as the loop end condition  */
	end = false
	
	if (rows.Next()){
		
		fmt.Fprintln(w, "[") /* JSON start */
		
		/* Result rows loop */
		for (!end) {
			var sid int /* Var for song id */
			var artist string /* Var for song artist */
			var song string /* Var for song name */
			var genre int /* Var for genre id */
			var slength int	/* Var for song length */	
			err = rows.Scan(&sid, &artist, &song, &genre, &slength) /* Sets the values from the row */
			checkErr(err)
			
			var songGenre string /* Var for genre name */
			songGenreR, err := db.Query("SELECT name FROM genres WHERE id=?", genre) /* Query: Selects the genre name that matches with the genre id */
			songGenreR.Next()
			err = songGenreR.Scan(&songGenre) /* Sets the genre name */
			checkErr(err)
			
			row := Song {Id :sid, Artist:artist, Song:song, Genre:songGenre, Length:slength}
			rowE, err := json.Marshal(row) /* JSON encode of the row */
			
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
		
		fmt.Fprintln(w, "]") /* JSON end */
	
	} else {
		fmt.Fprintln(w, "[") /* JSON start */
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]") /* JSON end */
	}
	
    db.Close() /* Closes the database */
}

/* Function for searching song by genre name */
func getSongByGenre(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") /* Sets content type to JSON */
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db") /* Opens the database */
    checkErr(err)
	
	vars := mux.Vars(r)
    genreRequest := vars["songGenre"] /* Sets the song genre var with the requested one */
	
	var genreId int
	genreIdR, err := db.Query("SELECT id FROM genres WHERE name=?", genreRequest) /* Selects the id of the selected genre */
	
	if (genreIdR.Next()){
		
		err = genreIdR.Scan(&genreId)
		checkErr(err)
		
		rows, err := db.Query("SELECT * FROM songs WHERE genre=?", genreId) /* Query: Selects the songs of the selected genre */
		checkErr(err)
		rows.Next()
		
		var end bool /* Var used as the loop end condition  */
		end = false
		
		fmt.Fprintln(w, "[") /* JSON start */
		
		/* Result rows loop */
		for (!end) {
			var sid int /* Var for song id */
			var artist string /* Var for song artist */
			var song string /* Var for song name */
			var genre int /* Var for genre id */
			var slength int	/* Var for song length */	
			err = rows.Scan(&sid, &artist, &song, &genre, &slength) /* Sets the values from the row */
			checkErr(err)
			
			row := Song {Id :sid, Artist:artist, Song:song, Genre:genreRequest, Length:slength}
			rowE, err := json.Marshal(row) /* JSON encode of the row */
			
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
		
		fmt.Fprintln(w, "]") /* JSON end */
	
	} else {
		fmt.Fprintln(w, "[") /* JSON start */
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]") /* JSON end */
	}
	
    db.Close() /* Closes the database */
}

/* Function for searching songs by length range */
func getSongsByLength(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") /* Sets content type to JSON */
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db") /* Opens the database */
    checkErr(err)
	
	vars := mux.Vars(r)
    minLength := vars["minLength"] /* Sets the minimum length var with the requested one */
	maxLength := vars["maxLength"] /* Sets the maximum length var with the requested one */
	
	rows, err := db.Query("SELECT * FROM songs WHERE length BETWEEN ? AND ?", minLength, maxLength) /* Query: Selects all songs of the selected length range */
    checkErr(err)	
	
	var end bool /* Var used as the loop end condition  */
	end = false
	
	if (rows.Next()){
		
		fmt.Fprintln(w, "[") /* JSON start */
	
		/* Result rows loop */
		for (!end) {
			var sid int /* Var for song id */
			var artist string /* Var for song artist */
			var song string /* Var for song name */
			var genre int /* Var for genre id */
			var slength int	/* Var for song length */	
			err = rows.Scan(&sid, &artist, &song, &genre, &slength) /* Sets the values from the row */
			checkErr(err)
			
			var songGenre string /* Var for genre name */
			songGenreR, err := db.Query("SELECT name FROM genres WHERE id=?", genre) /* Query: Selects the genre name that matches with the genre id */
			songGenreR.Next()
			err = songGenreR.Scan(&songGenre) /* Sets the genre name */
			checkErr(err)
			
			row := Song {Id :sid, Artist:artist, Song:song, Genre:songGenre, Length:slength}
			rowE, err := json.Marshal(row) /* JSON encode of the row */
			
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
		
		fmt.Fprintln(w, "]") /* JSON end */
	
	} else {
		fmt.Fprintln(w, "[") /* JSON start */
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]") /* JSON end */
	}
	 
    db.Close() /* Closes the database */
}


/* Function that returns all genres in the database */
func getAllGenres(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") /* Sets content type to JSON */
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db") /* Opens the database */
    checkErr(err)
	
	rows, err := db.Query("SELECT * FROM genres") /* Query: Selects all genres in the database */
    checkErr(err)
	
	fmt.Fprintln(w, "[") /* JSON start */
	
	var end bool /* Var used as the loop end condition  */
	end = false
	
	rows.Next()
	
	/* Result rows loop */
	for (!end) {
        var gid int /* Var for genre id */
        var gname string /* Var for genre name */
        err = rows.Scan(&gid, &gname) /* Sets the values from the row */
        checkErr(err)
		
		row := Genre {Id :gid, Name:gname}
		rowE, err := json.Marshal(row) /* JSON encode of the row */
		
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

	fmt.Fprintln(w, "]") /* JSON end */

    db.Close() /* Closes the database */
	
}

/* Function for searching genre by genre name */
func getGenreByName(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") /* Sets content type to JSON */
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db") /* Opens the database */
    checkErr(err)
	
	vars := mux.Vars(r)
    genreRequest := vars["genreName"] /* Sets the genre name var with the requested one */
	
	rows, err := db.Query("SELECT * FROM genres WHERE name=?", genreRequest) /* Query: Selects all genres of the selected name */
    checkErr(err)
	
	var end bool /* Var used as the loop end condition  */
	end = false
	
	if (rows.Next()){
		
		fmt.Fprintln(w, "[") /* JSON start */
	
		/* Result rows loop */
		for (!end) {
			var gid int /* Var for genre id */
			var gname string /* Var for genre name */
			err = rows.Scan(&gid, &gname) /* Sets the values from the row */
			checkErr(err)
			
			row := Genre {Id :gid, Name:gname}
			rowE, err := json.Marshal(row) /* JSON encode of the row */
			
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
		
		fmt.Fprintln(w, "]") /* JSON end */
	
	} else {
		fmt.Fprintln(w, "[") /* JSON start */
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "]") /* JSON end */
	}
	 
    db.Close() /* Closes the database */
}

/* Function that returns all genres, number of songs, and total length of all the songs by genre */
func getAllGenresExtra(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8") /* Sets content type to JSON */
    w.WriteHeader(http.StatusOK)
	
	db, err := sql.Open("sqlite3", "../src/github.com/macarranza/API/database/jrdd.db") /* Opens the database */
    checkErr(err)
	
	rows, err := db.Query("SELECT * FROM genres") /* Query: Selects all genres in the database */
    checkErr(err)
	
	fmt.Fprintln(w, "[")
	
	var end bool /* Var used as the loop end condition  */
	end = false
	
	rows.Next()
	
	/* Result rows loop */
	for (!end) {
        var gid int /* Var for genre id */
        var gname string /* Var for genre name */
        err = rows.Scan(&gid, &gname) /* Sets the values from the row */
        checkErr(err)
		
		var numberOfSongs int /* Var for number of songs */
		var totalLength int /* Var for total length of songs */
		numberOfSongsR, err := db.Query("SELECT COUNT(DISTINCT ID) FROM songs WHERE genre=?", gid) /* Query: Counts the number of songs of the selected genre */
		totalLengthR, err := db.Query("SELECT TOTAL(length) FROM songs WHERE genre=?", gid) /* Query: Sums the length of all the songs of the selected genre */
		numberOfSongsR.Next()
		totalLengthR.Next()
		err = numberOfSongsR.Scan(&numberOfSongs) /* Sets the number of songs */
		checkErr(err)
		err = totalLengthR.Scan(&totalLength) /* Sets the total length of the songs */
		checkErr(err)
		
		row := GenreExtra {Id :gid, Name:gname, NumberOfSongs:numberOfSongs, TotalLength:totalLength}
		rowE, err := json.Marshal(row) /* JSON encode of the row */
		
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

	fmt.Fprintln(w, "]") /* JSON end */

    db.Close() /* Closes the database */
	
}

/* Checks if there was an error */
func checkErr(err error) { 
    if err != nil {
        panic(err)
    }
}