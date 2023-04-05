package chatgpt

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// AuthSession will check if session is expired and return a new accessToken
func (c *Client) AuthSession() (*gjson.Result, error) {
	req, err := http.NewRequest(http.MethodGet, AUTH_SESSION_URI, nil)
	if err != nil {
		return nil, fmt.Errorf("new request failed: %v", err)
	}

	resp, err := c.doRequest(req)

	if err != nil {
		return nil, fmt.Errorf("do request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %v", err)
	}

	if c.opts.Debug {
		log.Printf("http response info: %s\n", body)
	}

	res := gjson.ParseBytes(body)

	return &res, nil
}
