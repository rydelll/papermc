package papermc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const versionEndpoint = "/projects/%s"

// A VersionService is used to get information about a PaperMC project version.
type VersionService struct {
	project Project
	baseURL string
	client  *http.Client
}

// NewVersionService for a PaperMC project.
func NewVersionService(project Project, baseURL string, client *http.Client) *VersionService {
	return &VersionService{
		project: project,
		baseURL: baseURL,
		client:  client,
	}
}

// GetLatest version for a PaperMC project.
func (s *VersionService) GetLatest() (string, error) {
	vers, err := s.List()
	if err != nil {
		return "", err
	}

	return vers[len(vers)-1], nil
}

type versionResponse struct {
	ProjectID     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

// List versions for a PaperMC project.
func (s *VersionService) List() ([]string, error) {
	url := fmt.Sprintf(s.baseURL+versionEndpoint, s.project)
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

	vr := versionResponse{}
	err = json.Unmarshal(body, &vr)
	if err != nil {
		return nil, fmt.Errorf("parse response body: %w", err)
	}

	return vr.Versions, nil
}
