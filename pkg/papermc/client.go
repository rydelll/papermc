package papermc

import (
	"net/http"
	"time"
)

const paperBaseURL string = "https://api.papermc.io/v2"

type service struct {
	client   *Client
	endpoint endpoint
}

type Client struct {
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration

	ProjectVersion  *projectVersionService
	ProjectBuild    *projectBuildService
	ProjectDownload *projectDownloadService
}

// Set an option on the client.
type Option func(*Client)

// Create a new PaperMC client.
//
// Options can be changed via Set methods passed in as paramaters.
func NewClient(opts ...Option) *Client {
	c := &Client{
		baseURL: paperBaseURL,
		timeout: time.Second * 30,
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

// SetEndpoint Set the base URL for the PaperMC API.
//
// An example and the default url is: https://api.papermc.io/v2.
// Change at your own risk, newer and older versions are no supported.
func SetBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// SetTimeout Set the request timeout for the PaperMC API.
func SetTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}
