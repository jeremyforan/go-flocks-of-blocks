package timepicker

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/confirmationdialog"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

type TimePicker struct {
	slackType element.ElementType
	actionID  string

	initialTime string
	confirm     confirmationdialog.ConfirmationDialog
	focusOnLoad bool
	placeholder compositiontext.CompositionText

	optionals timePickerOptions
}

type timePickerOptions struct {
	InitialTime bool
	Confirm     bool
	FocusOnLoad bool
	Placeholder bool
}

// NewTimePicker public constructor
func NewTimePicker(actionId string) TimePicker {
	return TimePicker{
		slackType: element.TimePicker,
		actionID:  actionId,
		optionals: timePickerOptions{
			InitialTime: false,
			Confirm:     false,
			FocusOnLoad: false,
			Placeholder: false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *TimePicker) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *TimePicker) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m TimePicker) UpdateActionId(actionId string) TimePicker {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// initialTime

func (m *TimePicker) setInitialTime(initialTime string) {
	m.initialTime = initialTime
}

func (m *TimePicker) removeInitialTime() {
	m.initialTime = ""
}

// UpdateInitialTime public update initialTime
func (m TimePicker) UpdateInitialTime(initialTime string) TimePicker {
	m.setInitialTime(initialTime)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *TimePicker) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *TimePicker) removeConfirm() {
	m.optionals.Confirm = false
}

// UpdateConfirm public update confirm
func (m TimePicker) UpdateConfirm(confirm confirmationdialog.ConfirmationDialog) TimePicker {
	m.setConfirm(confirm)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *TimePicker) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *TimePicker) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// UpdateFocusOnLoad public update focusOnLoad
func (m TimePicker) UpdateFocusOnLoad(focusOnLoad bool) TimePicker {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *TimePicker) setPlaceholder(placeholder compositiontext.CompositionText) {
	m.placeholder = placeholder
	m.optionals.Placeholder = true
}

func (m *TimePicker) removePlaceholder() {
	m.optionals.Placeholder = false
}

// UpdatePlaceholder public update placeholder
func (m TimePicker) UpdatePlaceholder(placeholder string) TimePicker {
	m.setPlaceholder(compositiontext.NewPlainText(placeholder))
	return m
}

//////////////////////////////////////////////////
// abstract

type timePickerAbstract struct {
	Type        string
	ActionID    string
	InitialTime string
	Confirm     confirmationdialog.ConfirmationDialog
	FocusOnLoad bool
	Placeholder compositiontext.CompositionText

	Optionals timePickerOptions
}

// abstract public method
func (m TimePicker) abstraction() timePickerAbstract {
	return timePickerAbstract{
		Type:        m.slackType.String(),
		ActionID:    m.actionID,
		InitialTime: m.initialTime,
		Confirm:     m.confirm,
		FocusOnLoad: m.focusOnLoad,
		Placeholder: m.placeholder,

		Optionals: m.optionals,
	}
}

//////////////////////////////////////////////////
// template

// Template public method
func (m timePickerAbstract) Template() string {
	return `{
"type": "{{.Type}}",
"action_id": "{{.ActionID}}"

{{if .Optionals.InitialTime}},
	"initial_time": "{{.InitialTime}}"
{{end}}

{{if .Optionals.Confirm}},
	"confirm": {{.Confirm.Render}}
{{end}}
		
{{if .Optionals.FocusOnLoad}},
	"focus_on_load": {{.FocusOnLoad}}
{{end}}
		
{{if .Optionals.Placeholder}},
	"placeholder": {{.Placeholder.Render}}
{{end}}

	}`
}

//////////////////////////////////////////////////
// render

// Render public method
func (m TimePicker) Render() string {
	output := common.Render(m.abstraction())
	return common.Pretty(output)
}