package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
)

//Plain TextInput
// https://api.slack.com/reference/block-kit/block-elements#input

type PlainTextInput struct {
	slackType ElementType
	actionID  string

	initialValue         string
	multiline            bool
	minLength            int
	maxLength            int
	dispatchActionConfig composition.DispatchActionConfig
	focusOnLoad          bool
	placeholder          composition.CompositionText

	optionals plainTextInputOptions
}

// optionals
type plainTextInputOptions struct {
	InitialValue         bool
	Multiline            bool
	MinLength            bool
	MaxLength            bool
	DispatchActionConfig bool
	FocusOnLoad          bool
	Placeholder          bool
}

//////////////////////////////////////////////////

func NewPlainTextInput(actionId string) PlainTextInput {
	return PlainTextInput{
		slackType: PlainTextInputElement,
		actionID:  actionId,
		optionals: plainTextInputOptions{
			InitialValue:         false,
			Multiline:            false,
			MinLength:            false,
			MaxLength:            false,
			DispatchActionConfig: false,
			FocusOnLoad:          false,
			Placeholder:          false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *PlainTextInput) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *PlainTextInput) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m PlainTextInput) UpdateActionId(actionId string) PlainTextInput {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// initialValue

func (m *PlainTextInput) setInitialValue(initialValue string) {
	m.initialValue = initialValue
	m.optionals.InitialValue = true
}

func (m *PlainTextInput) removeInitialValue() {
	m.initialValue = ""
	m.optionals.InitialValue = false
}

// UpdateInitialValue public update initial value
func (m PlainTextInput) UpdateInitialValue(initialValue string) PlainTextInput {
	m.setInitialValue(initialValue)
	return m
}

//////////////////////////////////////////////////
// multiline

func (m *PlainTextInput) setMultiline(multiline bool) {
	m.multiline = multiline
	m.optionals.Multiline = true
}

func (m *PlainTextInput) removeMultiline() {
	m.multiline = false
	m.optionals.Multiline = false
}

// EnableMultiline public update multiline
func (m PlainTextInput) EnableMultiline() PlainTextInput {
	m.setMultiline(true)
	return m
}

func (m PlainTextInput) DisableMultiline() PlainTextInput {
	m.setMultiline(false)
	return m
}

//////////////////////////////////////////////////
// minLength

func (m *PlainTextInput) setMinLength(minLength int) {
	m.minLength = minLength
	m.optionals.MinLength = true
}

func (m *PlainTextInput) removeMinLength() {
	m.minLength = 0
	m.optionals.MinLength = false
}

// UpdateMinLength public update min length
func (m PlainTextInput) SetMinLength(minLength int) PlainTextInput {
	m.setMinLength(minLength)
	return m
}

//////////////////////////////////////////////////
// maxLength

func (m *PlainTextInput) setMaxLength(maxLength int) {
	m.maxLength = maxLength
	m.optionals.MaxLength = true
}

func (m *PlainTextInput) removeMaxLength() {
	m.maxLength = 0
	m.optionals.MaxLength = false
}

// UpdateMaxLength public update max length
func (m PlainTextInput) SetMaxLength(maxLength int) PlainTextInput {
	m.setMaxLength(maxLength)
	return m
}

//////////////////////////////////////////////////
// dispatchActionConfig

func (m *PlainTextInput) setDispatchActionConfig(dispatchActionConfig composition.DispatchActionConfig) {
	m.dispatchActionConfig = dispatchActionConfig
	m.optionals.DispatchActionConfig = true
}

func (m *PlainTextInput) removeDispatchActionConfig() {
	m.dispatchActionConfig = composition.DispatchActionConfig{}
	m.optionals.DispatchActionConfig = false
}

// AddDispatchActionConfig public update dispatch action config
func (m PlainTextInput) AddDispatchActionConfig(dispatchActionConfig composition.DispatchActionConfig) PlainTextInput {
	m.setDispatchActionConfig(dispatchActionConfig)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *PlainTextInput) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m PlainTextInput) FocusOnLoad() PlainTextInput {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m PlainTextInput) UnsetFocusOnLoad() PlainTextInput {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *PlainTextInput) setPlaceholder(placeholder string) {
	m.placeholder = composition.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *PlainTextInput) removePlaceholder() {
	m.optionals.Placeholder = false
}

// AddPlaceholder public set placeholder
func (m PlainTextInput) AddPlaceholder(placeholder string) PlainTextInput {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m PlainTextInput) RemovePlaceholder() PlainTextInput {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// abstraction

type plainTextInputAbstraction struct {
	Type     string
	ActionID string

	InitialValue         string
	Multiline            bool
	MinLength            int
	MaxLength            int
	DispatchActionConfig composition.DispatchActionConfig
	FocusOnLoad          bool
	Placeholder          composition.CompositionText

	Optionals plainTextInputOptions
}

// abstraction
func (m PlainTextInput) abstraction() plainTextInputAbstraction {
	return plainTextInputAbstraction{
		Type:     m.slackType.String(),
		ActionID: m.actionID,

		InitialValue:         m.initialValue,
		Multiline:            m.multiline,
		MinLength:            m.minLength,
		MaxLength:            m.maxLength,
		DispatchActionConfig: m.dispatchActionConfig,
		FocusOnLoad:          m.focusOnLoad,
		Placeholder:          m.placeholder,

		Optionals: m.optionals,
	}
}

// template

func (m plainTextInputAbstraction) Template() string {
	return `{
"type": "{{.Type}}",
"action_id": "{{.ActionID}}"

{{if .Optionals.InitialValue}},
	"initial_value": "{{.InitialValue}}"
{{end}}

{{if .Optionals.Multiline}},
	"multiline": {{.Multiline}}
{{end}}

{{if .Optionals.MinLength}},
	"min_length": {{.MinLength}}
{{end}}
	
{{if .Optionals.MaxLength}},
	"max_length": {{.MaxLength}}
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

// Render
func (m PlainTextInput) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}
