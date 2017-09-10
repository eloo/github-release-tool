package models

import "time"

// Release representing a Github release
type Release struct {
	Name        string    `json:"name"`
	TagName     string    `json:"tag_name"`
	Assets      []Asset   `json:"assets"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// Asset representing a Github asset
type Asset struct {
	URL         string    `json:"url"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updated_at"`
	DownloadURL string    `json:"browser_download_url"`
}
