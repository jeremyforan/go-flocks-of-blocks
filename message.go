package flocksofblocks

import (
	"bytes"
	"encoding/json"
	"net/url"
)

const (
	messageBlockLimit     = 50
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

// AddBlock add a block to the message
func (m Message) AddBlock(block Block) Message {
	m.Blocks = append(m.Blocks, block)
	return m
}

// Generate Url for slack interactive building site
func (m Message) GenerateKitBuilderUrl() *url.URL {
	compact := bytes.NewBuffer([]byte{})

	err := json.Compact(compact, []byte(m.Render()))
	if err != nil {
		return nil
	}

	//convert bytes to urlencoded string
	encoded := url.QueryEscape(compact.String())

	url, err := url.Parse(slackKitBuilderApiUrl + "#" + encoded)
	if err != nil {
		return nil
	}
	return url
}
