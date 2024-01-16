package client

import (
	"encoding/json"
	"fmt"
	"io"
)

// URL format: /projects/{project}
const versionEndpoint string = "/projects/%s"

type versionResponse struct {
	ProjectId     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

// TODO
func (s *ProjectService) GetVersion(ver string) (string, error) {
	return "", nil
}

// TODO
func (s *ProjectService) GetLatestVersion() (string, error) {
	url := fmt.Sprintf(s.baseURL+versionEndpoint, s.project)
	resp, err := s.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("could not reach %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body from %s: %w", url, err)
	}

	vr := versionResponse{}
	err = json.Unmarshal(body, &vr)
	if err != nil {
		return "", fmt.Errorf("failed parsing response body from %s: %w", url, err)
	}

	return vr.Versions[len(vr.Versions)-1], nil
}

// TODO
func (s *ProjectService) ListVersions() ([]string, error) {
	return nil, nil
}
