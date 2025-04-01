package show

import "trev/spot/v2/spotify"

// Show represents a simplified show
type Show struct {
	AvailableMarkets   []string             `json:"available_markets"`
	Copyrights         []CopyrightObject    `json:"copyrights"`
	Description        string               `json:"description"`
	HTMLDescription    string               `json:"html_description"`
	Explicit           bool                 `json:"explicit"`
	ExternalURLs       spotify.ExternalURLs `json:"external_urls"`
	Href               string               `json:"href"`
	ID                 string               `json:"id"`
	Images             []spotify.Image      `json:"images"`
	IsExternallyHosted bool                 `json:"is_externally_hosted"`
	Languages          []string             `json:"languages"`
	MediaType          string               `json:"media_type"`
	Name               string               `json:"name"`
	Publisher          string               `json:"publisher"`
	Type               string               `json:"type"`
	URI                string               `json:"uri"`
	TotalEpisodes      int                  `json:"total_episodes"`
}

// CopyrightObject represents a copyright statement
type CopyrightObject struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
