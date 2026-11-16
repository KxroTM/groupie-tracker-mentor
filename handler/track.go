package handler

import (
	"groupie-tracker-mentor/service"
	"net/http"
	"strings"
	"text/template"
)

func TrackHandler(w http.ResponseWriter, r *http.Request) {
	var track = template.Must(template.ParseFiles("./src/templates/song.html"))
	var id string
	path := r.URL.Path
	parms := strings.Split(path, "/")
	if parms[2] == "static" {
		Error404Handler(w, r)
		return
	}
	id = parms[2]

	song, err := service.GetSongById(id)
	if err != nil {
		if err == service.SpotifyError {
			SpotifyErrorHandler(w, r)
			return
		} else {
			Error404Handler(w, r)
			return
		}
	}

	err = track.ExecuteTemplate(w, "song.html", song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
