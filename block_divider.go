package flocksofblocks

type Divider struct {
	slackType BlockType // required
	blockId   string    // optional

	optionals dividerOptionals
}

func (d Divider) BlockRender() {}

func NewDividerBlock() Divider {
	divider := Divider{
		slackType: DividerBlock,
		optionals: dividerOptionals{
			BlockId: false,
		},
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
	return Render(d.abstraction())
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

type dividerAbstraction struct {
	Type      string
	BlockId   string
	Optionals dividerOptionals
}

// create an abstraction of the divider block
func (d Divider) abstraction() dividerAbstraction {
	return dividerAbstraction{
		Type:      d.Type(),
		BlockId:   d.blockId,
		Optionals: d.optionals,
	}
}

type dividerOptionals struct {
	BlockId bool
}

// SetBlockId sets the block id for the block.
func (d Divider) SetBlockId(blockId string) Divider {
	d.setBlockId(blockId)
	return d
}

// RemoveBlockId removes the block id from the block.
func (d Divider) RemoveBlockId() Divider {
	d.blockId = ""
	d.optionals.BlockId = false
	return d
}
