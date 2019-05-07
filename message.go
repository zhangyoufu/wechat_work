package wxwork

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

type Message struct {
	MsgType  string    `json:"msgtype"` // text/markdown/image/news
	Text     *Text     `json:"text,omitempty"`
	Markdown *Markdown `json:"markdown,omitempty"`
	Image    *Image    `json:"image,omitempty"`
	News     *News     `json:"news,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}

/*Supported syntax:
  # level 1 title
  ## level 2 title
  ### level 3 title
  #### level 4 title
  ##### level 5 title
  ###### level 6 title
  **bold**
  [link](http://example.com)
  `inline code`
  > reference
  <font color="info">green</font>
  <font color="comment">gray</font>
  <font color="warning">orange</font>
*/
type Markdown struct {
	Content string `json:"content"`
}

// JPG/PNG < 2MiB (before base64 encoding)
type Image struct {
	Base64 string `json:"base64"`
	MD5    string `json:"md5"` // before base64 encoding
}

type News struct {
	Articles []Article `json:"articles"` // 1~8 articles
}

type Article struct {
	Title       string `json:"title"`                 // truncated to 128 bytes
	Description string `json:"description,omitempty"` // truncated to 512 bytes
	URL         string `json:"url"`
	PicURL      string `json:"picurl,omitempty"` // JPG/PNG, large:1068×455px, small:150×150px
}

func NewTextMessage(text string) *Message {
	return &Message{
		MsgType: "text",
		Text: &Text{
			Content: text,
		},
	}
}

func NewMarkdownMessage(markdown string) *Message {
	return &Message{
		MsgType: "markdown",
		Markdown: &Markdown{
			Content: markdown,
		},
	}
}

func NewImageMessage(image []byte) *Message {
	dgst := md5.Sum(image)
	return &Message{
		MsgType: "image",
		Image: &Image{
			Base64: base64.RawStdEncoding.EncodeToString(image),
			MD5:    hex.EncodeToString(dgst[:]),
		},
	}
}
