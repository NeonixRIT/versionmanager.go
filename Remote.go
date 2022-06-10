package versionmanager

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type remote struct {
	Author    string
	Name      string
	Version   version
	URL       string
	Separator string
	Data      Release
}

// getReleaseData attempts to get a projects data from GitHub release API request
func (r *remote) getReleaseData() (Release, error) {
	var release Release

	projectURL := "https://api.github.com/repos/" + r.Author + "/" + r.Name + "/releases/latest"
	resp, err := http.Get(projectURL)
	if err != nil {
		return release, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return release, err
	}

	if err := json.Unmarshal(body, &release); err != nil {
		return release, err
	}

	if release.ErrorMessage != "" {
		return Release{}, errors.New("Release not found: " + r.Name)
	}

	return release, nil
}

// refresh attempts to update
func (r *remote) refresh() {
	release, err := r.getReleaseData()
	if err != nil {
		return
	}
	r.Data = release
	r.Version = makeVersion(release.Version, r.Separator)
}

// makeRemote does something
func makeRemote(author string, projectName string, separator string) remote {
	remoteVar := remote{
		Author:    author,
		Name:      projectName,
		Separator: separator,
		URL:       "https://api.github.com/repos/" + author + "/" + projectName + "/releases/latest"}
	data, err := remoteVar.getReleaseData()
	if err != nil {
		return remote{}
	}
	remoteVar.Data = data
	remoteVar.Version = makeVersion(data.Version, separator)
	return remoteVar
}
