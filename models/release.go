package models

import "time"

type Release struct {
	Name        string    `json:"name"`
	TagName     string    `json:"tag_name"`
	Assets      []Asset   `json:"assets"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Asset struct {
	Url         string    `json:"url"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updated_at"`
	DownloadUrl string    `json:"browser_download_url"`
}

type DownloadCandidate struct {
	File     Asset
	Checksum Asset
}
