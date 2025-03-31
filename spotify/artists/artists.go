package artists

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"trev/spot/v2/spotify"
)

const API_URL = "https://api.spotify.com/v1/artists/"

type Image struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Followers struct {
	Href  *string `json:"href"`
	Total int     `json:"total"`
}

type Artist struct {
	Name         string            `json:"name"`
	Genres       []string          `json:"genres"`
	Followers    Followers         `json:"followers"`
	Popularity   int               `json:"popularity"`
	Images       []Image           `json:"images"`
	Type         string            `json:"type"`
	Id           string            `json:"id"`
	Uri          string            `json:"uri"`
	Href         string            `json:"href"`
	ExternalUrls map[string]string `json:"external_urls"`
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

// TODO: use enums where possible

type Album struct {
	Name                 string            `json:"name"`
	TotalTracks          int               `json:"total_tracks"`
	Artists              []SimpleArtist    `json:"artists"`
	AlbumGroup           string            `json:"album_group"`
	AlbumType            string            `json:"album_type"`
	AvailableMarkets     []string          `json:"available_markets"`
	Images               []Image           `json:"images"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Restrictions         map[string]string `json:"restrictions"`
	Type                 string            `json:"type"`
	Id                   string            `json:"id"`
	Href                 string            `json:"href"`
	Uri                  string            `json:"uri"`
	ExternalUrls         map[string]string `json:"external_urls"`
}

type SimpleArtist struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	ExternalUrls map[string]string `json:"external_urls"`
	Id           string            `json:"id"`
	Href         string            `json:"href"`
	Uri          string            `json:"uri"`
}

type PaginatedAlbums struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
	Items    []Album `json:"items"`
}

var ALBUM_PARAMS = map[string]string{"include_groups": "album"}

func GetArtistAlbums(client *http.Client, artistId string, auth http.Header) (*[]Album, error) {
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
		items := PaginatedAlbums{}
		if err := json.Unmarshal(body, &items); err != nil {
			return nil, err
		}
		return &items.Items, nil
	} else {
		fmt.Printf("request error %s\n", resp.Status)
		return nil, spotify.ToSpotifyError(body)
	}
}
