package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/google/uuid"
	"github.com/launchdarkly/eventsource"
	"github.com/tidwall/gjson"
)

// ChatText chat reply with text format
type ChatText struct {
	ConversationID string // conversation context id
	MessageID      string // current message id, can used as next chat's parent_message_id
	Content        string // text content
}

// ChatStream chat reply with sream
type ChatStream struct {
	Stream chan *ChatText // chat text stream
	Err    error          // error message
}

// GetChatText will return text message
func (c *Client) GetChatText(message string, args ...string) (*ChatText, error) {
	resp, err := c.sendMessage(message, args...)
	if err != nil {
		return nil, fmt.Errorf("send message failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %v", err)
	}

	arr := strings.Split(string(body), "\n\n")

	const TEXT_ARR_MIN_LEN = 3
	const TEXT_STR_MIN_LEN = 6

	l := len(arr)
	if l < TEXT_ARR_MIN_LEN {
		return nil, fmt.Errorf("invalid reply message: %s", body)
	}

	str := arr[l-TEXT_ARR_MIN_LEN]
	if len(str) < TEXT_STR_MIN_LEN {
		return nil, fmt.Errorf("invalid reply message: %s", body)
	}

	text := str[TEXT_STR_MIN_LEN:]

	return c.parseChatText(text)
}

// GetChatStream will return text stream
func (c *Client) GetChatStream(message string, args ...string) (*ChatStream, error) {
	resp, err := c.sendMessage(message, args...)
	if err != nil {
		return nil, fmt.Errorf("send message failed: %v", err)
	}

	chatStream := &ChatStream{
		Stream: make(chan *ChatText),
		Err:    nil,
	}

	decoder := eventsource.NewDecoder(resp.Body)

	go func() {
		defer resp.Body.Close()
		defer close(chatStream.Stream)

		for {
			event, err := decoder.Decode()
			if err != nil {
				chatStream.Err = fmt.Errorf("decode data failed: %v", err)
				return
			}

			text := event.Data()
			if text == "" || text == EOF_TEXT {
				// read data finished, success return
				return
			}

			chatText, err := c.parseChatText(text)
			if err != nil {
				chatStream.Err = fmt.Errorf("parse chat text failed: %v", err)
				return
			}

			chatStream.Stream <- chatText
		}
	}()

	return chatStream, nil
}

// parseChatText will return a ChatText struct from ChatText json
func (c *Client) parseChatText(text string) (*ChatText, error) {
	if text == "" || text == EOF_TEXT {
		return nil, fmt.Errorf("invalid chat text: %s", text)
	}

	res := gjson.Parse(text)

	conversationID := res.Get("conversation_id").String()
	messageID := res.Get("message.id").String()
	content := res.Get("message.content.parts.0").String()

	if conversationID == "" || messageID == "" {
		return nil, fmt.Errorf("invalid chat text")
	}

	return &ChatText{
		ConversationID: conversationID,
		MessageID:      messageID,
		Content:        content,
	}, nil
}

// sendMessage will send message to ChatGPT server
func (c *Client) sendMessage(message string, args ...string) (*http.Response, error) {
	accessToken, err := c.getAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get accessToken failed: %v", err)
	}

	var messageID string
	var conversationID string
	var parentMessageID string

	messageID = uuid.NewString()
	if len(args) > 0 {
		conversationID = args[0]
	}
	if len(args) > 1 {
		parentMessageID = args[1]
	}
	if parentMessageID == "" {
		parentMessageID = uuid.NewString()
	}

	params := MixMap{
		"action":            "next",
		"model":             "text-davinci-002-render",
		"parent_message_id": parentMessageID,
		"messages": []MixMap{
			{
				"role": "user",
				"id":   messageID,
				"content": MixMap{
					"content_type": "text",
					"parts":        []string{message},
				},
			},
		},
	}

	if conversationID != "" {
		params["conversation_id"] = conversationID
	}

	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("marshal request body failed: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, CONVERSATION_URI, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("new request failed: %v", err)
	}

	for _, cookie := range c.opts.Cookies {
		req.AddCookie(cookie)
	}
	req.Header.Set("User-Agent", c.opts.UserAgent)

	bearerToken := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Set("Authorization", bearerToken)
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Content-Type", "application/json")

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

	return resp, err
}
