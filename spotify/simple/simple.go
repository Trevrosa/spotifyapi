package simple

import "trev/spot/v2/spotify"

type SimpleArtist struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	ExternalUrls map[string]string `json:"external_urls"`
	Id           string            `json:"id"`
	Href         string            `json:"href"`
	Uri          string            `json:"uri"`
}

// SimpleAlbum represents a simplified album
type SimpleAlbum struct {
	AlbumType            string                `json:"album_type"`
	TotalTracks          int                   `json:"total_tracks"`
	AvailableMarkets     []string              `json:"available_markets"`
	ExternalURLs         spotify.ExternalURLs  `json:"external_urls"`
	Href                 string                `json:"href"`
	ID                   string                `json:"id"`
	Images               []spotify.Image       `json:"images"`
	Name                 string                `json:"name"`
	ReleaseDate          string                `json:"release_date"`
	ReleaseDatePrecision string                `json:"release_date_precision"`
	Restrictions         *spotify.Restrictions `json:"restrictions,omitempty"`
	Type                 string                `json:"type"`
	URI                  string                `json:"uri"`
	Artists              []SimpleArtist `json:"artists"`
}
