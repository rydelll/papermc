package client

import (
	"encoding/json"
	"fmt"
	"io"
)

// URL format: /projects/{project}/versions/{version}/builds
const buildEndpoint string = "/projects/%s/versions/%s/builds"

// TODO
type Build struct {
	Version  string
	Build    int
	Jar      string
	Checksum string
}

type buildResponse struct {
	ProjectId   string  `json:"project_id"`
	ProjectName string  `json:"project_name"`
	Version     string  `json:"version"`
	Builds      []build `json:"builds"`
}

type build struct {
	Build    int      `json:"build"`
	Time     string   `json:"time"`
	Channel  string   `json:"channel"`
	Promoted bool     `json:"promoted"`
	Changes  []change `json:"changes"`
	Download download `json:"downloads"`
}

type change struct {
	Commit  string `json:"commit"`
	Summary string `json:"summary"`
	Message string `json:"message"`
}

type download struct {
	Application application `json:"application"`
}

type application struct {
	Name   string `json:"name"`
	SHA256 string `json:"sha256"`
}

// TODO
func (s *ProjectService) GetBuild(ver string, build int) (Build, error) {
	return Build{}, nil
}

// TODO
func (s *ProjectService) GetLatestBuild(ver string) (Build, error) {
	url := fmt.Sprintf(s.baseURL+buildEndpoint, s.project, ver)
	resp, err := s.client.Get(url)
	if err != nil {
		return Build{}, fmt.Errorf("could not reach %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Build{}, fmt.Errorf("failed reading response body from %s: %w", url, err)
	}

	br := buildResponse{}
	err = json.Unmarshal(body, &br)
	if err != nil {
		return Build{}, fmt.Errorf("failed parsing response body from %s: %w", url, err)
	}

	return Build{
			Version:  br.Version,
			Build:    br.Builds[len(br.Builds)-1].Build,
			Jar:      br.Builds[len(br.Builds)-1].Download.Application.Name,
			Checksum: br.Builds[len(br.Builds)-1].Download.Application.SHA256,
		},
		nil
}

// TODO
func (s *ProjectService) ListBuilds(ver string) ([]Build, error) {
	return nil, nil
}
