package client

import (
	"log/slog"
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
	logger  *slog.Logger
	project Project
}

// TODO
func NewProjectService(c *Client, project Project) *ProjectService {
	return &ProjectService{
		baseURL: c.baseURL,
		client:  c.client,
		logger:  c.logger,
		project: project,
	}
}
