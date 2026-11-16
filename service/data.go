package service

import (
	"fmt"
	"groupie-tracker-mentor/models"
)

func GetAllJulSong() ([]models.Song, error) {
	albumIDs, err := GetArtistAlbumIDs(ApiToken, models.JulId)
	if err != nil {
		return nil, SpotifyError
	}

	albums, err := GetAlbumsDetails(ApiToken, albumIDs)
	if err != nil {
		return nil, SpotifyError
	}

	var songs []models.Song
	for _, album := range albums {
		image := ""
		if len(album.Images) > 0 {
			image = album.Images[0].URL
		}

		for _, track := range album.Tracks.Items {
			var artistNames []string
			for _, artist := range track.Artists {
				artistNames = append(artistNames, artist.Name)
			}

			songs = append(songs, models.Song{
				ID:          track.ID,
				Name:        track.Name,
				Image:       image,
				ReleaseDate: album.ReleaseDate,
				Artists:     artistNames,
				Duration:    formatDuration(track.DurationMs),
				Type:        album.AlbumType,
				AlbumName:   album.Name,
				SpotifyURL:  album.ExternalURLs.Spotify,
			})
		}
	}

	return songs, nil
}

func GetLastSong() models.Song {
	return SortByDate(SongCache)
}

func GetSongById(id string) (models.Song, error) {
	var result models.Song

	if !SpotifyState {
		return result, SpotifyError
	}

	for _, song := range SongCache {
		if song.ID == id {
			result = song
		}
	}
	if result.ID == "" {
		return result, ErrSongNotFound
	}
	return result, nil
}

func formatDuration(ms int) string {
	totalSeconds := ms / 1000
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}
