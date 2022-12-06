package selectmenu

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/confirmationdialog"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

// InputElement

type SelectMenuWithUserList struct {
	slackType element.ElementType
	actionID  string

	confirm     confirmationdialog.ConfirmationDialog
	focusOnLoad bool
	placeholder compositiontext.CompositionText

	// User List
	initialUser string

	optionals multiSelectMenuWithUserListOptions
}

type multiSelectMenuWithUserListOptions struct {
	Confirm     bool
	FocusOnLoad bool
	Placeholder bool

	// User List
	InitialUser bool
}

func NewSelectMenuWithUserList(actionId string) SelectMenuWithUserList {
	return SelectMenuWithUserList{
		slackType: element.SelectMenuWithUserList,
		actionID:  actionId,
		optionals: multiSelectMenuWithUserListOptions{
			Confirm:     false,
			FocusOnLoad: false,
			Placeholder: false,
			InitialUser: false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *SelectMenuWithUserList) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *SelectMenuWithUserList) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m SelectMenuWithUserList) UpdateActionId(actionId string) SelectMenuWithUserList {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithUserList) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithUserList) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithUserList) AddConfirmDialog(confirm confirmationdialog.ConfirmationDialog) SelectMenuWithUserList {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *SelectMenuWithUserList) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *SelectMenuWithUserList) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m SelectMenuWithUserList) FocusOnLoad() SelectMenuWithUserList {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m SelectMenuWithUserList) UnsetFocusOnLoad() SelectMenuWithUserList {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *SelectMenuWithUserList) setPlaceholder(placeholder string) {
	m.placeholder = compositiontext.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *SelectMenuWithUserList) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m SelectMenuWithUserList) AddPlaceholder(placeholder string) SelectMenuWithUserList {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m SelectMenuWithUserList) RemovePlaceholder() SelectMenuWithUserList {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// initialUsers

func (m *SelectMenuWithUserList) setInitialUser(initialUser string) {
	m.initialUser = initialUser
	m.optionals.InitialUser = true
}

func (m *SelectMenuWithUserList) removeInitialUser() {
	m.optionals.InitialUser = false
}

// AddInitialUser public add initial user
func (m SelectMenuWithUserList) SetInitialUser(initialUser string) SelectMenuWithUserList {
	m.setInitialUser(initialUser)
	return m
}

// ClearInitialUsers clear initial users
func (m SelectMenuWithUserList) ClearInitialUsers() SelectMenuWithUserList {
	m.removeInitialUser()
	return m
}

// ////////////////////////////////////////////////
// abstract

// abstracted type
type selectMenuWithUserListAbstraction struct {
	Type     string
	ActionId string

	Confirm     confirmationdialog.ConfirmationDialog
	FocusOnLoad bool
	Placeholder compositiontext.CompositionText

	// User List
	InitialUser string

	Optionals multiSelectMenuWithUserListOptions
}

// create abstract
func (m SelectMenuWithUserList) abstraction() selectMenuWithUserListAbstraction {
	return selectMenuWithUserListAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm:     m.confirm,
		FocusOnLoad: m.focusOnLoad,
		Placeholder: m.placeholder,

		// User List
		InitialUser: m.initialUser,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m selectMenuWithUserListAbstraction) Template() string {
	return `{
"type": "{{ .Type }}",
"action_id": "{{ .ActionId }}"

{{if .Optionals.InitialUser}},
	"initial_user": "{{ .InitialUser}}"
{{end}}

{{if .Optionals.Confirm }},
	"confirm": {{ .Confirm.Render }}
{{end}}

{{if .Optionals.FocusOnLoad }},
	"focus_on_load": {{ .FocusOnLoad }}
{{end}}

{{if .Optionals.Placeholder }},
	"placeholder": {{ .Placeholder.Render }}
{{end}}

}`
}

func (m SelectMenuWithUserList) ElementRender() {}

func (m SelectMenuWithUserList) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

func (m SelectMenuWithUserList) Section() block.Section {
	s := block.NewSection("newSection").AddAccessory(m)
	return s
}
