package main

import "net/http"

type Song struct {
	Id int `json:"ID"`
	Artist string `json:"artist"`
	Song string `json:"song"`
	Genre string `json:"genre"`
	Length int `json:"length"`
}

type Genre struct {
	Id int `json:"ID"`
	Name string `json:"name"`
}

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
        "Index",
        "GET",
        "/",
        Index,
    },
	Route{
        "AllSongs",
        "GET",
        "/songs",
        getAllSongs,
    },
    Route{
        "SongByArtist",
        "GET",
        "/songs/artist/{songArtist}",
        getSongByArtist,
    },
    Route{
        "SongByName",
        "GET",
        "/songs/name/{songName}",
        getSongByName,
    },
    Route{
        "SongByGenre",
        "GET",
        "/songs/genre/{songGenre}",
        getSongByGenre,
    },
	Route{
        "AllGenres",
        "GET",
        "/genres",
        getAllGenres,
    },
	Route{
        "GenreByName",
        "GET",
        "/genres/name/{genreName}",
        getGenreByName,
    },
}
