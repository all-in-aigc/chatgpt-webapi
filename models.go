package chatgpt

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// GetModels to get all availabel model s
func (c *Client) GetModels() (*gjson.Result, []*http.Cookie, error) {
	req, err := http.NewRequest(http.MethodGet, GET_MODELS_URI, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("new request failed: %v", err)
	}

	accessToken, err := c.getAccessToken()
	if err != nil {
		return nil, nil, fmt.Errorf("get accessToken failed: %v", err)
	}

	bearerToken := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Set("Authorization", bearerToken)

	resp, err := c.doRequest(req)

	if err != nil {
		return nil, nil, fmt.Errorf("do request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("read response body failed: %v", err)
	}

	if c.opts.Debug {
		log.Printf("http response info: %s\n", body)
	}

	res := gjson.ParseBytes(body)

	return &res, resp.Cookies(), nil
}
