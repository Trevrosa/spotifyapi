package audiobook

import "trev/spot/v2/spotify"

// Audiobook represents a simplified audiobook
type Audiobook struct {
	Authors          []NarratorAuthor     `json:"authors"`
	AvailableMarkets []string             `json:"available_markets"`
	Copyrights       []CopyrightObject    `json:"copyrights"`
	Description      string               `json:"description"`
	HTMLDescription  string               `json:"html_description"`
	Edition          string               `json:"edition"`
	Explicit         bool                 `json:"explicit"`
	ExternalURLs     spotify.ExternalURLs `json:"external_urls"`
	Href             string               `json:"href"`
	ID               string               `json:"id"`
	Images           []spotify.Image      `json:"images"`
	Languages        []string             `json:"languages"`
	MediaType        string               `json:"media_type"`
	Name             string               `json:"name"`
	Narrators        []Narrator           `json:"narrators"`
	Publisher        string               `json:"publisher"`
	Type             string               `json:"type"`
	URI              string               `json:"uri"`
	TotalChapters    int                  `json:"total_chapters"`
}

// NarratorAuthor represents an author
type NarratorAuthor struct {
	Name string `json:"name"`
}

// Narrator represents a narrator
type Narrator struct {
	Name string `json:"name"`
}

type CopyrightObject struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
