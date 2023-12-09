package client

import (
	"net/http"
	"time"
)

const paperBaseURL string = "https://api.papermc.io/v2"

const (
	folia     project = "folia"
	paper     project = "paper"
	velocity  project = "velocity"
	waterfall project = "waterfall"
)

type project string

type Client struct {
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration

	Folia     *projectService
	Paper     *projectService
	Velocity  *projectService
	Waterfall *projectService
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

	c.Folia = &projectService{client: c, project: folia}
	c.Paper = &projectService{client: c, project: paper}
	c.Velocity = &projectService{client: c, project: velocity}
	c.Waterfall = &projectService{client: c, project: waterfall}

	return c
}

// SetBaseURL Set the base URL for the PaperMC API.
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
