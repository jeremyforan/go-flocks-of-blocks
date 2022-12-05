package optiongroup

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/option"
)

type OptionGroup struct {
	label   compositiontext.CompositionText
	options []option.Option
}

func NewOptionGroup(label string) OptionGroup {
	return OptionGroup{
		label:   compositiontext.NewPlainText(label),
		options: []option.Option{},
	}
}

// setLabel sets the label for the block.
func (o *OptionGroup) setLabel(label compositiontext.CompositionText) {
	o.label = label
}

// SetLabel sets the label for the block.
func (o OptionGroup) SetLabel(label string) OptionGroup {
	o.setLabel(compositiontext.NewPlainText(label))
	return o
}

func (o OptionGroup) AddOption(option option.Option) OptionGroup {
	o.options = append(o.options, option)
	return o
}

// compositionOptionAbstraction is used to render the composition option
type optionGroupAbstraction struct {
	Label   compositiontext.CompositionText
	Options []option.Option
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
	return common.Render(o.abstraction())
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
