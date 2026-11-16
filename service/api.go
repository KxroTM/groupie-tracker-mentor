package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"groupie-tracker-mentor/models"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var ApiToken string

var client = &http.Client{
	Timeout: 10 * time.Second,
}

const maxRetries = 5

func doWithRetry(req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		resp, err := client.Do(req)
		if err != nil {
			return resp, err
		}

		if resp.StatusCode == http.StatusUnauthorized && strings.HasPrefix(req.Header.Get("Authorization"), "Bearer ") && attempt < maxRetries {
			resp.Body.Close()

			newToken, err := GetSpotifyValidToken(ClientID, ClientSecret)
			if err != nil {
				return resp, err
			}
			ApiToken = newToken
			req.Header.Set("Authorization", "Bearer "+newToken)

			if err := resetBody(req); err != nil {
				return resp, err
			}
			continue
		}

		if resp.StatusCode != http.StatusTooManyRequests || attempt >= maxRetries {
			return resp, nil
		}

		wait := retryAfterDuration(resp)
		log.Printf("rate limit Spotify atteint, réessaie dans %s secondes", wait)
		resp.Body.Close()
		time.Sleep(wait)

		if err := resetBody(req); err != nil {
			return resp, err
		}
	}
}

func resetBody(req *http.Request) error {
	if req.GetBody == nil {
		return nil
	}
	body, err := req.GetBody()
	if err != nil {
		return err
	}
	req.Body = body
	return nil
}

func retryAfterDuration(resp *http.Response) time.Duration {
	seconds, err := strconv.Atoi(resp.Header.Get("Retry-After"))
	if err != nil || seconds <= 0 {
		return 2 * time.Second
	}
	return time.Duration(seconds) * time.Second
}

func GetSpotifyValidToken(clientID, clientSecret string) (string, error) {

	postBody := bytes.NewBufferString("grant_type=client_credentials")

	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", postBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))

	resp, err := doWithRetry(request)
	if err != nil {
		return "", SpotifyError
	}

	if resp.StatusCode != http.StatusOK {
		return "", SpotifyError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", SpotifyError
	}

	var jsonToken models.TokenApi

	err = json.Unmarshal(body, &jsonToken)
	if err != nil {
		return "", SpotifyError
	}
	return jsonToken.AccessToken, nil
}

func GetJulId(apitoken string) (string, error) {

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search?q=JUL&type=artist", bytes.NewBufferString(""))
	req.Header.Add("Authorization", "Bearer "+apitoken)

	resp, err := doWithRetry(req)
	if err != nil {
		return "", SpotifyError
	}

	if resp.StatusCode != http.StatusOK {
		return "", SpotifyError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", SpotifyError
	}

	return string(body), nil
}

func GetArtistAlbumIDs(apitoken, artistID string) ([]string, error) {
	var ids []string

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/albums?include_groups=album,single&limit=10"

	for endpoint != "" {
		req, err := http.NewRequest("GET", endpoint, bytes.NewBufferString(""))
		req.Header.Add("Authorization", "Bearer "+apitoken)

		resp, err := doWithRetry(req)
		if err != nil {
			return nil, SpotifyError
		}

		if resp.StatusCode != http.StatusOK {
			return nil, SpotifyError
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, SpotifyError
		}

		var page models.SimplifiedAlbumsResponse
		err = json.Unmarshal(body, &page)
		if err != nil {
			return nil, SpotifyError
		}

		for _, item := range page.Items {
			ids = append(ids, item.ID)
		}

		if page.Next != nil {
			endpoint = *page.Next
		} else {
			endpoint = ""
		}
	}

	return ids, nil
}

func GetAlbumsDetails(apitoken string, albumIDs []string) ([]models.SpotifyAlbumResponse, error) {
	var albums []models.SpotifyAlbumResponse

	for _, id := range albumIDs {
		endpoint := "https://api.spotify.com/v1/albums/" + id
		req, err := http.NewRequest("GET", endpoint, bytes.NewBufferString(""))
		req.Header.Add("Authorization", "Bearer "+apitoken)

		resp, err := doWithRetry(req)
		if err != nil {
			return nil, SpotifyError
		}

		if resp.StatusCode != http.StatusOK {
			return nil, SpotifyError
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, SpotifyError
		}

		var album models.SpotifyAlbumResponse
		err = json.Unmarshal(body, &album)
		if err != nil {
			return nil, SpotifyError
		}

		albums = append(albums, album)
	}

	return albums, nil
}
