package flocksofblocks

type RadioButton struct {
	slackType ElementType
	actionID  string

	options       []Option
	initialOption Option
	ConfirmationDialog
	focusOnLoad bool

	optionals radioButtonOptions
}

type radioButtonOptions struct {
	InitialOption bool
	Confirm       bool
	FocusOnLoad   bool
}

// NewRadioButton public constructor
func NewRadioButton(actionId string) RadioButton {
	return RadioButton{
		slackType: RadioButtonElement,
		actionID:  actionId,
		options:   []Option{},
		optionals: radioButtonOptions{
			InitialOption: false,
			Confirm:       false,
			FocusOnLoad:   false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *RadioButton) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *RadioButton) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m RadioButton) UpdateActionId(actionId string) RadioButton {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// options

// AddOption adds an option to the RadioButton
func (m *RadioButton) addOption(option Option) {
	m.options = append(m.options, option)
}

// AddOptions adds multiple options to the RadioButton
func (m *RadioButton) addOptions(options []Option) {
	for _, option := range options {
		m.addOption(option)
	}
}

// RemoveOption removes an option from the RadioButton
func (m *RadioButton) removeOption(option Option) {
	for i, v := range m.options {
		if v == option {
			m.options = append(m.options[:i], m.options[i+1:]...)
		}
	}
}

// RemoveOptions removes multiple options from the RadioButton
func (m *RadioButton) removeOptions(options []Option) {
	for _, option := range options {
		m.removeOption(option)
	}
}

// AddOption public update options
func (m RadioButton) AddOption(option Option) RadioButton {
	m.addOption(option)
	return m
}

// AddOptions public update options
func (m RadioButton) AddOptions(options []Option) RadioButton {
	m.addOptions(options)
	return m
}

// RemoveOption public update options
func (m RadioButton) RemoveOption(option Option) RadioButton {
	m.removeOption(option)
	return m
}

// RemoveOptions public update options
func (m RadioButton) RemoveOptions(options []Option) RadioButton {
	m.removeOptions(options)
	return m
}

//////////////////////////////////////////////////
// initialOption

func (m *RadioButton) setInitialOption(option Option) {
	m.initialOption = option
	m.optionals.InitialOption = true
}

func (m *RadioButton) removeInitialOption() {
	m.initialOption = Option{}
	m.optionals.InitialOption = false
}

// UpdateInitialOption public update initialOption
func (m RadioButton) UpdateInitialOption(option Option) RadioButton {
	m.setInitialOption(option)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *RadioButton) setConfirm(confirm ConfirmationDialog) {
	m.ConfirmationDialog = confirm
	m.optionals.Confirm = true
}

func (m *RadioButton) removeConfirm() {
	m.ConfirmationDialog = ConfirmationDialog{}
	m.optionals.Confirm = false
}

// UpdateConfirm public update confirm
func (m RadioButton) AddConfirmationDialog(confirm ConfirmationDialog) RadioButton {
	m.setConfirm(confirm)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *RadioButton) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

func (m *RadioButton) removeFocusOnLoad() {
	m.focusOnLoad = false
	m.optionals.FocusOnLoad = false
}

// UpdateFocusOnLoad public update focusOnLoad
func (m RadioButton) FocusOnLoad() RadioButton {
	m.setFocusOnLoad(true)
	return m
}

//////////////////////////////////////////////////
// abstraction

type radioButtonAbstraction struct {
	Type     string
	ActionID string

	Options       []Option
	InitialOption Option
	Confirm       ConfirmationDialog
	FocusOnLoad   bool

	Optionals radioButtonOptions
}

func (m RadioButton) abstraction() radioButtonAbstraction {
	return radioButtonAbstraction{
		Type:     m.slackType.String(),
		ActionID: m.actionID,

		Options:       m.options,
		InitialOption: m.initialOption,
		Confirm:       m.ConfirmationDialog,
		FocusOnLoad:   m.focusOnLoad,

		Optionals: m.optionals,
	}
}

// Render
func (m RadioButton) Render() string {
	output := Render(m.abstraction())
	return Pretty(output)
}

// template

func (m radioButtonAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"action_id": "{{.ActionID}}",
	"options": [{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]

{{if .Optionals.InitialOption}}
	"initial_option": {{.InitialOption.Render}}
{{end}}

{{if .Optionals.Confirm}}
	"confirm": {{.Confirm.Render}}
{{end}}

{{if .Optionals.FocusOnLoad}}
	"focus_on_load": {{.FocusOnLoad}}
{{end}}
}`
}
