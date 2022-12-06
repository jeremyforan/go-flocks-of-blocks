package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
)

type NumberInput struct {
	slackType        ElementType
	actionID         string
	isDecimalAllowed bool

	initialValue int
	minValue     int
	maxValue     int

	dispatchActionConfig composition.DispatchActionConfig

	focusOnLoad bool
	placeholder composition.CompositionText

	optionals numberInputOptions
}

// optionals struct
type numberInputOptions struct {
	MinValue             bool
	MaxValue             bool
	InitialValue         bool
	DispatchActionConfig bool
	FocusOnLoad          bool
	Placeholder          bool
}

func NewNumberInput(actionId string) NumberInput {
	return NumberInput{
		slackType:        NumberInput,
		actionID:         actionId,
		isDecimalAllowed: false,

		optionals: numberInputOptions{
			MinValue:             false,
			MaxValue:             false,
			InitialValue:         false,
			DispatchActionConfig: false,
			FocusOnLoad:          false,
			Placeholder:          false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (n *NumberInput) setActionId(actionId string) {
	n.actionID = actionId
}

func (n *NumberInput) removeActionId() {
	n.actionID = ""
}

// UpdateActionId public update action id
func (n NumberInput) UpdateActionId(actionId string) NumberInput {
	n.setActionId(actionId)
	return n
}

//////////////////////////////////////////////////
// isDecimalAllowed

func (n *NumberInput) setIsDecimalAllowed(isDecimalAllowed bool) {
	n.isDecimalAllowed = isDecimalAllowed
}

func (n *NumberInput) removeIsDecimalAllowed() {
	n.isDecimalAllowed = false
}

// UpdateIsDecimalAllowed public update isDecimalAllowed
func (n NumberInput) DecimalAllowed() NumberInput {
	n.setIsDecimalAllowed(true)
	return n
}

func (n NumberInput) UnsetDecimalAllowed() NumberInput {
	n.setIsDecimalAllowed(false)
	return n
}

//////////////////////////////////////////////////
// initialValue

func (n *NumberInput) setInitialValue(initialValue int) {
	n.initialValue = initialValue
	n.optionals.InitialValue = true
}

func (n *NumberInput) removeInitialValue() {
	n.optionals.InitialValue = false
}

// InitialValue public update initialValue
func (n NumberInput) InitialValue(initialValue int) NumberInput {
	n.setInitialValue(initialValue)
	return n
}

//////////////////////////////////////////////////
// minValue

func (n *NumberInput) setMinValue(minValue int) {
	n.minValue = minValue
	n.optionals.MinValue = true
}

func (n *NumberInput) removeMinValue() {
	n.optionals.MinValue = false
}

// UpdateMinValue public update minValue
func (n NumberInput) MinValue(minValue int) NumberInput {
	n.setMinValue(minValue)
	return n
}

//////////////////////////////////////////////////
// maxValue

func (n *NumberInput) setMaxValue(maxValue int) {
	n.maxValue = maxValue
	n.optionals.MaxValue = true
}

func (n *NumberInput) removeMaxValue() {
	n.optionals.MaxValue = false
}

// UpdateMaxValue public update maxValue
func (n NumberInput) MaxValue(maxValue int) NumberInput {
	n.setMaxValue(maxValue)
	return n
}

//////////////////////////////////////////////////
// dispatchActionConfig

func (n *NumberInput) setDispatchActionConfig(dispatchActionConfig composition.DispatchActionConfig) {
	n.dispatchActionConfig = dispatchActionConfig
	n.optionals.DispatchActionConfig = true
}

func (n *NumberInput) removeDispatchActionConfig() {
	n.optionals.DispatchActionConfig = false
}

// DispatchAction public update dispatchActionConfig
func (n NumberInput) DispatchAction(dispatchActionConfig composition.DispatchActionConfig) NumberInput {
	n.setDispatchActionConfig(dispatchActionConfig)
	return n
}

//////////////////////////////////////////////////
// focusOnLoad

func (n *NumberInput) setFocusOnLoad(focusOnLoad bool) {
	n.focusOnLoad = focusOnLoad
	n.optionals.FocusOnLoad = true
}

func (n *NumberInput) removeFocusOnLoad() {
	n.optionals.FocusOnLoad = false
}

// FocusOnLoad public update focusOnLoad
func (n NumberInput) FocusOnLoad() NumberInput {
	n.setFocusOnLoad(true)
	return n
}

func (n NumberInput) UnsetFocusOnLoad() NumberInput {
	n.setFocusOnLoad(false)
	return n
}

//////////////////////////////////////////////////
// placeholder

func (n *NumberInput) setPlaceholder(placeholder string) {
	n.placeholder = composition.NewPlainText(placeholder)
	n.optionals.Placeholder = true
}

func (n *NumberInput) removePlaceholder() {
	n.optionals.Placeholder = false
}

// Placeholder public update placeholder
func (n NumberInput) Placeholder(placeholder string) NumberInput {
	n.setPlaceholder(placeholder)
	return n
}

//////////////////////////////////////////////////
// abstraction

type numberInputAbstraction struct {
	Type                 string
	ActionID             string
	IsDecimalAllowed     bool
	InitialValue         int
	MinValue             int
	MaxValue             int
	DispatchActionConfig composition.DispatchActionConfig
	FocusOnLoad          bool
	Placeholder          composition.CompositionText

	Optionals numberInputOptions
}

func (n NumberInput) abstraction() numberInputAbstraction {
	return numberInputAbstraction{
		Type:                 n.slackType.String(),
		ActionID:             n.actionID,
		IsDecimalAllowed:     n.isDecimalAllowed,
		InitialValue:         n.initialValue,
		MinValue:             n.minValue,
		MaxValue:             n.maxValue,
		DispatchActionConfig: n.dispatchActionConfig,
		FocusOnLoad:          n.focusOnLoad,
		Placeholder:          n.placeholder,

		Optionals: n.optionals,
	}
}

//////////////////////////////////////////////////
// template

func (n numberInputAbstraction) Template() string {
	return `{
"type": "{{.Type}}",
"action_id": "{{.ActionID}}",
"is_decimal_allowed": {{.IsDecimalAllowed}}

{{if .Optionals.InitialValue}},
	"initial_value": "{{.InitialValue}}"
{{end}}

{{if .Optionals.MinValue}},
	"min_value": "{{.MinValue}}"
{{end}}

{{if .Optionals.MaxValue}},
	"max_value": "{{.MaxValue}}"
{{end}}

{{if .Optionals.DispatchActionConfig}},
	"dispatch_action_config": {{.DispatchActionConfig.Render}}
{{end}}

{{if .Optionals.FocusOnLoad}},
	"focus_on_load": {{.FocusOnLoad}}
{{end}}

{{if .Optionals.Placeholder}},
	"placeholder": {{.Placeholder.Render}}
{{end}}
}`
}

// render public render
func (n NumberInput) Render() string {
	raw := common.Render(n.abstraction())
	return common.Pretty(raw)
}

// element interface
func (n NumberInput) InputElement() {}

// Input
func (n NumberInput) Input(label string) block.Input {
	return block.NewInput(label, n)
}
