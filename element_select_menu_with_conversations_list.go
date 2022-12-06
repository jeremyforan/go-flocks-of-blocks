package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
)

// InputElement

type SelectMenuWithConversationsList struct {
	slackType ElementType
	actionID  string

	confirm            composition.ConfirmationDialog
	responseUrlEnabled bool
	focusOnLoad        bool
	placeholder        composition.CompositionText

	// Conversation
	defaultToCurrentConversation bool
	initialConversation          string
	filter                       composition.Filter

	optionals selectMenuWithConversationsListOptions
}

type selectMenuWithConversationsListOptions struct {
	Confirm bool

	FocusOnLoad bool
	Placeholder bool

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversation          bool
	Filter                       bool
	ResponseUrlEnabled           bool
}

// abstracted type
type selectMenuWithConversationsListAbstraction struct {
	Type     string
	ActionId string

	Confirm composition.ConfirmationDialog

	FocusOnLoad bool
	Placeholder composition.CompositionText

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversation          string
	Filter                       composition.Filter
	ResponseUrlEnabled           bool

	Optionals selectMenuWithConversationsListOptions
}

func NewSelectMenuWithConversationsList(actionId string) SelectMenuWithConversationsList {
	return SelectMenuWithConversationsList{
		slackType: SelectMenuWithConversationsListElement,
		actionID:  actionId,
		optionals: selectMenuWithConversationsListOptions{
			Confirm:                      false,
			FocusOnLoad:                  false,
			Placeholder:                  false,
			DefaultToCurrentConversation: false,
			InitialConversation:          false,
			Filter:                       false,
			ResponseUrlEnabled:           false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *SelectMenuWithConversationsList) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *SelectMenuWithConversationsList) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m SelectMenuWithConversationsList) UpdateActionId(actionId string) SelectMenuWithConversationsList {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithConversationsList) setConfirm(confirm composition.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithConversationsList) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithConversationsList) AddConfirmDialog(confirm composition.ConfirmationDialog) SelectMenuWithConversationsList {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *SelectMenuWithConversationsList) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// responseUrlEnabled

func (m *SelectMenuWithConversationsList) setResponseUrlEnabled(responseUrlEnabled bool) {
	m.responseUrlEnabled = responseUrlEnabled
	m.optionals.ResponseUrlEnabled = responseUrlEnabled
}

func (m *SelectMenuWithConversationsList) unsetResponseUrlEnabled() {
	m.optionals.ResponseUrlEnabled = false
}

func (m SelectMenuWithConversationsList) EnableResponseUrlEnabled() SelectMenuWithConversationsList {
	m.setResponseUrlEnabled(true)
	return m
}

func (m SelectMenuWithConversationsList) DisableResponseUrlEnabled() SelectMenuWithConversationsList {
	m.setResponseUrlEnabled(false)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *SelectMenuWithConversationsList) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m SelectMenuWithConversationsList) FocusOnLoad() SelectMenuWithConversationsList {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m SelectMenuWithConversationsList) UnsetFocusOnLoad() SelectMenuWithConversationsList {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *SelectMenuWithConversationsList) setPlaceholder(placeholder string) {
	m.placeholder = composition.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *SelectMenuWithConversationsList) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m SelectMenuWithConversationsList) AddPlaceholder(placeholder string) SelectMenuWithConversationsList {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m SelectMenuWithConversationsList) RemovePlaceholder() SelectMenuWithConversationsList {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// defaultToCurrentConversation

// setDefaultToCurrentConversation public set default to current conversation
func (m *SelectMenuWithConversationsList) setDefaultToCurrentConversation(defaultToCurrentConversation bool) {
	m.defaultToCurrentConversation = defaultToCurrentConversation
	m.optionals.DefaultToCurrentConversation = defaultToCurrentConversation

}

// unsetDefaultToCurrentConversation public remove default to current conversation
func (m *SelectMenuWithConversationsList) unsetDefaultToCurrentConversation() {
	m.setDefaultToCurrentConversation(false)
}

// DefaultToCurrentConversation public set default to current conversation
func (m SelectMenuWithConversationsList) DefaultToCurrentConversation() SelectMenuWithConversationsList {
	m.setDefaultToCurrentConversation(true)
	return m
}

// UnsetDefaultToCurrentConversation public remove default to current conversation
func (m SelectMenuWithConversationsList) UnsetDefaultToCurrentConversation() SelectMenuWithConversationsList {
	m.unsetDefaultToCurrentConversation()
	return m
}

//////////////////////////////////////////////////
// initialConversations

// addInitialConversation private add initial conversation
func (m *SelectMenuWithConversationsList) addInitialConversation(initialConversation string) {
	m.initialConversation = initialConversation
	m.optionals.InitialConversation = true
}

// removeInitialConversations private remove initial conversations
func (m *SelectMenuWithConversationsList) removeInitialConversation() {
	m.optionals.InitialConversation = false
}

// SetInitialConversation public add initial conversation
func (m SelectMenuWithConversationsList) SetInitialConversation(initialConversation string) SelectMenuWithConversationsList {
	m.addInitialConversation(initialConversation)
	return m
}

// ClearInitialConversations clear initial conversations
func (m SelectMenuWithConversationsList) UnsetInitialConversation() SelectMenuWithConversationsList {
	m.removeInitialConversation()
	return m
}

//////////////////////////////////////////////////
// filter

func (m *SelectMenuWithConversationsList) setFilter(filter composition.Filter) {
	m.filter = filter
	m.optionals.Filter = true
}

func (m *SelectMenuWithConversationsList) removeFilter() {
	m.optionals.Filter = false
}

// AddFilter public set filter
func (m SelectMenuWithConversationsList) AddFilter(filter composition.Filter) SelectMenuWithConversationsList {
	m.setFilter(filter)
	return m
}

// RemoveFilter public remove filter
func (m SelectMenuWithConversationsList) RemoveFilter() SelectMenuWithConversationsList {
	m.removeFilter()
	return m
}

// create abstract
func (m SelectMenuWithConversationsList) abstraction() selectMenuWithConversationsListAbstraction {
	return selectMenuWithConversationsListAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm: m.confirm,

		FocusOnLoad: m.focusOnLoad,
		Placeholder: m.placeholder,

		// Conversation
		DefaultToCurrentConversation: m.defaultToCurrentConversation,
		InitialConversation:          m.initialConversation,
		Filter:                       m.filter,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m selectMenuWithConversationsListAbstraction) Template() string {
	return `{	
"type": "{{ .Type }}",
"action_id": "{{ .ActionId }}"

{{if .Optionals.InitialConversation}},
	"initial_conversation":  "{{ .InitialConversation }}"
{{end}}

{{if .Optionals.DefaultToCurrentConversation}},
	"default_to_current_conversation": {{ .DefaultToCurrentConversation }}
{{end}}

{{if .Optionals.Confirm }},
	"confirm": {{ .Confirm.Render }}
{{end}}

{{ if .Optionals.ResponseUrlEnabled }},
	"response_url_enabled": {{ .ResponseUrlEnabled }}
{{end}}

{{if .Optionals.Filter }},
	{{ .Filter.Render }}
{{end}}

{{if .Optionals.FocusOnLoad }},
	"focus_on_load": {{ .FocusOnLoad }}
{{end}}

{{if .Optionals.Placeholder }},
	"placeholder": {{ .Placeholder.Render }}
{{end}}

}`
}

func (m SelectMenuWithConversationsList) ElementRender() {}

func (m SelectMenuWithConversationsList) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

func (m SelectMenuWithConversationsList) Section() block.Section {
	s := block.NewSection("newSection").AddAccessory(m)
	return s
}
