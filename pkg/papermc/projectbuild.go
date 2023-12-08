package papermc

import (
	"encoding/json"
	"fmt"
	"io"
)

type projectBuildService service

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

func (srv *projectBuildService) GetLatest(project Project, version string) (BuildInfo, error) {
	url := fmt.Sprintf(srv.client.baseURL+string(srv.endpoint), project, version)
	resp, err := srv.client.httpClient.Get(url)
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
