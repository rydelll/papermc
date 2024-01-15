package client

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	defaultBaseURL string        = "https://api.papermc.io/v2"
	defaultTimeout time.Duration = time.Second * 30
)

type Client struct {
	baseURL string
	client  *http.Client
	logger  *slog.Logger
}

// New PaperMC client.
//
// Options can be changed via options methods passed in as paramaters.
func New(opts ...Option) *Client {
	c := &Client{
		baseURL: defaultBaseURL,
		client:  &http.Client{Timeout: defaultTimeout},
		logger:  slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
	}

	for _, opt := range opts {
		opt(c)
	}

	// Do stuff here related to the services on the client

	return c
}
