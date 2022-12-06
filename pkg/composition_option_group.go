package flocksofblocks

type OptionGroup struct {
	label   CompositionText
	options []Option
}

func NewOptionGroup(label string) OptionGroup {
	return OptionGroup{
		label:   NewPlainText(label),
		options: []Option{},
	}
}

// setLabel sets the label for the block.
func (o *OptionGroup) setLabel(label CompositionText) {
	o.label = label
}

// SetLabel sets the label for the block.
func (o OptionGroup) SetLabel(label string) OptionGroup {
	o.setLabel(NewPlainText(label))
	return o
}

func (o OptionGroup) AddOption(option Option) OptionGroup {
	o.options = append(o.options, option)
	return o
}

// compositionOptionAbstraction is used to render the composition option
type optionGroupAbstraction struct {
	Label   CompositionText
	Options []Option
}

// generate the abstraction for the block
func (o OptionGroup) abstraction() optionGroupAbstraction {
	return optionGroupAbstraction{
		Label:   o.label,
		Options: o.options,
	}
}

// Render renders the block to a string.
func (o OptionGroup) Render() string {
	return Render(o.abstraction())
}

// Template returns the template for the block.
func (o optionGroupAbstraction) Template() string {
	return `{
		"label": {{.Label.Render}},
		"options": [
			{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}
		]
	}`
}
