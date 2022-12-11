package chatgpt

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

// authSession will check if session is expired and return a new accessToken
func (c *Client) authSession() (*gjson.Result, error) {
	req, err := http.NewRequest(http.MethodGet, AUTH_SESSION_URI, nil)
	if err != nil {
		return nil, fmt.Errorf("new request failed: %v", err)
	}

	cookie := fmt.Sprintf("__Secure-next-auth.session-token=%s", c.opts.Token)

	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Cookie", cookie)

	cli := &http.Client{
		Timeout: c.opts.Timeout,
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %v", err)
	}

	res := gjson.ParseBytes(body)

	return &res, nil
}
