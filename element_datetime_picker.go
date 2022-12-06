package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
	"strconv"
	"time"
)

type DateTimePicker struct {
	slackType ElementType
	actionId  string

	initialDateTime time.Time
	confirm         composition.ConfirmationDialog
	focusOnLoad     bool

	options dateTimePickerOptions
}

type dateTimePickerOptions struct {
	InitialDateTime bool
	Confirm         bool
	FocusOnLoad     bool
}

// NewDateTimePicker creates a new date picker.
func NewDateTimePicker(actionId string) DateTimePicker {
	return DateTimePicker{
		slackType: DateTimePickerElement,
		actionId:  actionId,
		options: dateTimePickerOptions{
			InitialDateTime: false,
			Confirm:         false,
			FocusOnLoad:     false,
		},
	}
}

// SetInitialDateTime sets the initial date for the date picker.
func (d *DateTimePicker) setInitialDateTime(initialDateTime time.Time) {
	d.initialDateTime = initialDateTime
	d.options.InitialDateTime = true
}

func (d *DateTimePicker) removeInitialDateTime() {
	d.options.InitialDateTime = false
}

// SetConfirm sets the confirmation dialog for the date picker.
func (d *DateTimePicker) setConfirm(confirm composition.ConfirmationDialog) {
	d.confirm = confirm
	d.options.Confirm = true
}

func (d *DateTimePicker) removeConfirm() {
	d.options.Confirm = false
}

// SetFocusOnLoad sets the focus on load for the date picker.
func (d *DateTimePicker) setFocusOnLoad(focusOnLoad bool) {
	d.focusOnLoad = focusOnLoad
	d.options.FocusOnLoad = true
}

func (d *DateTimePicker) removeFocusOnLoad() {
	d.options.FocusOnLoad = false
}

// AddInitialDateTime chain function to add initial date to an existing date picker
func (d DateTimePicker) AddInitialDateTime(initialDateTime time.Time) DateTimePicker {
	d.setInitialDateTime(initialDateTime)
	return d
}

// RemoveInitialDateTime remove add initial date from date picker
func (d DateTimePicker) RemoveInitialDateTime() DateTimePicker {
	d.removeInitialDateTime()
	return d
}

// AddConfirm chain function to add confirmation dialog to an existing date picker
func (d DateTimePicker) AddConfirmationDialog(confirm composition.ConfirmationDialog) DateTimePicker {
	d.setConfirm(confirm)
	return d
}

// RemoveConfirm remove add confirmation dialog from date picker
func (d DateTimePicker) RemoveConfirmationDialog() DateTimePicker {
	d.removeConfirm()
	return d
}

// AddFocusOnLoad chain function to add focus on load to an existing date picker
func (d DateTimePicker) AddFocusOnLoad(focusOnLoad bool) DateTimePicker {
	d.setFocusOnLoad(focusOnLoad)
	return d
}

// RemoveFocusOnLoad remove add focus on load from date picker
func (d DateTimePicker) RemoveFocusOnLoad() DateTimePicker {
	d.removeFocusOnLoad()
	return d
}

// abstraction type
type abstractDateTimePicker struct {
	Type     string
	ActionId string

	InitialDateTime string
	Confirm         composition.ConfirmationDialog
	FocusOnLoad     bool

	Optionals dateTimePickerOptions
}

// abstraction method
func (d DateTimePicker) abstraction() abstractDateTimePicker {
	unixString := strconv.FormatInt(d.initialDateTime.Unix(), 10)

	return abstractDateTimePicker{
		Type:     d.slackType.String(),
		ActionId: d.actionId,

		InitialDateTime: unixString,
		Confirm:         d.confirm,
		FocusOnLoad:     d.focusOnLoad,

		Optionals: d.options,
	}
}

// Render renders the date picker to a JSON string.
func (d DateTimePicker) Render() string {
	return common.Render(d.abstraction())
}

// Template function
func (d abstractDateTimePicker) Template() string {
	return `{
		"type": "{{.Type}}",
		"action_id": "{{.ActionId}}"{{if .Optionals.InitialDateTime}},
		"initial_date": "{{.InitialDateTime}}"{{end}}{{if .Optionals.Confirm}},
		"confirm": {{.Confirm.Render}}{{end}}{{if .Optionals.FocusOnLoad}},
		"focus_on_load": {{.FocusOnLoad}},{{end}}
	}`
}
