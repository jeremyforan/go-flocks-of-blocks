package divider

type dividerOptionals struct {
	BlockId bool
}

type buttonConstructionOptions func(*Divider)

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

func BlockId(blockId string) buttonConstructionOptions {
	return func(b *Divider) {
		b.setBlockId(blockId)
	}
}
