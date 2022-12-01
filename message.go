package flocksofblocks

import (
	"go-flocks-of-blocks/block"
	"go-flocks-of-blocks/common"
)

const (
	messageBlockLimit = 50
)

type Message struct {
	Blocks []block.Block
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
	return common.Pretty(common.Render(m))
}

// AddBlock add a block to the message
func (m Message) AddBlock(block block.Block) Message {
	m.Blocks = append(m.Blocks, block)
	return m
}
