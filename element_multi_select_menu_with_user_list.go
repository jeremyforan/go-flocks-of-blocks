package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
)

// InputElement

type MultiSelectMenuWithUserList struct {
	slackType ElementType
	actionID  string

	confirm          composition.ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      composition.CompositionText

	// User List
	initialUsers []string

	optionals multiSelectMenuWithUserListOptions
}

type multiSelectMenuWithUserListOptions struct {
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool

	// User List
	InitialUsers bool
}

func NewMultiSelectMenuWithUserList(actionId string) MultiSelectMenuWithUserList {
	return MultiSelectMenuWithUserList{
		slackType: MultiSelectMenuWithUserListElement,
		actionID:  actionId,
		optionals: multiSelectMenuWithUserListOptions{
			Confirm:          false,
			MaxSelectedItems: false,
			FocusOnLoad:      false,
			Placeholder:      false,
			InitialUsers:     false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenuWithUserList) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithUserList) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenuWithUserList) UpdateActionId(actionId string) MultiSelectMenuWithUserList {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithUserList) setConfirm(confirm composition.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithUserList) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithUserList) AddConfirmDialog(confirm composition.ConfirmationDialog) MultiSelectMenuWithUserList {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *MultiSelectMenuWithUserList) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenuWithUserList) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithUserList) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithUserList) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithUserList {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithUserList) UnsetMaxSelectedItems() MultiSelectMenuWithUserList {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenuWithUserList) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithUserList) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithUserList) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithUserList {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithUserList) UnsetFocusOnLoad() MultiSelectMenuWithUserList {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenuWithUserList) setPlaceholder(placeholder string) {
	m.placeholder = composition.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithUserList) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithUserList) AddPlaceholder(placeholder string) MultiSelectMenuWithUserList {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithUserList) RemovePlaceholder() MultiSelectMenuWithUserList {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// initialUsers

func (m *MultiSelectMenuWithUserList) addInitialUser(initialUser string) {
	m.initialUsers = append(m.initialUsers, initialUser)
	m.optionals.InitialUsers = true
}

func (m *MultiSelectMenuWithUserList) removeInitialUsers() {
	m.initialUsers = []string{}
	m.optionals.InitialUsers = false
}

// AddInitialUser public add initial user
func (m MultiSelectMenuWithUserList) AddInitialUser(initialUser string) MultiSelectMenuWithUserList {
	m.addInitialUser(initialUser)
	return m
}

// ClearInitialUsers clear initial users
func (m MultiSelectMenuWithUserList) ClearInitialUsers() MultiSelectMenuWithUserList {
	m.removeInitialUsers()
	return m
}

// ////////////////////////////////////////////////
// abstract

// abstracted type
type multiSelectMenuWithUserListAbstraction struct {
	Type     string
	ActionId string

	Confirm          composition.ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      composition.CompositionText

	// User List
	InitialUsers []string

	Optionals multiSelectMenuWithUserListOptions
}

// create abstract
func (m MultiSelectMenuWithUserList) abstraction() multiSelectMenuWithUserListAbstraction {
	return multiSelectMenuWithUserListAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,

		// User List
		InitialUsers: common.RemoveDuplicateString(m.initialUsers),

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m multiSelectMenuWithUserListAbstraction) Template() string {
	return `{
"action_id": "{{ .ActionId }}",
		
"type": "{{ .Type }}"	

{{if .Optionals.InitialUsers}},
	"initial_users": [{{range $index, $user := .InitialUsers}}{{if $index}},{{end}}"{{ $user}}"{{end}}]
{{end}}

{{if .Optionals.Confirm }},
	"confirm": {{ .Confirm.Render }}
{{end}}

{{if .Optionals.MaxSelectedItems }},
	"max_selected_items": {{ .MaxSelectedItems }}
{{end}}

{{if .Optionals.FocusOnLoad }},
	"focus_on_load": {{ .FocusOnLoad }}
{{end}}

{{if .Optionals.Placeholder }},
	"placeholder": {{ .Placeholder.Render }}
{{end}}

}`
}

func (m MultiSelectMenuWithUserList) ElementRender() {}

func (m MultiSelectMenuWithUserList) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

func (m MultiSelectMenuWithUserList) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}
