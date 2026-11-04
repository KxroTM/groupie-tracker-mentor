package main

import (
	"fmt"
	"groupie-tracker-mentor/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("src/styles"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", handler.HomeHandler)
	fmt.Println("Server is runing on http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
}
