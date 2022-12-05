package input

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

type Input struct {
	slackType block.BlockType
	label     compositiontext.CompositionText
	element   element.InputElement

	dispatchAction bool
	blockID        string
	hint           compositiontext.CompositionText
	slackOptional  bool

	optionals inputOptional
}

type inputOptional struct {
	DispatchAction bool
	BlockID        bool
	Hint           bool
	SlackOptional  bool
}

func NewInput(label string, element element.InputElement) Input {
	return Input{
		slackType: block.Input,
		label:     compositiontext.NewPlainText(label),
		element:   element,
	}
}

// set dispatch action
func (i *Input) setDispatchAction(dispatchAction bool) {
	i.dispatchAction = dispatchAction
	i.optionals.DispatchAction = true
}

// set block id
func (i *Input) setBlockID(blockID string) {
	i.blockID = blockID
	i.optionals.BlockID = true
}

// remove block id
func (i *Input) removeBlockID() {
	i.optionals.BlockID = false
}

// setHint
func (i *Input) setHint(s string) {
	i.hint = compositiontext.NewPlainText(s)
	i.optionals.Hint = true
}

// removeHint remove hint
func (i *Input) removeHint() {
	i.optionals.Hint = false
}

// set label
func (i *Input) setLabel(label string) {
	i.label = compositiontext.NewPlainText(label)
}

// SetSlackOptional set slack optional
func (i *Input) setSlackOptional() {
	i.slackOptional = true
	i.optionals.SlackOptional = true
}

// RemoveSlackOptional remove slack optional
func (i *Input) removeSlackOptional() {
	i.optionals.SlackOptional = false
}

type abstractionInput struct {
	Type           string
	Label          compositiontext.CompositionText
	Element        element.InputElement
	DispatchAction bool
	BlockID        string
	Hint           compositiontext.CompositionText
	SlackOptional  bool
	Optionals      inputOptional
}

// create abstraction input
func (i Input) abstraction() abstractionInput {
	return abstractionInput{
		Type:           i.slackType.String(),
		Label:          i.label,
		Element:        i.element,
		DispatchAction: i.dispatchAction,
		BlockID:        i.blockID,
		Hint:           i.hint,
		SlackOptional:  i.slackOptional,
		Optionals:      i.optionals,
	}
}

// Template for input
func (i abstractionInput) Template() string {
	return `{
"type": "{{.Type}}",
"label": {{.Label.Render}},
	
"element": {{.Element.Render}}

{{if .Optionals.DispatchAction}},
		"dispatch_action": "{{.DispatchAction}}"
{{end}}

{{if .Optionals.BlockID}},
	"block_id": "{{.BlockID}}"
{{end}}

{{if .Optionals.Hint}},
	"hint": {{.Hint.Render}}
{{end}}

{{if .Optionals.SlackOptional}},
	"optional": "{{.SlackOptional}}"
{{end}}
	}`
}

// Render render input
func (i Input) Render() string {
	raw := common.Render(i.abstraction())
	return common.Pretty(raw)
}

type InputType interface {
	Input()
}
