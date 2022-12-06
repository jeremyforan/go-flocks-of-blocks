package selectmenu

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/confirmationdialog"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/option"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/optiongroup"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

// InputElement

type SelectMenuWithStaticOption struct {
	slackType element.ElementType
	actionID  string
	options   []option.Option

	optionGroups  []optiongroup.OptionGroup
	initialOption option.Option
	confirm       confirmationdialog.ConfirmationDialog

	focusOnLoad bool
	placeholder compositiontext.CompositionText

	optionals SelectMenuWithStaticOptionOptions
}

type SelectMenuWithStaticOptionOptions struct {
	OptionGroups  bool
	InitialOption bool
	Confirm       bool

	FocusOnLoad bool
	Placeholder bool
}

func (m SelectMenuWithStaticOption) emptyAllFalseOptions() SelectMenuWithStaticOptionOptions {
	return SelectMenuWithStaticOptionOptions{
		OptionGroups:  false,
		InitialOption: false,
		Confirm:       false,

		FocusOnLoad: false,
		Placeholder: false,
	}
}

func NewSelectMenuWithStaticOptions(actionId string) SelectMenuWithStaticOption {
	return SelectMenuWithStaticOption{
		slackType: element.SelectMenuWithStaticOptions,
		actionID:  actionId,
		options:   []option.Option{},
		optionals: SelectMenuWithStaticOptionOptions{
			OptionGroups:  false,
			InitialOption: false,
			Confirm:       false,

			FocusOnLoad: false,
			Placeholder: false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *SelectMenuWithStaticOption) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *SelectMenuWithStaticOption) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m SelectMenuWithStaticOption) UpdateActionId(actionId string) SelectMenuWithStaticOption {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// options

func (m *SelectMenuWithStaticOption) setOptions(options []option.Option) {
	m.options = options
}

func (m *SelectMenuWithStaticOption) addOption(option option.Option) {
	m.options = append(m.options, option)
}

func (m *SelectMenuWithStaticOption) removeOptions() {
	m.options = []option.Option{}
}

// AddOption public add option
func (m SelectMenuWithStaticOption) AddOption(option option.Option) SelectMenuWithStaticOption {
	m.addOption(option)
	return m
}

// ClearOptions clear options
func (m SelectMenuWithStaticOption) ClearOptions() SelectMenuWithStaticOption {
	m.removeOptions()
	return m
}

func (m *SelectMenuWithStaticOption) setOptionGroups(optionGroups []optiongroup.OptionGroup) {
	m.optionGroups = optionGroups
	m.optionals.OptionGroups = true
}

func (m *SelectMenuWithStaticOption) removeOptionGroups() {
	m.optionals.OptionGroups = false
}

// ClearOptionGroups clear option groups
func (m SelectMenuWithStaticOption) ClearOptionGroups() SelectMenuWithStaticOption {
	m.removeOptionGroups()
	return m
}

// AddOptionGroup public add option group
func (m SelectMenuWithStaticOption) AddOptionGroup(optionGroup optiongroup.OptionGroup) SelectMenuWithStaticOption {
	m.setOptionGroups(append(m.optionGroups, optionGroup))
	return m
}

//////////////////////////////////////////////////
// all options

// ClearAllOptions clear all options
func (m SelectMenuWithStaticOption) ClearAllOptions() SelectMenuWithStaticOption {
	m.removeOptions()
	m.removeInitialOptions()
	return m
}

//////////////////////////////////////////////////
// initialOptions

func (m *SelectMenuWithStaticOption) addInitialOption(initialOption option.Option) {
	m.addOption(initialOption)
	m.initialOption = initialOption
	m.optionals.InitialOption = true
}

func (m *SelectMenuWithStaticOption) removeInitialOptions() {
	m.optionals.InitialOption = false
}

func (m *SelectMenuWithStaticOption) setInitialOptions(initialOption option.Option) {
	m.initialOption = initialOption
	m.optionals.InitialOption = true
}

// ClearInitialOptions clear initial options
func (m SelectMenuWithStaticOption) ClearInitialOptions() SelectMenuWithStaticOption {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m SelectMenuWithStaticOption) AddInitialOption(initialOption option.Option) SelectMenuWithStaticOption {
	m.addInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithStaticOption) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithStaticOption) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithStaticOption) AddConfirmDialog(confirm confirmationdialog.ConfirmationDialog) SelectMenuWithStaticOption {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *SelectMenuWithStaticOption) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *SelectMenuWithStaticOption) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m SelectMenuWithStaticOption) FocusOnLoad() SelectMenuWithStaticOption {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m SelectMenuWithStaticOption) UnsetFocusOnLoad() SelectMenuWithStaticOption {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *SelectMenuWithStaticOption) setPlaceholder(placeholder string) {
	m.placeholder = compositiontext.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *SelectMenuWithStaticOption) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m SelectMenuWithStaticOption) SetPlaceholder(placeholder string) SelectMenuWithStaticOption {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m SelectMenuWithStaticOption) RemovePlaceholder() SelectMenuWithStaticOption {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// abstract

// abstracted type
type SelectMenuWithStaticOptionAbstraction struct {
	Type             string
	ActionId         string
	Options          []option.Option
	OptionGroups     []optiongroup.OptionGroup
	InitialOption    option.Option
	Confirm          confirmationdialog.ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      compositiontext.CompositionText

	Optionals SelectMenuWithStaticOptionOptions
}

func (m SelectMenuWithStaticOption) abstraction() SelectMenuWithStaticOptionAbstraction {
	return SelectMenuWithStaticOptionAbstraction{
		Type:          m.slackType.String(),
		ActionId:      m.actionID,
		Options:       m.options,
		OptionGroups:  m.optionGroups,
		InitialOption: m.initialOption,
		Confirm:       m.confirm,
		FocusOnLoad:   m.focusOnLoad,
		Placeholder:   m.placeholder,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m SelectMenuWithStaticOptionAbstraction) Template() string {
	return `
{
"action_id": "{{ .ActionId }}",
"type": "{{ .Type }}",	

{{if .Optionals.OptionGroups }}	
	"option_groups": [{{range $index, $option := .OptionGroups}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]
{{else}}
	"options": [{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]
{{end}}

{{if .Optionals.InitialOption}},
	"initial_option": {{ .InitialOption.Render }}
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

// Render returns json string
func (m SelectMenuWithStaticOption) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

// ElementRender
func (m SelectMenuWithStaticOption) ElementRender() {}

// SectionBlock public section block
func (m SelectMenuWithStaticOption) Section() block.Section {
	s := block.NewSection("newSection").AddAccessory(m)
	return s
}
