package handler

import (
	"groupie-tracker-mentor/service"
	"net/http"
	"text/template"
)

var home = template.Must(template.ParseFiles("./src/templates/home.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	songs := service.GetAllJulSong()
	err := home.ExecuteTemplate(w, "home.html", songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
