package chatgpt

import (
	"time"
)

const (
	BASE_URI         = "https://chat.openai.com"
	AUTH_SESSION_URI = "https://chat.openai.com/api/auth/session"
	CONVERSATION_URI = "https://chat.openai.com/backend-api/conversation"
	USER_AGENT       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
	EOF_TEXT         = "[DONE]"
)

// MixMap is a type alias for map[string]interface{}
type MixMap = map[string]interface{}

// Client is a ChatGPT request client
type Client struct {
	opts Options // custom options
}

// NewClient will return a ChatGPT request client
func NewClient(options ...Option) *Client {
	cli := &Client{
		opts: Options{
			Timeout:   30 * time.Second, // set default timeout
			UserAgent: USER_AGENT,       // set default user-agent
		},
	}

	// load custom options
	for _, option := range options {
		option(cli)
	}

	return cli
}
