package chatgpt

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	BASE_URI         = "https://chat.openai.com"
	AUTH_SESSION_URI = "https://chat.openai.com/api/auth/session"
	CONVERSATION_URI = "https://chat.openai.com/backend-api/conversation"
	GET_MODELS_URI   = "https://chat.openai.com/backend-api/models"
	USER_AGENT       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
	EOF_TEXT         = "[DONE]"
)

// MixMap is a type alias for map[string]interface{}
type MixMap = map[string]interface{}

// Client is a ChatGPT request client
type Client struct {
	opts    Options // custom options
	httpCli *http.Client
}

// NewClient will return a ChatGPT request client
func NewClient(options ...Option) *Client {
	cli := &Client{
		opts: Options{
			Timeout:   30 * time.Second,              // set default timeout
			UserAgent: USER_AGENT,                    // set default user-agent
			Model:     "text-davinci-002-render-sha", // set default chat model
		},
	}

	// load custom options
	for _, option := range options {
		option(cli)
	}

	cli.initHttpClient()

	return cli
}

func (c *Client) initHttpClient() {
	transport := &http.Transport{}

	if c.opts.Proxy != "" {
		proxy, err := url.Parse(c.opts.Proxy)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxy)
		}
	}

	c.httpCli = &http.Client{
		Timeout:   c.opts.Timeout,
		Transport: transport,
	}
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	if c.opts.UserAgent != "" {
		req.Header.Set("User-Agent", c.opts.UserAgent)
	}
	if c.opts.Cookie != "" {
		req.Header.Set("Cookie", c.opts.Cookie)
	} else if len(c.opts.Cookies) > 0 {
		for _, cookie := range c.opts.Cookies {
			req.AddCookie(cookie)
		}
	}

	if c.opts.Debug {
		reqInfo, _ := httputil.DumpRequest(req, true)
		log.Printf("http request info: \n%s\n", reqInfo)
	}

	resp, err := c.httpCli.Do(req)

	if c.opts.Debug {
		respInfo, _ := httputil.DumpResponse(resp, false)
		log.Printf("http response info: \n%s\n", respInfo)
	}

	return resp, err
}

// WithModel: set chat model
func (c *Client) WithModel(model string) *Client {
	c.opts.Model = model

	return c
}
