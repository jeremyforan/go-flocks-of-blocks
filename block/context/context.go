package context

import (
	"go-flocks-of-blocks/block"
	"go-flocks-of-blocks/element"
)

type Context struct {
	slackType block.BlockType
	elements  []element.Element
	blockId   string

	optional contextOptions
}

type contextOptions struct {
	BlockId bool
}

// NewContext creates a new context.
func NewContext() Context {
	return Context{
		slackType: block.Context,
		elements:  []element.Element{},
		optional: contextOptions{
			BlockId: false,
		},
	}
}

// SetBlockId sets the block id for the block.
func (c *Context) setBlockId(blockId string) {
	c.blockId = blockId
	c.optional.BlockId = true
}

func (c *Context) removeBlockId() {
	c.blockId = ""
	c.optional.BlockId = false
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
