package divider

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
