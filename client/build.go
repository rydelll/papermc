package client

import "net/http"

// URL format: /projects/{project}/versions/{version}/builds
const endpointProjectBuild buildEndpoint = "/projects/%s/versions/%s/builds"

type buildEndpoint string

type BuildService struct {
	client   *http.Client
	project  Project
	endpoint buildEndpoint
}

func NewBuildService(client *http.Client, project Project) *BuildService {
	return &BuildService{
		client:   client,
		project:  project,
		endpoint: endpointProjectBuild,
	}
}

func (s *BuildService) Get(ver string) (string, error) {
	return "specific build", nil
}

func (s *BuildService) GetLatest() (string, error) {
	return "latest build", nil
}

func (s *BuildService) List() ([]string, error) {
	vers := make([]string, 2)
	vers[0] = "first build"
	vers[1] = "second build"
	return vers, nil
}
