package flocksofblocks

type OverflowMenu struct {
	slackType ElementType
	actionID  string

	options []Option
	confirm ConfirmationDialog

	optionals overflowMenuOptions
}

type overflowMenuOptions struct {
	Confirm bool
}

// NewOverflowMenu creates a new OverflowMenu
func NewOverflowMenu(actionId string) OverflowMenu {
	return OverflowMenu{
		slackType: OverflowMenuElement,
		actionID:  actionId,
		options:   []Option{},
		optionals: overflowMenuOptions{
			Confirm: false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (o *OverflowMenu) setActionId(actionId string) {
	o.actionID = actionId
}

func (o *OverflowMenu) removeActionId() {
	o.actionID = ""
}

// UpdateActionId public update action id
func (o OverflowMenu) UpdateActionId(actionId string) OverflowMenu {
	o.setActionId(actionId)
	return o
}

//////////////////////////////////////////////////
// options

// AddOption adds an option to the OverflowMenu
func (o *OverflowMenu) addOption(option Option) {
	o.options = append(o.options, option)
}

// RemoveOption removes an option from the OverflowMenu
func (o *OverflowMenu) removeOption(option Option) {
	for i, v := range o.options {
		if v == option {
			o.options = append(o.options[:i], o.options[i+1:]...)
		}
	}
}

// AddOption public update options
func (o OverflowMenu) AddOption(options Option) OverflowMenu {
	o.addOption(options)
	return o
}

//////////////////////////////////////////////////
// confirm

func (o *OverflowMenu) setConfirm(confirm ConfirmationDialog) {
	o.confirm = confirm
	o.optionals.Confirm = true
}

func (o *OverflowMenu) removeConfirm() {
	o.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (o OverflowMenu) AddConfirmDialog(confirm ConfirmationDialog) OverflowMenu {
	o.setConfirm(confirm)
	o.optionals.Confirm = true
	return o
}

// RemoveConfirmDialog public remove confirm
func (o *OverflowMenu) RemoveConfirmDialog() {
	o.optionals.Confirm = false
}

//////////////////////////////////////////////////
// abstraction

type overflowMenuAbstraction struct {
	Type     string
	ActionID string
	Options  []Option
	Confirm  ConfirmationDialog

	Optionals overflowMenuOptions
}

// abstractOverflowMenu abstracts the OverflowMenu
func (o OverflowMenu) abstractOverflowMenu() overflowMenuAbstraction {
	return overflowMenuAbstraction{
		Type:     o.slackType.String(),
		ActionID: o.actionID,
		Options:  o.options,
		Confirm:  o.confirm,

		Optionals: o.optionals,
	}
}

// Template returns the template for the OverflowMenu
func (o overflowMenuAbstraction) Template() string {
	return `{
"type": "{{.Type}}",
"action_id": "{{.ActionID}}",
"options": [{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]

{{if .Optionals.Confirm}},
	"confirm": {{.Confirm.Render}}
{{end}}

}`
}

// Render
func (o OverflowMenu) Render() string {
	raw := Render(o.abstractOverflowMenu())
	return Pretty(raw)
}
