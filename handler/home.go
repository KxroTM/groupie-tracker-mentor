package handler

import (
	"groupie-tracker-mentor/models"
	"groupie-tracker-mentor/service"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var home = template.Must(template.ParseFiles("./src/templates/home.html"))
	if r.URL.Path != "/" {
		Error404Handler(w, r)
		return
	}

	songs := service.SongCache

	songType := r.URL.Query().Get("type")
	year := r.URL.Query().Get("year")
	feat := r.URL.Query().Get("feat")

	if year != "" {
		songs = service.GetSongByYear(songs, year)
	}
	if songType != "" {
		songs = service.GetType(songs, songType)
	}
	if feat == "true" {
		songs = service.GetFeats(songs)
	}

	if !service.SpotifyState {
		SpotifyErrorHandler(w, r)
		return
	}

	data := models.HomeData{
		Songs:    songs,
		LastSong: service.GetLastSong(),
		Years:    service.GetAllYears(service.SongCache),
	}

	err := home.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
