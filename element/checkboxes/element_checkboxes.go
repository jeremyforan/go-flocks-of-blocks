package checkboxes

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/confirmationdialog"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/option"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

type Checkboxes struct {
	slackType element.ElementType // required
	actionId  string              // required
	options   []option.Option     // required

	initialOptions     []option.Option // optional
	confirmationDialog confirmationdialog.ConfirmationDialog
	focusOnLoad        bool

	optional checkboxOptional
}

// checkboxOptional is a struct to keep track of which optional fields are set.
type checkboxOptional struct {
	InitialOptions     bool
	ConfirmationDialog bool
	FocusOnLoad        bool
}

// NewCheckboxes creates a new checkboxes element.
func NewCheckboxes(actionId string) Checkboxes {
	return Checkboxes{
		slackType: element.Checkboxes,
		actionId:  actionId,
		options:   []option.Option{},
		optional: checkboxOptional{
			InitialOptions:     false,
			ConfirmationDialog: false,
			FocusOnLoad:        false,
		},
	}
}

// AddOption adds an option to the checkboxes element.
func (c *Checkboxes) addOption(option option.Option) {
	c.options = append(c.options, option)
}

func (c *Checkboxes) addInitialOption(option option.Option) {
	c.initialOptions = append(c.initialOptions, option)
	c.optional.InitialOptions = true
}

// AddConfirmationDialog adds a confirmation dialog to the checkboxes element.
func (c *Checkboxes) addConfirmationDialog(confirmationDialog confirmationdialog.ConfirmationDialog) {
	c.confirmationDialog = confirmationDialog
	c.optional.ConfirmationDialog = true
}

// RemoveConfirmationDialog removes the confirmation dialog from the checkboxes element.
func (c *Checkboxes) removeConfirmationDialog() {
	c.optional.ConfirmationDialog = false
}

// SetFocusOnLoad sets the focus on load flag for the checkboxes element.
func (c *Checkboxes) setFocusOnLoad(focusOnLoad bool) {
	c.focusOnLoad = focusOnLoad
	c.optional.FocusOnLoad = true
}

// RemoveFocusOnLoad removes the focus on load flag from the checkboxes element.
func (c *Checkboxes) removeFocusOnLoad() {
	c.optional.FocusOnLoad = false
}

type abstractCheckboxes struct {
	Type               string
	ActionID           string
	Options            []option.Option
	InitialOptions     []option.Option
	ConfirmationDialog confirmationdialog.ConfirmationDialog
	FocusOnLoad        bool

	Optional checkboxOptional
}

// create a new abstract checkboxes element
func (c Checkboxes) abstraction() abstractCheckboxes {
	return abstractCheckboxes{
		Type:               c.slackType.String(),
		ActionID:           c.actionId,
		Options:            c.options,
		InitialOptions:     c.initialOptions,
		ConfirmationDialog: c.confirmationDialog,
		FocusOnLoad:        c.focusOnLoad,
		Optional:           c.optional,
	}
}

func (c abstractCheckboxes) Template() string {
	return `{
		"type": "{{.Type}}",
		"action_id": "{{.ActionID}}",
		"options": [
			{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}
		]{{if .Optional.InitialOptions}},
		"initial_options": [{{range $index, $option := .InitialOptions}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]{{end}}{{if .Optional.ConfirmationDialog}},
		"confirmation_dialog": {{.ConfirmationDialog.Render}}{{end}}{{if .Optional.FocusOnLoad}},
		"focus_on_load": "{{.FocusOnLoad}}"{{end}}
	}`
}

// Render renders the checkboxes element to JSON.
func (c Checkboxes) Render() string {
	return common.Render(c.abstraction())
}
