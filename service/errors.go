package service

import "errors"

var (
	// 404 Not Found
	ErrSongNotFound = errors.New("Son introuvable")
	// 5xx Spotify error
	SpotifyError = errors.New("Spotify injoignable")
)
