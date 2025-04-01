package playlist

import "trev/spot/v2/spotify"

// Playlist represents a simplified playlist
type Playlist struct {
	Collaborative bool                 `json:"collaborative"`
	Description   string               `json:"description"`
	ExternalURLs  spotify.ExternalURLs `json:"external_urls"`
	Href          string               `json:"href"`
	ID            string               `json:"id"`
	Images        []spotify.Image      `json:"images"`
	Name          string               `json:"name"`
	Owner         PlaylistOwner        `json:"owner"`
	Public        bool                 `json:"public"`
	SnapshotID    string               `json:"snapshot_id"`
	Tracks        PlaylistTracks       `json:"tracks"`
	Type          string               `json:"type"`
	URI           string               `json:"uri"`
}

// PlaylistOwner represents a playlist owner
type PlaylistOwner struct {
	ExternalURLs spotify.ExternalURLs `json:"external_urls"`
	Href         string               `json:"href"`
	ID           string               `json:"id"`
	Type         string               `json:"type"`
	URI          string               `json:"uri"`
	DisplayName  string               `json:"display_name"`
}

// PlaylistTracks represents tracks in a playlist
type PlaylistTracks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
