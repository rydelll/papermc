package client

import (
	"fmt"
	"net/http"
)

// URL format: /projects/{project}/versions/{version}/builds/{build}/downloads/{jar}
const endpointProjectDownload downloadEndpoint = "/projects/%s/versions/%s/builds/%d/downloads/%s"

type downloadEndpoint string

type DownloadService struct {
	client   *http.Client
	project  Project
	endpoint downloadEndpoint
}

func NewDownloadService(client *http.Client, project Project) *DownloadService {
	return &DownloadService{
		client:   client,
		project:  project,
		endpoint: endpointProjectDownload,
	}
}

func (s *DownloadService) Download(ver string, build int, jar string) {
	fmt.Println("Downloading...")
	fmt.Println("Done!")
}
