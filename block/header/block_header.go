package header

import (
	"go-flocks-of-blocks/block"
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/compositiontext"
)

type Header struct {
	slackType block.BlockType
	text      compositiontext.CompositionText

	blockId string

	optional headerOptions
}

type headerOptions struct {
	BlockId bool
}

// NewHeader creates a new header.
func NewHeader(text string) Header {
	return Header{
		slackType: block.Header,
		text:      compositiontext.NewPlainText(text),
		optional: headerOptions{
			BlockId: false,
		},
	}
}

// SetBlockId sets the block id for the block.
func (h *Header) setBlockId(blockId string) {
	h.blockId = blockId
	h.optional.BlockId = true
}

func (h *Header) removeBlockId() {
	h.blockId = ""
	h.optional.BlockId = false
}

// AddBlockId chain function to add block id to an existing header
func (h Header) AddBlockId(blockId string) Header {
	h.setBlockId(blockId)
	return h
}

// RemoveBlockId remove add block id from header
func (h Header) RemoveBlockId() Header {
	h.removeBlockId()
	return h
}

// abstraction is a helper struct to generate the abstraction for the header.
type headerAbstraction struct {
	Type     string
	Text     compositiontext.CompositionText
	BlockId  string
	Optional headerOptions
}

// generate abstraction from header
func (h Header) abstraction() headerAbstraction {
	return headerAbstraction{
		Type:    string(h.slackType),
		Text:    h.text,
		BlockId: h.blockId,

		Optional: h.optional,
	}
}

// template is a helper struct to generate the template for the header.
func (f headerAbstraction) Template() string {
	return `{
		"type": "{{.Type}}",
		"text": {{.Text.Render}}{{if .Optional.BlockId}},
		"block_id": "{{.BlockId}}"
		{{end}}
	}`
}

// render is a helper function to generate the json for the file.
func (h Header) Render() string {
	return common.Render(h.abstraction())
}
