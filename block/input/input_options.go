package input

// 	dispatchAction bool
//	blockID        string
//	hint           compositiontext.CompositionText
//	slackOptional  bool

func (i Input) SetLabel(label string) Input {
	i.setLabel(label)
	return i
}

// EnableDispatchAction bool as chain method
func (i Input) EnableDispatchAction() Input {
	i.setDispatchAction(true)
	return i
}

// DisableDispatchAction bool as chain method
func (i Input) DisableDispatchAction() Input {
	i.setDispatchAction(false)
	return i
}

// SetBlockID BlockID string as chain method
func (i Input) SetBlockID(blockID string) Input {
	i.setBlockID(blockID)
	return i
}

// RemoveBlockID string as chain method
func (i Input) RemoveBlockID() Input {
	i.removeBlockID()
	return i
}

// SetHint as chain method
func (i Input) SetHint(s string) Input {
	i.setHint(s)
	return i
}

// RemoveHint as chain method
func (i Input) RemoveHint() Input {
	i.removeHint()
	return i
}

// SetSlackOptional as chain method
func (i Input) MakeOptional() Input {
	i.setSlackOptional()
	return i
}

func (i Input) RemoveOptional() Input {
	i.removeSlackOptional()
	return i
}
