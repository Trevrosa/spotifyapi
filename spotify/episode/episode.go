package episode

import "trev/spot/v2/spotify"

// Episode represents a simplified episode
type Episode struct {
	AudioPreviewURL      string               `json:"audio_preview_url"`
	Description          string               `json:"description"`
	HTMLDescription      string               `json:"html_description"`
	DurationMs           int                  `json:"duration_ms"`
	Explicit             bool                 `json:"explicit"`
	ExternalURLs         spotify.ExternalURLs `json:"external_urls"`
	Href                 string               `json:"href"`
	ID                   string               `json:"id"`
	Images               []spotify.Image      `json:"images"`
	IsExternallyHosted   bool                 `json:"is_externally_hosted"`
	IsPlayable           bool                 `json:"is_playable"`
	Language             string               `json:"language"`
	Languages            []string             `json:"languages"`
	Name                 string               `json:"name"`
	ReleaseDate          string               `json:"release_date"`
	ReleaseDatePrecision string               `json:"release_date_precision"`
	ResumePoint          ResumePoint   `json:"resume_point"`
	Type                 string               `json:"type"`
	URI                  string               `json:"uri"`
}

// ResumePoint represents a resume point for an episode
type ResumePoint struct {
	FullyPlayed      bool `json:"fully_played"`
	ResumePositionMs int  `json:"resume_position_ms"`
}
