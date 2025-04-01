package artists

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"trev/spot/v2/spotify"
	"trev/spot/v2/spotify/album"
)

const API_URL = "https://api.spotify.com/v1/artists/"

type Followers struct {
	Href  *string `json:"href"`
	Total int     `json:"total"`
}

type Artist struct {
	Name         string               `json:"name"`
	Genres       []string             `json:"genres"`
	Followers    Followers            `json:"followers"`
	Popularity   int                  `json:"popularity"`
	Images       []spotify.Image      `json:"images"`
	Type         string               `json:"type"`
	Id           string               `json:"id"`
	Uri          string               `json:"uri"`
	Href         string               `json:"href"`
	ExternalUrls spotify.ExternalURLs `json:"external_urls"`
}

func GetArtist(client *http.Client, id string, auth http.Header) (*Artist, error) {
	resp, err := spotify.GetAuthed(client, API_URL+id, auth)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		artist := Artist{}
		if err := json.Unmarshal(body, &artist); err != nil {
			return nil, err
		}
		return &artist, nil
	} else {
		fmt.Printf("request error %s\n", resp.Status)
		return nil, spotify.ToSpotifyError(body)
	}
}

func GetArtists(client *http.Client, ids []string, auth http.Header) (*[]Artist, error) {
	type multipleArtists struct {
		Artists []Artist
	}

	parsedIds := strings.Join(ids, ",")
	resp, err := spotify.GetAuthed(client, API_URL+parsedIds, auth)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		multipleArtists := multipleArtists{}
		if err := json.Unmarshal(body, &multipleArtists); err != nil {
			return nil, err
		}
		return &multipleArtists.Artists, nil
	} else {
		fmt.Printf("request error %s\n", resp.Status)
		return nil, spotify.ToSpotifyError(body)
	}
}

var ALBUM_PARAMS = map[string]string{"include_groups": "album"}

func GetArtistAlbums(client *http.Client, artistId string, auth http.Header) (*[]albums.Album, error) {
	resp, err := spotify.GetFormAuthed(client, fmt.Sprintf("%s%s/%s", API_URL, artistId, "albums"), ALBUM_PARAMS, auth)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		items := albums.PaginatedAlbums{}
		if err := json.Unmarshal(body, &items); err != nil {
			return nil, err
		}
		return &items.Items, nil
	} else {
		fmt.Printf("request error %s\n", resp.Status)
		return nil, spotify.ToSpotifyError(body)
	}
}
