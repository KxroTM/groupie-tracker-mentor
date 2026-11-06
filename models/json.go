package models

var JulId = "3IW7ScrzXmPvZhB27hmfgy"

type TokenApi struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type SimplifiedAlbumsResponse struct {
	Items []struct {
		ID string `json:"id"`
	} `json:"items"`
	Next *string `json:"next"`
}

type SeveralAlbumsResponse struct {
	Albums []SpotifyAlbumResponse `json:"albums"`
}

type SpotifyAlbumResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	AlbumType    string `json:"album_type"`
	ReleaseDate  string `json:"release_date"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Tracks struct {
		Items []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			DurationMs int    `json:"duration_ms"`
			Artists    []struct {
				Name string `json:"name"`
			} `json:"artists"`
		} `json:"items"`
	} `json:"tracks"`
}
