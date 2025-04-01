package search

import (
	"encoding/json"
	"io"
	"net/http"
	"trev/spot/v2/spotify"
	albums "trev/spot/v2/spotify/album"
	artists "trev/spot/v2/spotify/artist"
	"trev/spot/v2/spotify/audiobook"
	"trev/spot/v2/spotify/episode"
	"trev/spot/v2/spotify/playlist"
	"trev/spot/v2/spotify/show"
	"trev/spot/v2/spotify/track"
)

const API_URL = "https://api.spotify.com/v1/search"

func Search[T []spotify.SpotifyType](client *http.Client, query string, types T, auth http.Header) (*T, error) {
	params := make(map[string]string)
	params["q"] = query

	typesString := ""
	for _, kind := range types {
		typesString += kind.SpotifyType() + ","
	}

	params["type"] = typesString

	resp, err := spotify.GetFormAuthed(client, API_URL, params, auth)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type PaginatedObject[T any] = spotify.PaginatedObject[T]

	type SearchResponse struct {
		Tracks     *PaginatedObject[track.Track]         `json:"tracks,omitempty"`
		Artists    *PaginatedObject[artists.Artist]      `json:"artists,omitempty"`
		Albums     *PaginatedObject[albums.Album]        `json:"albums,omitempty"`
		Playlists  *PaginatedObject[playlist.Playlist]   `json:"playlists,omitempty"`
		Shows      *PaginatedObject[show.Show]           `json:"shows,omitempty"`
		Episodes   *PaginatedObject[episode.Episode]     `json:"episodes,omitempty"`
		Audiobooks *PaginatedObject[audiobook.Audiobook] `json:"audiobooks,omitempty"`
	}

	if resp.StatusCode == 200 {
		searchResponse := SearchResponse{}
		if err := json.Unmarshal(body, &searchResponse); err != nil {
			return nil, err
		}
		searchResults := make([]spotify.SpotifyType, 0, 10)

		if searchResponse.Albums != nil {
			for _, album := range searchResponse.Albums.Items {
				searchResults = append(searchResults, album)
			}
		}
	}
}
