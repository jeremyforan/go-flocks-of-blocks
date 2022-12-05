package urlinput

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/dispatchactionconfig"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
	"net/url"
)

type URLInput struct {
	slackType element.ElementType
	actionID  string

	initialValue         *url.URL
	dispatchActionConfig dispatchactionconfig.DispatchActionConfig
	focusOnLoad          bool
	placeholder          compositiontext.CompositionText

	optionals urlInputOptions
}

type urlInputOptions struct {
	InitialValue         bool
	DispatchActionConfig bool
	FocusOnLoad          bool
	Placeholder          bool
}

func NewURLInput(actionId string) URLInput {
	return URLInput{
		slackType: element.UrlInput,
		actionID:  actionId,
		optionals: urlInputOptions{
			InitialValue:         false,
			DispatchActionConfig: false,
			FocusOnLoad:          false,
			Placeholder:          false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *URLInput) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *URLInput) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m URLInput) UpdateActionId(actionId string) URLInput {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// initialValue

func (m *URLInput) setInitialValue(initialValue *url.URL) {
	m.initialValue = initialValue
	m.optionals.InitialValue = true
}

func (m *URLInput) removeInitialValue() {
	m.optionals.InitialValue = false
}

// UpdateInitialValue public update initial value
func (m URLInput) UpdateInitialValue(initialValue *url.URL) URLInput {
	m.setInitialValue(initialValue)
	return m
}

// RemoveInitialValue public remove initial value
func (m URLInput) RemoveInitialValue() URLInput {
	m.removeInitialValue()
	return m
}

//////////////////////////////////////////////////
// dispatchActionConfig

func (m *URLInput) setDispatchActionConfig(dispatchActionConfig dispatchactionconfig.DispatchActionConfig) {
	m.dispatchActionConfig = dispatchActionConfig
	m.optionals.DispatchActionConfig = true
}

func (m *URLInput) removeDispatchActionConfig() {
	m.optionals.DispatchActionConfig = false
}

// UpdateDispatchActionConfig public update dispatch action config
func (m URLInput) UpdateDispatchActionConfig(dispatchActionConfig dispatchactionconfig.DispatchActionConfig) URLInput {
	m.setDispatchActionConfig(dispatchActionConfig)
	return m
}

// RemoveDispatchActionConfig public remove dispatch action config
func (m URLInput) RemoveDispatchActionConfig() URLInput {
	m.removeDispatchActionConfig()
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *URLInput) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *URLInput) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// UpdateFocusOnLoad public update focus on load
func (m URLInput) UpdateFocusOnLoad(focusOnLoad bool) URLInput {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// RemoveFocusOnLoad public remove focus on load
func (m URLInput) RemoveFocusOnLoad() URLInput {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *URLInput) setPlaceholder(placeholder compositiontext.CompositionText) {
	m.placeholder = placeholder
	m.optionals.Placeholder = true
}

func (m *URLInput) removePlaceholder() {
	m.optionals.Placeholder = false
}

// UpdatePlaceholder public update placeholder
func (m URLInput) UpdatePlaceholder(placeholder string) URLInput {
	m.setPlaceholder(compositiontext.NewPlainText(placeholder))
	return m
}

// RemovePlaceholder public remove placeholder
func (m URLInput) RemovePlaceholder() URLInput {
	m.removePlaceholder()
	return m
}

//////////////////////////////////////////////////
// abstraction

type urlInputAbstraction struct {
	Type     string
	ActionID string

	InitialValue         string
	DispatchActionConfig dispatchactionconfig.DispatchActionConfig
	FocusOnLoad          bool
	Placeholder          compositiontext.CompositionText

	Optionals urlInputOptions
}

// abstract method
func (m URLInput) abstraction() urlInputAbstraction {
	url := ""
	if m.optionals.InitialValue {
		url = m.initialValue.String()
	}

	return urlInputAbstraction{
		Type:     m.slackType.String(),
		ActionID: m.actionID,

		InitialValue:         url,
		DispatchActionConfig: m.dispatchActionConfig,
		FocusOnLoad:          m.focusOnLoad,
		Placeholder:          m.placeholder,

		Optionals: m.optionals,
	}
}

//////////////////////////////////////////////////
// render

// Render method
func (m URLInput) Render() string {
	output := common.Render(m.abstraction())
	return common.Pretty(output)
}

//////////////////////////////////////////////////
// template

// Template method
func (m urlInputAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"action_id": "{{.ActionID}}"

{{if .Optionals.InitialValue}},
	"initial_value": "{{.InitialValue}}"
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
}
`
}
