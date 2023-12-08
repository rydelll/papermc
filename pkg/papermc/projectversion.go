package papermc

import (
	"encoding/json"
	"fmt"
	"io"
)

type projectVersionService service

type projectVersionResponse struct {
	ProjectId     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

func (srv *projectVersionService) GetLatest(project Project) (string, error) {
	url := fmt.Sprintf(srv.client.baseURL+string(srv.endpoint), project)
	resp, err := srv.client.httpClient.Get(url)
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
