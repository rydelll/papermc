package client

import "net/http"

// URL format: /projects/{project}
const endpointProjectVersion versionEndpoint = "/projects/%s"

type versionEndpoint string

type VersionService struct {
	client   *http.Client
	project  Project
	endpoint versionEndpoint
}

func NewVersionService(client *http.Client, project Project) *VersionService {
	return &VersionService{
		client:   client,
		project:  project,
		endpoint: endpointProjectVersion,
	}
}

func (s *VersionService) Get(ver string) (string, error) {
	return "specific version", nil
}

func (s *VersionService) GetLatest() (string, error) {
	return "latest version", nil
}

func (s *VersionService) List() ([]string, error) {
	vers := make([]string, 2)
	vers[0] = "first version"
	vers[1] = "second version"
	return vers, nil
}
