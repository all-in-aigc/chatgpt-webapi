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
	// Cookie will set in request headers with string format
	Cookie string
	// Proxy is used to proxy request
	Proxy string
	// AccessToken is used to authorization
	AccessToken string
	// Model is the chat model
	Model string
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

// WithCookie is used to set request cookies in header
func WithCookie(cookie string) Option {
	return func(c *Client) {
		c.opts.Cookie = cookie
	}
}

// WithProxy is used to set request proxy
func WithProxy(proxy string) Option {
	return func(c *Client) {
		c.opts.Proxy = proxy
	}
}

// WithAccessToken is used to set accessToken
func WithAccessToken(accessToken string) Option {
	return func(c *Client) {
		c.opts.AccessToken = accessToken
	}
}

// WithModel is used to set chat model
func WithModel(model string) Option {
	return func(c *Client) {
		c.opts.Model = model
	}
}
