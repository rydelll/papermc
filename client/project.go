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

type Project string

type ProjectService struct {
	baseURL string
	client  *http.Client
	logger  *slog.Logger
	project Project
}

func NewProjectService(c *Client, project Project) *ProjectService {
	return &ProjectService{
		baseURL: c.baseURL,
		client:  c.client,
		logger:  c.logger,
		project: project,
	}
}
