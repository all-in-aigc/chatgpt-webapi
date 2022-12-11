package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

// ChatText chat reply with text format
type ChatText struct {
	Content        string // text content
	ConversationID string // conversation context id
	MessageID      string // current message id, can used as next chat's parent_message_id
}

// ChatStream chat reply with sream
type ChatStream struct {
}

// GetChatText will return reply text
func (c *Client) GetChatText(message string, args ...string) (*ChatText, error) {
	res, err := c.getChatReply(message, args...)
	if err != nil {
		return nil, fmt.Errorf("get chat reply failed: %v", err)
	}

	content := res.Get("message.content.parts.0").String()
	conversationID := res.Get("conversation_id").String()
	messageID := res.Get("message.id").String()

	text := &ChatText{
		Content:        content,
		ConversationID: conversationID,
		MessageID:      messageID,
	}

	return text, nil
}

// getChatReply will return reply message with json format
func (c *Client) getChatReply(message string, args ...string) (*gjson.Result, error) {
	resp, err := c.sendMessage(message, args...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %v", err)
	}

	arr := strings.Split(string(body), "\n\n")

	l := len(arr)
	if l == 0 {
		return nil, fmt.Errorf("invalid reply message: %s", body)
	}

	str := arr[l-3]
	if str == "" || !strings.Contains(str, "data") || str == "data: [DONE]" {
		return nil, fmt.Errorf("invalid reply message: %s", body)
	}

	res := gjson.Parse(str[6:])

	return &res, nil
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

	bearerToken := fmt.Sprintf("Bearer %s", accessToken)

	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Authorization", bearerToken)
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Content-Type", "application/json")

	cli := &http.Client{
		Timeout: c.opts.Timeout,
	}

	resp, err := cli.Do(req)

	return resp, err
}
