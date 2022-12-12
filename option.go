package chatgpt

import (
	"net/http"
	"time"
)

// Options can set custom options for ChatGPT request client
type Options struct {
	// Debug is used to output debug message
	Debug bool
	// Timeout is used to end http request after timeout duration
	Timeout time.Duration
	// UserAgent is used for custom user-agent
	UserAgent string
	// Cookies is request cookies for each api
	Cookies []*http.Cookie
}

// Option is used to set custom option
type Option func(*Client)

// WithDebug is used to output debug message
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.opts.Debug = debug
	}
}

// WithTimeout is used to set request timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.opts.Timeout = timeout
	}
}

// WithUserAgent is used to set request user-agent
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.opts.UserAgent = userAgent
	}
}

// WithCookies is used to set request cookies
func WithCookies(cookies []*http.Cookie) Option {
	return func(c *Client) {
		c.opts.Cookies = cookies
	}
}
