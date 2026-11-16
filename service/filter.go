package service

import (
	"groupie-tracker-mentor/models"
	"slices"
	"strings"
)

func SortByDate(tracks []models.Song) models.Song {
	var result models.Song
	for _, track := range tracks {
		if track.ReleaseDate > result.ReleaseDate {
			result = track
		}
	}
	return result
}

func GetFeats(tracks []models.Song) []models.Song {
	var result []models.Song

	for _, track := range tracks {
		if len(track.Artists) > 1 {
			result = append(result, track)
		}
	}
	return result
}

func GetType(tracks []models.Song, songtype string) []models.Song {
	var result []models.Song

	for _, track := range tracks {
		if track.Type == songtype {
			result = append(result, track)
		}
	}
	return result
}

func GetSongByYear(tracks []models.Song, year string) []models.Song {
	var result []models.Song

	for _, track := range tracks {
		date := strings.Split(track.ReleaseDate, "-")[0]
		if date == year {
			result = append(result, track)
		}
	}
	return result
}

func GetAllYears(tracks []models.Song) []string {
	var result []string

	for _, track := range tracks {
		if !slices.Contains(result, strings.Split(track.ReleaseDate, "-")[0]) {
			result = append(result, strings.Split(track.ReleaseDate, "-")[0])
		}
	}
	slices.Sort(result)
	return result
}
