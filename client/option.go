package client

import (
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

// WithTimeout modifies request timeout for the PaperMC API.
//
// The default request timeout is 30 seconds.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}
