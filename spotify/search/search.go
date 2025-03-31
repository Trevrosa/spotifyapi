package search

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"trev/spot/v2/spotify"
)

type QueryType string

const (
	Album     QueryType = "album"
	Artist    QueryType = "artist"
	Playlist  QueryType = "playlist"
	Track     QueryType = "track"
	Show      QueryType = "show"
	Episode   QueryType = "episode"
	Audiobook QueryType = "audiobook"
)

const API_URL = "https://api.spotify.com/v1/search"

func Search[T spotify.SpotifyType](client *http.Client, query string, types []QueryType, auth http.Header) (*T, error) {
	params := make(map[string]string)
	params["q"] = query

	typesString := ""
	for _, kind := range types {
		typesString += string(kind) + ","
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

	if resp.StatusCode == 200 {
		json.Unmarshal()

	}
}
