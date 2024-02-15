package papermc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	buildListEndpoint = "/projects/%s/versions/%s"
	buildInfoEndpoint = "/projects/%s/versions/%s/builds/%d"
)

// A BuildService is used to get information about a PaperMC project build.
type BuildService struct {
	project Project
	baseURL string
	client  *http.Client
}

// NewBuildService for a PaperMC project.
func NewBuildService(project Project, baseURL string, client *http.Client) *BuildService {
	return &BuildService{
		project: project,
		baseURL: baseURL,
		client:  client,
	}
}

// ProjectInfo information required for downloading a JAR.
type ProjectInfo struct {
	Version  string
	Build    int
	JAR      string
	Checksum string
}

type buildInfoResponse struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Build       int    `json:"build"`
	Time        string `json:"time"`
	Channel     string `json:"channel"`
	Promoted    bool   `json:"promoted"`
	Changes     []struct {
		Commit  string `json:"commit"`
		Summary string `json:"summary"`
		Message string `json:"message"`
	} `json:"changes"`
	Download struct {
		Application struct {
			Name   string `json:"name"`
			SHA256 string `json:"sha256"`
		} `json:"application"`
	} `json:"downloads"`
}

// Get project information for a build of a PaperMC project version.
func (s *BuildService) Get(ver string, build int) (ProjectInfo, error) {
	url := fmt.Sprintf(s.baseURL+buildInfoEndpoint, s.project, ver, build)
	resp, err := s.client.Get(url)
	if err != nil {
		return ProjectInfo{}, fmt.Errorf("fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ProjectInfo{}, fmt.Errorf("response: HTTP status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProjectInfo{}, fmt.Errorf("read response body: %w", err)
	}

	br := buildInfoResponse{}
	err = json.Unmarshal(body, &br)
	if err != nil {
		return ProjectInfo{}, fmt.Errorf("parse response body: %w", err)
	}

	return ProjectInfo{
		Version:  br.Version,
		Build:    br.Build,
		JAR:      br.Download.Application.Name,
		Checksum: br.Download.Application.SHA256,
	}, nil
}

// GetLatest project information for latest build of a PaperMC project version.
func (s *BuildService) GetLatest(ver string) (ProjectInfo, error) {
	builds, err := s.List(ver)
	if err != nil {
		return ProjectInfo{}, err
	}

	project, err := s.Get(ver, builds[len(builds)-1])
	if err != nil {
		return ProjectInfo{}, err
	}

	return project, nil
}

type buildListResponse struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []int  `json:"builds"`
}

// List builds for a PaperMC project version.
func (s *BuildService) List(ver string) ([]int, error) {
	url := fmt.Sprintf(s.baseURL+buildListEndpoint, s.project, ver)
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response: HTTP status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	br := buildListResponse{}
	err = json.Unmarshal(body, &br)
	if err != nil {
		return nil, fmt.Errorf("parse response body: %w", err)
	}

	return br.Builds, nil
}
