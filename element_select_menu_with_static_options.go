package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
)

// InputElement

type SelectMenuWithStaticOption struct {
	slackType ElementType
	actionID  string
	options   []composition.Option

	optionGroups  []composition.OptionGroup
	initialOption composition.Option
	confirm       composition.ConfirmationDialog

	focusOnLoad bool
	placeholder composition.CompositionText

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
		slackType: SelectMenuWithStaticOptionsElement,
		actionID:  actionId,
		options:   []composition.Option{},
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

func (m *SelectMenuWithStaticOption) setOptions(options []composition.Option) {
	m.options = options
}

func (m *SelectMenuWithStaticOption) addOption(option composition.Option) {
	m.options = append(m.options, option)
}

func (m *SelectMenuWithStaticOption) removeOptions() {
	m.options = []composition.Option{}
}

// AddOption public add option
func (m SelectMenuWithStaticOption) AddOption(option composition.Option) SelectMenuWithStaticOption {
	m.addOption(option)
	return m
}

// ClearOptions clear options
func (m SelectMenuWithStaticOption) ClearOptions() SelectMenuWithStaticOption {
	m.removeOptions()
	return m
}

func (m *SelectMenuWithStaticOption) setOptionGroups(optionGroups []composition.OptionGroup) {
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
func (m SelectMenuWithStaticOption) AddOptionGroup(optionGroup composition.OptionGroup) SelectMenuWithStaticOption {
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

func (m *SelectMenuWithStaticOption) addInitialOption(initialOption composition.Option) {
	m.addOption(initialOption)
	m.initialOption = initialOption
	m.optionals.InitialOption = true
}

func (m *SelectMenuWithStaticOption) removeInitialOptions() {
	m.optionals.InitialOption = false
}

func (m *SelectMenuWithStaticOption) setInitialOptions(initialOption composition.Option) {
	m.initialOption = initialOption
	m.optionals.InitialOption = true
}

// ClearInitialOptions clear initial options
func (m SelectMenuWithStaticOption) ClearInitialOptions() SelectMenuWithStaticOption {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m SelectMenuWithStaticOption) AddInitialOption(initialOption composition.Option) SelectMenuWithStaticOption {
	m.addInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithStaticOption) setConfirm(confirm composition.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithStaticOption) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithStaticOption) AddConfirmDialog(confirm composition.ConfirmationDialog) SelectMenuWithStaticOption {
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
	m.placeholder = composition.NewPlainText(placeholder)
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
	Options          []composition.Option
	OptionGroups     []composition.OptionGroup
	InitialOption    composition.Option
	Confirm          composition.ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      composition.CompositionText

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
