package chatgpt

import "time"

// Options can set custom options for ChatGPT request client
type Options struct {
	// Timeout is used to end http request after timeout duration
	Timeout time.Duration
	// Token is session token for each api
	Token string
}

// Option is used to set custom option
type Option func(*Client)

// WithToken is used to set token option
func WithToken(token string) Option {
	return func(c *Client) {
		c.opts.Token = token
	}
}

// WithTimeout is used to set request timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.opts.Timeout = timeout
	}
}
