package track

import (
	"trev/spot/v2/spotify"
	"trev/spot/v2/spotify/simple"
)

type Track struct {
	Album            simple.SimpleAlbum    `json:"album"`
	Artists          []simple.SimpleArtist `json:"artists"`
	AvailableMarkets []string              `json:"available_markets"`
	DiscNumber       int                   `json:"disc_number"`
	DurationMs       int                   `json:"duration_ms"`
	Explicit         bool                  `json:"explicit"`
	ExternalIDs      spotify.ExternalIDs   `json:"external_ids"`
	ExternalURLs     spotify.ExternalURLs  `json:"external_urls"`
	Href             string                `json:"href"`
	ID               string                `json:"id"`
	IsPlayable       bool                  `json:"is_playable"`
	LinkedFrom       *TrackLink            `json:"linked_from,omitempty"`
	Restrictions     *spotify.Restrictions `json:"restrictions,omitempty"`
	Name             string                `json:"name"`
	Popularity       int                   `json:"popularity"`
	PreviewURL       string                `json:"preview_url"`
	TrackNumber      int                   `json:"track_number"`
	Type             string                `json:"type"`
	URI              string                `json:"uri"`
	IsLocal          bool                  `json:"is_local"`
}

// TrackLink represents the track that a track was linked from
type TrackLink struct {
	ExternalURLs spotify.ExternalURLs `json:"external_urls"`
	Href         string               `json:"href"`
	ID           string               `json:"id"`
	Type         string               `json:"type"`
	URI          string               `json:"uri"`
}