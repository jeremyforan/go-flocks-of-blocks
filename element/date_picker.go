package element

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
	"time"
)

//InputElement

type DatePicker struct {
	slackType ElementType
	actionId  string

	initialDate time.Time
	confirm     composition.ConfirmationDialog
	placeholder composition.CompositionText
	focus       bool

	options optionalDatePicker
}

type optionalDatePicker struct {
	InitialDate bool
	Confirm     bool
	Placeholder bool
	Focus       bool
}

func NewDatePicker(actionId string) DatePicker {
	return DatePicker{
		slackType: DatePickerElement,
		actionId:  actionId,
		options: optionalDatePicker{
			InitialDate: false,
			Confirm:     false,
			Placeholder: false,
			Focus:       false,
		},
	}
}

// SetInitialDate sets the initial date for the date picker.
func (d *DatePicker) setInitialDate(initialDate time.Time) {
	// todo: implement a parser for the date format YYYY-MM-DD
	d.initialDate = initialDate
	d.options.InitialDate = true
}

// removeInitialDate removes the initial date for the date picker.
func (d *DatePicker) removeInitialDate() {
	d.options.InitialDate = false
}

// SetConfirm sets the confirmation dialog for the date picker.
func (d *DatePicker) setConfirm(confirm composition.ConfirmationDialog) {
	d.confirm = confirm
	d.options.Confirm = true
}

// removeConfirm removes the confirmation dialog for the date picker.
func (d *DatePicker) removeConfirm() {
	d.options.Confirm = false
}

// SetPlaceholder sets the placeholder for the date picker.
func (d *DatePicker) setPlaceholder(placeholder string) {
	d.placeholder = composition.NewPlainText(placeholder)
	d.options.Placeholder = true
}

// removePlaceholder removes the placeholder for the date picker.
func (d *DatePicker) removePlaceholder() {
	d.options.Placeholder = false
}

// SetFocus sets the focus for the date picker.
func (d *DatePicker) setFocus(focus bool) {
	d.focus = focus
	d.options.Focus = true
}

// removeFocus removes the focus for the date picker.
func (d *DatePicker) removeFocus() {
	d.options.Focus = false
}

type abstractDatePicker struct {
	Type        string
	ActionId    string
	InitalDate  string
	Confirm     composition.ConfirmationDialog
	Placeholder composition.CompositionText
	Focus       bool
	Optionals   optionalDatePicker
}

// abstraction
func (d DatePicker) abstraction() abstractDatePicker {
	return abstractDatePicker{
		Type:     d.slackType.String(),
		ActionId: d.actionId,

		InitalDate:  d.initialDate.Format("2006-01-02"),
		Confirm:     d.confirm,
		Placeholder: d.placeholder,
		Focus:       d.focus,

		Optionals: d.options,
	}
}

// Template
func (d abstractDatePicker) Template() string {
	return `{
		"type": "{{.Type}}",
		"action_id": "{{.ActionId}}"{{if .Optionals.InitialDate}},
		"initial_date": "{{.InitalDate}}"{{end}}{{if .Optionals.Confirm}}, 
		"confirm": {{.Confirm.Render}}{{end}}{{if .Optionals.Placeholder}},
		"placeholder": {{.Placeholder.Render}}{{end}}{{if .Optionals.Focus}},
		"initial_focus": {{.Focus}}{{end}}
	}`
}

// Render
func (d DatePicker) Render() string {
	return common.Render(d.abstraction())
}

// AddInitialDate chain function to add initial date to an existing date picker
func (d DatePicker) AddInitialDate(initialDate time.Time) DatePicker {
	d.setInitialDate(initialDate)
	return d
}

// RemoveInitialDate remove add initial date from date picker
func (d DatePicker) RemoveInitialDate() DatePicker {
	d.removeInitialDate()
	return d
}

// AddConfirm chain function to add confirm to an existing date picker
func (d DatePicker) AddConfirm(confirm composition.ConfirmationDialog) DatePicker {
	d.setConfirm(confirm)
	return d
}

// RemoveConfirm remove add confirm from date picker
func (d DatePicker) RemoveConfirm() DatePicker {
	d.removeConfirm()
	return d
}

// AddPlaceholder chain function to add placeholder to an existing date picker
func (d DatePicker) AddPlaceholder(placeholder string) DatePicker {
	d.setPlaceholder(placeholder)
	return d
}

// RemovePlaceholder remove add placeholder from date picker
func (d DatePicker) RemovePlaceholder() DatePicker {
	d.removePlaceholder()
	return d
}

// MakeFocused chain function to add focus to an existing date picker
func (d DatePicker) MakeFocused() DatePicker {
	d.setFocus(true)
	return d
}

// RemoveInitialFocus remove add focus from date picker
func (d DatePicker) RemoveInitialFocus() DatePicker {
	d.setFocus(false)
	return d
}
