package models

type Album struct {
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
