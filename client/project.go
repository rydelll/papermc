package client

import "net/http"

const (
	Folia     Project = "folia"
	Paper     Project = "paper"
	Velocity  Project = "velocity"
	Waterfall Project = "waterfall"
)

type Project string

type ProjectService struct {
	Version  *VersionService
	Build    *BuildService
	Download *DownloadService
}

func NewProjectService(client *http.Client, project Project) *ProjectService {
	return &ProjectService{
		Version:  NewVersionService(client, project),
		Build:    NewBuildService(client, project),
		Download: NewDownloadService(client, project),
	}
}
