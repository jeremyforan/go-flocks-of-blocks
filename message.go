package flocksofblocks

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	slackKitBuilderApiUrl = "https://app.slack.com/block-kit-builder/"
)

type Message struct {
	Blocks []Block
}

// NewMessage Create a new message
func NewMessage() Message {
	return Message{}
}

func (m Message) Template() string {
	return `{
	"blocks": [{{ range $index, $element := .Blocks}}{{if $index}},{{end}}{{$element.Render}}{{end}}]
	}`
}

// Render the message
func (m Message) Render() string {
	raw := Render(m)
	return Pretty(raw)
}

// AddBlock add a block to the message.
func (m Message) AddBlock(block Block) Message {
	m.Blocks = append(m.Blocks, block)
	return m
}

// Stringer function which also parses the message to a pretty json format.
func (m Message) String() string {
	return m.Render()
}

func minify(s string) string {
	var buff bytes.Buffer
	err := json.Compact(&buff, []byte(s))
	if err != nil {
		return ""
	}
	return buff.String()
}

// GenerateKitBuilderUrl generates a url to the slack kit builder with the message block encoded in the url.
// This is useful for testing and validation. The URL is https://app.slack.com/block-kit-builder/#<url-encoded-message>
func (m Message) GenerateKitBuilderUrl() *url.URL {
	compact := minify(m.Render())

	//convert bytes to urlencoded string
	encoded := url.PathEscape(compact)

	//build the url
	url, err := url.Parse(slackKitBuilderApiUrl + "#" + encoded)
	if err != nil {
		return nil
	}
	return url
}

// todo: add a means of copying an existing message/block/element/
