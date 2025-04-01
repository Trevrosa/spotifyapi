package albums

import (
	"trev/spot/v2/spotify"
	"trev/spot/v2/spotify/simple"
)

// TODO: use enums where possible

type Album struct {
	Name                 string                `json:"name"`
	TotalTracks          int                   `json:"total_tracks"`
	Artists              []simple.SimpleArtist `json:"artists"`
	AlbumGroup           string                `json:"album_group"`
	AlbumType            string                `json:"album_type"`
	AvailableMarkets     []string              `json:"available_markets"`
	Images               []spotify.Image       `json:"images"`
	ReleaseDate          string                `json:"release_date"`
	ReleaseDatePrecision string                `json:"release_date_precision"`
	Restrictions         *spotify.Restrictions         `json:"restrictions"`
	Type                 string                `json:"type"`
	Id                   string                `json:"id"`
	Href                 string                `json:"href"`
	Uri                  string                `json:"uri"`
	ExternalUrls         map[string]string     `json:"external_urls"`
}

func (album Album) SpotifyType() string {
	return "album"
}

type PaginatedAlbums = spotify.PaginatedObject[Album]
