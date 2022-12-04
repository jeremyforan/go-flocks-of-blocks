package selectmenu

import (
	"go-flocks-of-blocks/block/section"
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/compositiontext"
	"go-flocks-of-blocks/composition/confirmationdialog"
	"go-flocks-of-blocks/element"
)

// InputElement

type SelectMenuWithPublicChannelsSelect struct {
	slackType element.ElementType
	actionID  string

	confirm            confirmationdialog.ConfirmationDialog
	responseUrlEnabled bool
	focusOnLoad        bool
	placeholder        compositiontext.CompositionText

	// Public Channel
	initialChannel string

	optionals SelectMenuWithPublicChannelsSelectOptions
}

type SelectMenuWithPublicChannelsSelectOptions struct {
	Confirm            bool
	ResponseUrlEnabled bool
	FocusOnLoad        bool
	Placeholder        bool

	// Public Channel
	InitialChannel bool
}

// abstracted type
type SelectMenuWithPublicChannelsSelectAbstraction struct {
	Type     string
	ActionId string

	Confirm            confirmationdialog.ConfirmationDialog
	ResponseUrlEnabled bool
	FocusOnLoad        bool
	Placeholder        compositiontext.CompositionText

	// Public Channel
	InitialChannel string

	Optionals SelectMenuWithPublicChannelsSelectOptions
}

func NewSelectMenuWithPublicChannelsSelect(actionId string) SelectMenuWithPublicChannelsSelect {
	return SelectMenuWithPublicChannelsSelect{
		slackType: element.SelectMenuWithPublicChannelsSelect,
		actionID:  actionId,

		optionals: SelectMenuWithPublicChannelsSelectOptions{
			Confirm:            false,
			ResponseUrlEnabled: false,
			FocusOnLoad:        false,
			Placeholder:        false,
			InitialChannel:     false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *SelectMenuWithPublicChannelsSelect) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *SelectMenuWithPublicChannelsSelect) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m SelectMenuWithPublicChannelsSelect) UpdateActionId(actionId string) SelectMenuWithPublicChannelsSelect {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithPublicChannelsSelect) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithPublicChannelsSelect) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithPublicChannelsSelect) AddConfirmDialog(confirm confirmationdialog.ConfirmationDialog) SelectMenuWithPublicChannelsSelect {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *SelectMenuWithPublicChannelsSelect) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// responseUrlEnabled

func (m *SelectMenuWithPublicChannelsSelect) setResponseUrlEnabled(responseUrlEnabled bool) {
	m.responseUrlEnabled = responseUrlEnabled
	m.optionals.ResponseUrlEnabled = responseUrlEnabled
}

func (m *SelectMenuWithPublicChannelsSelect) unsetResponseUrlEnabled() {
	m.optionals.ResponseUrlEnabled = false
}

func (m SelectMenuWithPublicChannelsSelect) EnableResponseUrlEnabled() SelectMenuWithPublicChannelsSelect {
	m.setResponseUrlEnabled(true)
	return m
}

func (m SelectMenuWithPublicChannelsSelect) DisableResponseUrlEnabled() SelectMenuWithPublicChannelsSelect {
	m.setResponseUrlEnabled(false)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *SelectMenuWithPublicChannelsSelect) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m SelectMenuWithPublicChannelsSelect) FocusOnLoad() SelectMenuWithPublicChannelsSelect {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m SelectMenuWithPublicChannelsSelect) UnsetFocusOnLoad() SelectMenuWithPublicChannelsSelect {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *SelectMenuWithPublicChannelsSelect) setPlaceholder(placeholder string) {
	m.placeholder = compositiontext.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *SelectMenuWithPublicChannelsSelect) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m SelectMenuWithPublicChannelsSelect) AddPlaceholder(placeholder string) SelectMenuWithPublicChannelsSelect {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m SelectMenuWithPublicChannelsSelect) RemovePlaceholder() SelectMenuWithPublicChannelsSelect {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// initialChannel

func (m *SelectMenuWithPublicChannelsSelect) setInitialChannel(initialChannel string) {
	m.initialChannel = initialChannel
	m.optionals.InitialChannel = true
}

func (m *SelectMenuWithPublicChannelsSelect) removeInitialChannel() {
	m.optionals.InitialChannel = false
}

// addInitialChannel private set initial channels
func (m *SelectMenuWithPublicChannelsSelect) addInitialChannel(initialChannel string) {
	m.initialChannel = initialChannel
	m.optionals.InitialChannel = true
}

// SetInitialChannel public set initial channels
func (m SelectMenuWithPublicChannelsSelect) SetInitialChannel(initialChannel string) SelectMenuWithPublicChannelsSelect {
	m.addInitialChannel(initialChannel)
	return m
}

// UnsetInitialChannel public remove initial channels
func (m SelectMenuWithPublicChannelsSelect) UnsetInitialChannel() SelectMenuWithPublicChannelsSelect {
	m.removeInitialChannel()
	return m
}

// create abstract
func (m SelectMenuWithPublicChannelsSelect) abstraction() SelectMenuWithPublicChannelsSelectAbstraction {
	return SelectMenuWithPublicChannelsSelectAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm: m.confirm,

		FocusOnLoad: m.focusOnLoad,
		Placeholder: m.placeholder,

		// Public Channel
		InitialChannel: m.initialChannel,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m SelectMenuWithPublicChannelsSelectAbstraction) Template() string {
	return `{
"type": "{{ .Type }}",
"action_id": "{{ .ActionId }}"

{{if .Optionals.InitialChannel}},
	"initial_channel": "{{ .InitialChannel }}"
{{end}}

{{if .Optionals.Confirm }},
	"confirm": {{ .Confirm.Render }}
{{end}}

{{ if .Optionals.ResponseUrlEnabled }},
	"response_url_enabled": {{ .ResponseUrlEnabled }}
{{end}}

{{if .Optionals.FocusOnLoad }},
	"focus_on_load": {{ .FocusOnLoad }}
{{end}}

{{if .Optionals.Placeholder }},
	"placeholder": {{ .Placeholder.Render }}
{{end}}

}`
}

func (m SelectMenuWithPublicChannelsSelect) ElementRender() {}

func (m SelectMenuWithPublicChannelsSelect) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

func (m SelectMenuWithPublicChannelsSelect) Section() section.Section {
	s := section.NewSection("newSection").AddAccessory(m)
	return s
}
