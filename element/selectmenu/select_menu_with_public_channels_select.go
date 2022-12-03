package selectmenu

import (
	"go-flocks-of-blocks/block/section"
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/compositiontext"
	"go-flocks-of-blocks/composition/confirmationdialog"
	"go-flocks-of-blocks/element"
)

// InputElement

type MultiSelectMenuWithPublicChannelsSelect struct {
	slackType element.ElementType
	actionID  string

	confirm          confirmationdialog.ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      compositiontext.CompositionText

	// Public Channel
	initialChannels []string

	optionals multiSelectMenuWithPublicChannelsSelectOptions
}

type multiSelectMenuWithPublicChannelsSelectOptions struct {
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool

	// Public Channel
	InitialChannels bool
}

// abstracted type
type multiSelectMenuWithPublicChannelsSelectAbstraction struct {
	Type     string
	ActionId string

	Confirm          confirmationdialog.ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      compositiontext.CompositionText

	// Public Channel
	InitialChannels []string

	Optionals multiSelectMenuWithPublicChannelsSelectOptions
}

func NewMultiSelectMenuWithPublicChannelsSelect(actionId string) MultiSelectMenuWithPublicChannelsSelect {
	return MultiSelectMenuWithPublicChannelsSelect{
		slackType: element.MultiSelectMenuWithPublicChannelsSelect,
		actionID:  actionId,

		optionals: multiSelectMenuWithPublicChannelsSelectOptions{
			Confirm:          false,
			MaxSelectedItems: false,
			FocusOnLoad:      false,
			Placeholder:      false,
			InitialChannels:  false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenuWithPublicChannelsSelect) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenuWithPublicChannelsSelect) UpdateActionId(actionId string) MultiSelectMenuWithPublicChannelsSelect {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithPublicChannelsSelect) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithPublicChannelsSelect) AddConfirmDialog(confirm confirmationdialog.ConfirmationDialog) MultiSelectMenuWithPublicChannelsSelect {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *MultiSelectMenuWithPublicChannelsSelect) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenuWithPublicChannelsSelect) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithPublicChannelsSelect) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithPublicChannelsSelect {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithPublicChannelsSelect) UnsetMaxSelectedItems() MultiSelectMenuWithPublicChannelsSelect {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenuWithPublicChannelsSelect) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithPublicChannelsSelect) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithPublicChannelsSelect {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithPublicChannelsSelect) UnsetFocusOnLoad() MultiSelectMenuWithPublicChannelsSelect {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenuWithPublicChannelsSelect) setPlaceholder(placeholder string) {
	m.placeholder = compositiontext.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithPublicChannelsSelect) AddPlaceholder(placeholder string) MultiSelectMenuWithPublicChannelsSelect {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithPublicChannelsSelect) RemovePlaceholder() MultiSelectMenuWithPublicChannelsSelect {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// initialChannels

func (m *MultiSelectMenuWithPublicChannelsSelect) setInitialChannels(initialChannels []string) {
	m.initialChannels = initialChannels
	m.optionals.InitialChannels = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeInitialChannels() {
	m.initialChannels = []string{}
	m.optionals.InitialChannels = false
}

// addInitialChannel private set initial channels
func (m *MultiSelectMenuWithPublicChannelsSelect) addInitialChannel(initialChannel string) {
	m.initialChannels = append(m.initialChannels, initialChannel)
	m.optionals.InitialChannels = true
}

// removeInitialChannel private remove initial channels
func (m *MultiSelectMenuWithPublicChannelsSelect) removeInitialChannel(initialChannel string) {
	for i, v := range m.initialChannels {
		if v == initialChannel {
			m.initialChannels = append(m.initialChannels[:i], m.initialChannels[i+1:]...)
		}
	}
}

// AddInitialChannels public set initial channels
func (m MultiSelectMenuWithPublicChannelsSelect) AddInitialChannels(initialChannel string) MultiSelectMenuWithPublicChannelsSelect {
	m.addInitialChannel(initialChannel)
	return m
}

// RemoveInitialChannels public remove initial channels
func (m MultiSelectMenuWithPublicChannelsSelect) RemoveInitialChannels(initialChannel string) MultiSelectMenuWithPublicChannelsSelect {
	m.removeInitialChannel(initialChannel)
	return m
}

// create abstract
func (m MultiSelectMenuWithPublicChannelsSelect) abstraction() multiSelectMenuWithPublicChannelsSelectAbstraction {
	return multiSelectMenuWithPublicChannelsSelectAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,

		// Public Channel
		InitialChannels: common.RemoveDuplicateString(m.initialChannels),

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m multiSelectMenuWithPublicChannelsSelectAbstraction) Template() string {
	return `{
	"action_id": "{{ .ActionId }}",
	"type": "{{ .Type }}"

{{if .Optionals.InitialChannels}},
	"initial_channels": [{{range $index, $channel := .InitialChannels}}{{if $index}},{{end}}"{{ $channel}}"{{end}}]
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

func (m MultiSelectMenuWithPublicChannelsSelect) ElementRender() {}

func (m MultiSelectMenuWithPublicChannelsSelect) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

func (m MultiSelectMenuWithPublicChannelsSelect) Section() section.Section {
	s := section.NewSection("newSection").AddAccessory(m)
	return s
}
