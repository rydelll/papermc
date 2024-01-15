package client

import (
	"log/slog"
	"time"
)

type Option func(*Client)

// WithBaseURL modifies for base URL for the PaperMC API.
//
// For example the default URL is: https://api.papermc.io/v2.
// Change at your own risk, newer and older versions are not supported.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithLogger sets a custom logger for log messages.
func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithTimeout modifies request timeout for the PaperMC API.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}
