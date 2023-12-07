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

func (pvs *projectVersionService) GetLatest(project Project) (string, error) {
	url := fmt.Sprintf(string(endpointProjectVersion), project)
	resp, err := pvs.client.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("could not reach %s", url)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body from %s", url)
	}

	vr := projectVersionResponse{}
	err = json.Unmarshal(body, &vr)
	if err != nil {
		return "", fmt.Errorf("failed parsing response body from %s", url)
	}

	return vr.Versions[len(vr.Versions)-1], nil
}
