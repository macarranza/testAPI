# testAPI

Installation:
- Move the API folder to your GOPATH folder (your GO workspace) 
- Execute: go install /yourGOPATH/API
- Execute the API.exe file created in the bin folder in your GOPATH folder 

Usage: After executing the API:
- http://localhost:8080/ - Index/Welcome/Services
- http://localhost:8080/songs - To get all songs in database
- http://localhost:8080/songs/name/{songName} - To search songs by song name
- http://localhost:8080/songs/artist/{songArtist} - To search songs by artist
- http://localhost:8080/songs/genre/{songGenre} - To search songs by genre
- http://localhost:8080/songs/length/{minLength}/{maxLength} - To search songs by length range
- http://localhost:8080/genres - To get all genres in database
- http://localhost:8080/genresExtra - To get number of songs, and total length of all songs by genre (All genres) 
- http://localhost:8080/genres/name/{genreName} - To search genres by genre name
