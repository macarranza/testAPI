package main

import "net/http"

/* Structs for the JSON encoding */

/* Song struct */
type Song struct {
	Id int `json:"ID"`
	Artist string `json:"artist"`
	Song string `json:"song"`
	Genre string `json:"genre"`
	Length int `json:"length"`
}

/* Genre struct */
type Genre struct {
	Id int `json:"ID"`
	Name string `json:"name"`
}

/* Genre struct for the song count and total length by genre */
type GenreExtra struct {
	Id int `json:"ID"`
	Name string `json:"name"`
	NumberOfSongs int `json:"numberOfSongs"`
	TotalLength int `json:"totalLength"`
}

/* Context struct for the Index page */
type Context struct {
	MessageTitle string 
	Message string
}

/* Service route struct */
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

/* Needed routes */
var routes = Routes{
	
	/* Route for index/services page */
	Route{
        "Index",
        "GET",
        "/",
        Index,
    },
	
	/* Route for getting all songs in database */
	Route{
        "AllSongs",
        "GET",
        "/songs",
        getAllSongs,
    },
	
	/* Route for getting songs by song artist */
    Route{
        "SongByArtist",
        "GET",
        "/songs/artist/{songArtist}",
        getSongByArtist,
    },
	
	/* Route for getting songs by song name */
    Route{
        "SongByName",
        "GET",
        "/songs/name/{songName}",
        getSongByName,
    },
	
	/* Route for getting songs by song genre */
    Route{
        "SongByGenre",
        "GET",
        "/songs/genre/{songGenre}",
        getSongByGenre,
    },
	
	/* Route for getting songs by length range */
	Route{
        "SongByLength",
        "GET",
        "/songs/length/{minLength}/{maxLength}",
        getSongsByLength,
    },
	
	/* Route for getting all genres in the database */
	Route{
        "AllGenres",
        "GET",
        "/genres",
        getAllGenres,
    },
	
	/* Route for getting all genres, number of songs and total length of all the songs by genre */
	Route{
        "AllGenresExtra",
        "GET",
        "/genresExtra",
        getAllGenresExtra,
    },
	
	/* Route for getting genres by genre name */
	Route{
        "GenreByName",
        "GET",
        "/genres/name/{genreName}",
        getGenreByName,
    },
}
