package chatgpt

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/tidwall/gjson"
)

// authSession will check if session is expired and return a new accessToken
func (c *Client) authSession() (*gjson.Result, error) {
	req, err := http.NewRequest(http.MethodGet, AUTH_SESSION_URI, nil)
	if err != nil {
		return nil, fmt.Errorf("new request failed: %v", err)
	}

	for _, cookie := range c.opts.Cookies {
		req.AddCookie(cookie)
	}
	req.Header.Set("User-Agent", c.opts.UserAgent)

	cli := &http.Client{
		Timeout: c.opts.Timeout,
	}

	if c.opts.Debug {
		reqInfo, _ := httputil.DumpRequest(req, true)
		log.Printf("http request info: \n%s\n", reqInfo)
	}

	resp, err := cli.Do(req)

	if c.opts.Debug {
		respInfo, _ := httputil.DumpResponse(resp, true)
		log.Printf("http response info: \n%s\n", respInfo)
	}

	if err != nil {
		return nil, fmt.Errorf("do request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %v", err)
	}

	res := gjson.ParseBytes(body)
	if res.String() == "" {
		return nil, fmt.Errorf("parse response body failed")
	}

	return &res, nil
}
