package flocksofblocks

// InputElement

type MultiSelectMenuWithConversationsList struct {
	slackType ElementType
	actionID  string

	confirm          ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      CompositionText

	// Conversation
	defaultToCurrentConversation bool
	initialConversations         []string
	filter                       Filter

	optionals multiSelectMenuWithConversationsListOptions
}

type multiSelectMenuWithConversationsListOptions struct {
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversations         bool
	Filter                       bool
}

// abstracted type
type multiSelectMenuWithConversationsListAbstraction struct {
	Type     string
	ActionId string

	Confirm          ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      CompositionText

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversations         []string
	Filter                       Filter

	Optionals multiSelectMenuWithConversationsListOptions
}

func NewMultiSelectMenuWithConversationsList(actionId string) MultiSelectMenuWithConversationsList {
	return MultiSelectMenuWithConversationsList{
		slackType: MultiSelectMenuWithConversationsListElement,
		actionID:  actionId,
		optionals: multiSelectMenuWithConversationsListOptions{
			Confirm:                      false,
			MaxSelectedItems:             false,
			FocusOnLoad:                  false,
			Placeholder:                  false,
			DefaultToCurrentConversation: false,
			InitialConversations:         false,
			Filter:                       false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenuWithConversationsList) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithConversationsList) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenuWithConversationsList) UpdateActionId(actionId string) MultiSelectMenuWithConversationsList {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithConversationsList) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithConversationsList) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithConversationsList) AddConfirmDialog(confirm ConfirmationDialog) MultiSelectMenuWithConversationsList {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *MultiSelectMenuWithConversationsList) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenuWithConversationsList) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithConversationsList) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithConversationsList) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithConversationsList {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithConversationsList) UnsetMaxSelectedItems() MultiSelectMenuWithConversationsList {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenuWithConversationsList) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithConversationsList) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithConversationsList) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithConversationsList {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithConversationsList) UnsetFocusOnLoad() MultiSelectMenuWithConversationsList {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenuWithConversationsList) setPlaceholder(placeholder string) {
	m.placeholder = NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithConversationsList) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithConversationsList) AddPlaceholder(placeholder string) MultiSelectMenuWithConversationsList {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithConversationsList) RemovePlaceholder() MultiSelectMenuWithConversationsList {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// defaultToCurrentConversation

// setDefaultToCurrentConversation public set default to current conversation
func (m *MultiSelectMenuWithConversationsList) setDefaultToCurrentConversation(defaultToCurrentConversation bool) {
	m.defaultToCurrentConversation = defaultToCurrentConversation
	m.optionals.DefaultToCurrentConversation = defaultToCurrentConversation

}

// unsetDefaultToCurrentConversation public remove default to current conversation
func (m *MultiSelectMenuWithConversationsList) unsetDefaultToCurrentConversation() {
	m.setDefaultToCurrentConversation(false)
}

// DefaultToCurrentConversation public set default to current conversation
func (m MultiSelectMenuWithConversationsList) DefaultToCurrentConversation() MultiSelectMenuWithConversationsList {
	m.setDefaultToCurrentConversation(true)
	return m
}

// UnsetDefaultToCurrentConversation public remove default to current conversation
func (m MultiSelectMenuWithConversationsList) UnsetDefaultToCurrentConversation() MultiSelectMenuWithConversationsList {
	m.unsetDefaultToCurrentConversation()
	return m
}

//////////////////////////////////////////////////
// initialConversations

// addInitialConversation private add initial conversation
func (m *MultiSelectMenuWithConversationsList) addInitialConversation(initialConversation string) {
	m.initialConversations = append(m.initialConversations, initialConversation)
	m.optionals.InitialConversations = true
}

// removeInitialConversations private remove initial conversations
func (m *MultiSelectMenuWithConversationsList) removeInitialConversations() {
	m.initialConversations = []string{}
	m.optionals.InitialConversations = false
}

// AddInitialConversation public add initial conversation
func (m MultiSelectMenuWithConversationsList) AddInitialConversation(initialConversation string) MultiSelectMenuWithConversationsList {
	m.addInitialConversation(initialConversation)
	return m
}

// ClearInitialConversations clear initial conversations
func (m MultiSelectMenuWithConversationsList) ClearInitialConversations() MultiSelectMenuWithConversationsList {
	m.removeInitialConversations()
	return m
}

//////////////////////////////////////////////////
// filter

func (m *MultiSelectMenuWithConversationsList) setFilter(filter Filter) {
	m.filter = filter
	m.optionals.Filter = true
}

func (m *MultiSelectMenuWithConversationsList) removeFilter() {
	m.optionals.Filter = false
}

// AddFilter public set filter
func (m MultiSelectMenuWithConversationsList) AddFilter(filter Filter) MultiSelectMenuWithConversationsList {
	m.setFilter(filter)
	return m
}

// RemoveFilter public remove filter
func (m MultiSelectMenuWithConversationsList) RemoveFilter() MultiSelectMenuWithConversationsList {
	m.removeFilter()
	return m
}

// create abstract
func (m MultiSelectMenuWithConversationsList) abstraction() multiSelectMenuWithConversationsListAbstraction {
	return multiSelectMenuWithConversationsListAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,

		// Conversation
		DefaultToCurrentConversation: m.defaultToCurrentConversation,
		InitialConversations:         RemoveDuplicateString(m.initialConversations),
		Filter:                       m.filter,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m multiSelectMenuWithConversationsListAbstraction) Template() string {
	return `{
	"action_id": "{{ .ActionId }}",

	"type": "{{ .Type }}"

{{if .Optionals.InitialConversations}},
	"initial_conversations": [{{range $index, $conversations := .InitialConversations}}{{if $index}},{{end}}"{{ $conversations}}"{{end}}]
{{end}}

{{if .Optionals.DefaultToCurrentConversation}},
	"default_to_current_conversation": {{ .DefaultToCurrentConversation }}
{{end}}

{{if .Optionals.Confirm }},
	"confirm": {{ .Confirm.Render }}
{{end}}

{{if .Optionals.MaxSelectedItems }},
	"max_selected_items": {{ .MaxSelectedItems }}
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

func (m MultiSelectMenuWithConversationsList) ElementRender() {}

func (m MultiSelectMenuWithConversationsList) Render() string {
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m MultiSelectMenuWithConversationsList) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}
