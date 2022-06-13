package versionmanagergo

// Release represents a parsed GitHub release API JSON. This does not contain all the information contained in the JSON
type Release struct {
	URL          string `json:"url"`
	Version      string `json:"tag_name"`
	FriendlyName string `json:"name"`
	Prerelease   bool   `json:"prerelease"`
	PublishDate  string `json:"published_at"`
	Description  string `json:"body"`
	ErrorMessage string `json:"message"`
}
