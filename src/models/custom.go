package models

// DownloadCandidate with an asset and optionally a checksum asset
type DownloadCandidate struct {
	File     Asset
	Checksum Asset
}
