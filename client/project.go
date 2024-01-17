package client

import (
	"net/http"
)

const (
	Folia     Project = "folia"
	Paper     Project = "paper"
	Velocity  Project = "velocity"
	Waterfall Project = "waterfall"
)

// TODO
type Project string

// TODO
type ProjectService struct {
	baseURL string
	client  *http.Client
	project Project
}

// TODO
func NewProjectService(c *Client, project Project) *ProjectService {
	return &ProjectService{
		baseURL: c.baseURL,
		client:  c.client,
		project: project,
	}
}
