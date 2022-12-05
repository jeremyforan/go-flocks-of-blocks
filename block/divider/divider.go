package divider

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
)

type Divider struct {
	slackType block.BlockType // required
	blockId   string          // optional

	optionals dividerOptionals
}

func (d Divider) BlockRender() {}

func NewDividerBlock(options ...buttonConstructionOptions) Divider {
	divider := Divider{
		slackType: block.Divider,
		optionals: dividerOptionals{
			BlockId: false,
		},
	}

	for _, option := range options {
		option(&divider)
	}

	return divider
}

// SetBlockId sets the block id for the block.
func (d *Divider) setBlockId(blockId string) {
	d.blockId = blockId
	d.optionals.BlockId = true
}

func (d *Divider) removeBlockId() {
	d.blockId = ""
	d.optionals.BlockId = false
}

// Render renders the block to a string.
func (d Divider) Render() string {
	return common.Render(d.abstraction())
}

// SlackType return slack type
func (d Divider) Type() string {
	return string(d.slackType)
}

func (d dividerAbstraction) Template() string {
	return `{
	"type": "{{.Type}}"
	{{if .Optionals.BlockId}},"block_id": "{{.BlockId}}"{{end}}
}`
}
