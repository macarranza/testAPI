package main

/* Html template for the Index/Services page */
const doc = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Index</title>
	<style>
		body {background-color: #D3D3D3;}
		h1, h2, h3 { color: white; background-color: black;}
		div {margin:30px;}
	</style>	
</head>
<body>
	<h1 align="center">{{.MessageTitle}}</h1>
	<br><br>
	<div>
		<h2>Services</h2>
		<b>http://localhost:8080/songs - To get all songs in database</b><br>
		<b>http://localhost:8080/songs/name/{songName} - To search songs by song name</b><br>
		<b>http://localhost:8080/songs/artist/{songArtist} - To search songs by artist</b><br>
		<b>http://localhost:8080/songs/genre/{songGenre} - To search songs by genre</b><br>
		<b>http://localhost:8080/songs/length/{minLength}/{maxLength} - To search songs by length range</b><br>
		<b>http://localhost:8080/genres - To get all genres in database</b><br>
		<b>http://localhost:8080/genresExtra - To get number of songs, and total length of all songs by genre (All genres)</b><br>
		<b>http://localhost:8080/genres/name/{genreName} - To search genres by genre name</b><br>
	</div>
	<h3 align="center">{{.Message}}</h3>
</body>
</html>
`
