package models

type Song struct {
	ID          string
	Name        string
	Image       string
	ReleaseDate string
	Artists     []string
	Duration    string
	Type        string
	AlbumName   string
	SpotifyURL  string
}
