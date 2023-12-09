package client

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	// URL format: /projects/{project}
	endpointProjectVersion endpoint = "/projects/%s"
	// URL format: /projects/{project}/versions/{version}/builds
	endpointProjectBuild endpoint = "/projects/%s/versions/%s/builds"
	// URL format: /projects/{project}/versions/{version}/builds/{build}/downloads/{jar}
	endpointProjectDownload endpoint = "/projects/%s/versions/%s/builds/%d/downloads/%s"
)

type endpoint string

type projectService struct {
	client  *Client
	project project
}

type projectVersionResponse struct {
	ProjectId     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

type projectBuildResponse struct {
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

type BuildInfo struct {
	version  string
	build    int
	jar      string
	checksum string
}

func (ps *projectService) LatestVersion() (string, error) {
	url := fmt.Sprintf(ps.client.baseURL+string(endpointProjectVersion), ps.project)
	resp, err := ps.client.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("could not reach %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body from %s: %w", url, err)
	}

	vr := projectVersionResponse{}
	err = json.Unmarshal(body, &vr)
	if err != nil {
		return "", fmt.Errorf("failed parsing response body from %s: %w", url, err)
	}

	return vr.Versions[len(vr.Versions)-1], nil
}

func (ps *projectService) LatestBuild(version string) (BuildInfo, error) {
	url := fmt.Sprintf(ps.client.baseURL+string(endpointProjectBuild), ps.project, version)
	resp, err := ps.client.httpClient.Get(url)
	if err != nil {
		return BuildInfo{}, fmt.Errorf("could not reach %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return BuildInfo{}, fmt.Errorf("failed reading response body from %s: %w", url, err)
	}

	br := projectBuildResponse{}
	err = json.Unmarshal(body, &br)
	if err != nil {
		return BuildInfo{}, fmt.Errorf("failed parsing response body from %s: %w", url, err)
	}

	return BuildInfo{
			version:  version,
			build:    br.Builds[len(br.Builds)-1].Build,
			jar:      br.Builds[len(br.Builds)-1].Download.Application.Name,
			checksum: br.Builds[len(br.Builds)-1].Download.Application.SHA256,
		},
		nil
}

func (ps *projectService) Download(info BuildInfo) error {
	out, err := os.Create(fmt.Sprintf("./%s.jar", ps.project))
	if err != nil {
		return fmt.Errorf("could not create %s.jar file: %w", ps.project, err)
	}
	defer out.Close()

	url := fmt.Sprintf(ps.client.baseURL+string(endpointProjectDownload), ps.project, info.version, info.build, info.jar)
	resp, err := ps.client.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("could not reach %s: %w", url, err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save request body to %s.jar file: %w", ps.project, err)
	}

	err = validateChecksum(ps.project, info.checksum)
	if err != nil {
		return err
	}

	return nil
}

func validateChecksum(project project, checksum string) error {
	file, err := os.Open(fmt.Sprintf("./%s.jar", project))
	if err != nil {
		return fmt.Errorf("could not read %s.jar file: %w", project, err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return fmt.Errorf("could not create SHA256 hash of %s.jar: %w", project, err)
	}

	if hex.EncodeToString(hasher.Sum(nil)) != checksum {
		return fmt.Errorf("invalid SHA256 hash, expected %s but got %x", checksum, hasher.Sum(nil))
	}

	return nil
}
