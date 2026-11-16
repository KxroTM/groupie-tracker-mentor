package main

import (
	"fmt"
	"groupie-tracker-mentor/handler"
	"groupie-tracker-mentor/service"
	"net/http"
)

func main() {
	service.GenDotEnv()
	token, err := service.GetSpotifyValidToken(service.ClientID, service.ClientSecret)
	if err != nil {
		service.SpotifyState = false
	} else {
		service.ApiToken = token
		songs, err := service.GetAllJulSong()
		if err != nil {
			service.SpotifyState = false
		}
		service.SongCache = songs
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("src/styles"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/tracks/", handler.TrackHandler)
	mux.HandleFunc("/notfound", handler.Error404Handler)
	mux.HandleFunc("/spotifyError", handler.SpotifyErrorHandler)

	fmt.Println("Server is runing on http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
}
