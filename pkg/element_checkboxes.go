package flocksofblocks

type Checkboxes struct {
	slackType ElementType // required
	actionId  string      // required
	options   []Option    // required

	initialOptions     []Option // optional
	confirmationDialog ConfirmationDialog
	focusOnLoad        bool

	optional checkboxOptional
}

// checkboxOptional is a struct to keep track of which optional fields are set.
type checkboxOptional struct {
	InitialOptions     bool
	ConfirmationDialog bool
	FocusOnLoad        bool
}

// NewCheckboxes creates a new checkbox element.
func NewCheckboxes(actionId string) Checkboxes {
	return Checkboxes{
		slackType: CheckboxesElement,
		actionId:  actionId,
		options:   []Option{},
		optional: checkboxOptional{
			InitialOptions:     false,
			ConfirmationDialog: false,
			FocusOnLoad:        false,
		},
	}
}

// AddOption adds an option to the checkboxes element.
func (c *Checkboxes) addOption(option Option) {
	c.options = append(c.options, option)
}

func (c *Checkboxes) addInitialOption(option Option) {
	c.initialOptions = append(c.initialOptions, option)
	c.optional.InitialOptions = true
}

// AddConfirmationDialog adds a confirmation dialog to the checkboxes element.
func (c *Checkboxes) addConfirmationDialog(confirmationDialog ConfirmationDialog) {
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
	Options            []Option
	InitialOptions     []Option
	ConfirmationDialog ConfirmationDialog
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
	return Render(c.abstraction())
}

// AddOption add option to checkboxes
func (c Checkboxes) AddOption(option Option) Checkboxes {
	c.addOption(option)
	return c
}

// AddInitialOption add initial option to checkboxes
func (c Checkboxes) AddInitialOption(option Option) Checkboxes {
	c.addInitialOption(option)
	c.addOption(option)
	return c
}

// AddConfirmationDialog add confirmation dialog to checkboxes
func (c Checkboxes) AddConfirmationDialog(confirmationDialog ConfirmationDialog) Checkboxes {
	c.addConfirmationDialog(confirmationDialog)
	return c
}

// RemoveConfirmationDialog remove confirmation dialog from checkboxes
func (c Checkboxes) RemoveConfirmationDialog() Checkboxes {
	c.removeConfirmationDialog()
	return c
}

// FocusOnLoad set focus on load to checkboxes
func (c Checkboxes) FocusOnLoad() Checkboxes {
	c.setFocusOnLoad(true)
	return c
}

// DisableFocusOnLoad remove focus on load from checkboxes
func (c Checkboxes) DisableFocusOnLoad() Checkboxes {
	c.setFocusOnLoad(false)
	return c
}