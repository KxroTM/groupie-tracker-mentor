package handler

import (
	"net/http"
	"text/template"
)

var error404 = template.Must(template.ParseFiles("./src/templates/404.html"))
var spotifyerror = template.Must(template.ParseFiles("./src/templates/spotify.html"))

func Error404Handler(w http.ResponseWriter, r *http.Request) {
	err := error404.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SpotifyErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := spotifyerror.ExecuteTemplate(w, "spotify.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
