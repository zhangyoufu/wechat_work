package wxwork

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func noRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

var httpClient = &http.Client{
	CheckRedirect: noRedirect,
}

type Client struct {
	key string
}

func NewClient(key string) *Client {
	return &Client{
		key: key,
	}
}

func (c *Client) Send(msg *Message) (err error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return
	}
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + url.QueryEscape(c.key)
	_, err = httpClient.Post(url, "application/json", bytes.NewReader(body))
	return
}

func (c *Client) SendText(text string) (err error) {
	return c.Send(NewTextMessage(text))
}

func (c *Client) SendMarkdown(markdown string) (err error) {
	return c.Send(NewMarkdownMessage(markdown))
}

func (c *Client) SendImage(image []byte) (err error) {
	return c.Send(NewImageMessage(image))
}
