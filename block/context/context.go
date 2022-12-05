package context

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

type Context struct {
	slackType block.BlockType
	elements  []element.Element

	blockId string

	optionals contextOptions
}

type contextOptions struct {
	BlockId bool
}

// NewContext creates a new context.
func NewContext() Context {
	return Context{
		slackType: block.Context,
		elements:  []element.Element{},
		optionals: contextOptions{
			BlockId: false,
		},
	}
}

// setBlockId sets the block id for the block.
func (c *Context) setBlockId(blockId string) {
	c.blockId = blockId
	c.optionals.BlockId = true
}

// removeBlockId removes the block id from the context.
func (c *Context) removeBlockId() {
	c.optionals.BlockId = false
}

// AddBlockId chain function to add block id to an existing context
func (c Context) AddBlockId(blockId string) Context {
	c.setBlockId(blockId)
	return c
}

// RemoveBlockId remove add block id from context
func (c Context) RemoveBlockId() Context {
	c.removeBlockId()
	return c
}

// addElement adds an element to the context.
func (c *Context) addElement(element element.Element) {
	c.elements = append(c.elements, element)
}

// AddElement chain function to add an element to an existing context
func (c Context) AddElement(element element.Element) Context {
	c.addElement(element)
	return c
}

// ContextAbstraction is the abstraction of the context block.
type ContextAbstraction struct {
	Type     string
	BlockId  string
	Elements []element.Element
}

// BlockRender is the implementation of the BlockRender interface.
func (c Context) BlockRender() {}

// BlockRenderAbstraction is the implementation of the BlockRenderAbstraction interface.
func (c Context) abstraction() ContextAbstraction {
	return ContextAbstraction{
		Type:     c.slackType.String(),
		BlockId:  c.blockId,
		Elements: c.elements,
	}
}

// Template is the template for the context block.
func (c ContextAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"elements": [{{range $index, $element := .Elements}}{{$element.Render}}{{if not $last}},{{end}}{{end}}]

{{if .Optionals.BlockId}},
	"block_id": "{{.BlockId}}"
{{end}}
	}`
}

// BlockRenderAbstraction is the implementation of the BlockRenderAbstraction interface.
