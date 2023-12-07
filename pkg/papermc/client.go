package papermc

import (
	"log/slog"
	"net/http"
	"time"
)

type endpoint string

type service struct {
	client   *Client
	endpoint endpoint
}

type Client struct {
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration
	logger     *slog.Logger

	ProjectVersion  *projectVersionService
	ProjectBuild    *projectBuildService
	ProjectDownload *projectDownloadService
}

// Set an option on the client.
type Option func(*Client)

// Create a new PaperMC client.
func NewClient(opts ...Option) *Client {
	c := &Client{
		baseURL: paperBaseURL,
		timeout: time.Second * 30,
		logger:  slog.Default(),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.httpClient = &http.Client{
		Timeout: c.timeout,
	}

	c.ProjectVersion = &projectVersionService{client: c, endpoint: endpointProjectVersion}
	c.ProjectBuild = &projectBuildService{client: c, endpoint: endpointProjectBuild}
	c.ProjectDownload = &projectDownloadService{client: c, endpoint: endpointProjectDownload}

	return c
}

// Set the endpoint for the PaperMC API.
func SetBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// Set the request timeout for the PaperMC API.
func SetTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// Set a custom structured logger for the PaperMC API.
func SetLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}
