package papermc

import (
	"net/http"
	"time"
)

const (
	defaultBaseURL string        = "https://api.papermc.io/v2"
	defaultTimeout time.Duration = time.Second * 30
)

// A PaperMC client is used to fetch information and executables.
type Client struct {
	baseURL string
	client  *http.Client

	// Folia is a fork of the Paper Minecraft server that adds regionized
	// multithreading. Folia is designed to address the constant bottleneck of
	// the server running on a single thread causing performance issues. It is
	// not a drop-in replacement for Paper as it breaks most public plugins.
	Folia *ProjectService

	// Paper is a Minecraft game server based on Spigot, designed to greatly
	// improve performance and offer more advanced features and API.
	Paper *ProjectService

	// Travertine is a fork of Waterfall with additional protocols. Travertine
	// aims to support older client versions then Waterfall.
	Travertine *ProjectService

	// Velocity is the modern, high-performance proxy for Minecraft. Designed
	// with performance and stability in mind, it's a full alternative to
	// Waterfall with its own diverse plugin ecosystem.
	Velocity *ProjectService

	// Waterfall is an upgraded BungeeCord Minecraft proxy, offering full
	// compatibility with improvements to performance and stability.
	Waterfall *ProjectService
}

// New PaperMC client.
//
// Options can be changed via option methods passed in as paramaters.
func NewClient(opts ...Option) *Client {
	c := &Client{
		baseURL: defaultBaseURL,
		client:  &http.Client{Timeout: defaultTimeout},
	}

	// Apply optional configurations
	for _, opt := range opts {
		opt(c)
	}

	c.Folia = NewProjectService(Folia, c)
	c.Paper = NewProjectService(Paper, c)
	c.Travertine = NewProjectService(Travertine, c)
	c.Velocity = NewProjectService(Velocity, c)
	c.Waterfall = NewProjectService(Waterfall, c)

	return c
}
