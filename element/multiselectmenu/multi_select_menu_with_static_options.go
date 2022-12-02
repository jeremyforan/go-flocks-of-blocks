package multiselectmenu

import (
	"go-flocks-of-blocks/block/section"
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/compositiontext"
	"go-flocks-of-blocks/composition/confirmationdialog"
	"go-flocks-of-blocks/composition/option"
	"go-flocks-of-blocks/composition/optiongroup"
	"go-flocks-of-blocks/element"
)

// InputElement

type MultiSelectMenuWithStaticOption struct {
	slackType element.ElementType
	actionID  string
	options   []option.Option

	optionGroups     []optiongroup.OptionGroup
	initialOptions   []option.Option
	confirm          confirmationdialog.ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      compositiontext.CompositionText

	optionals multiSelectMenuWithStaticOptionOptions
}

type multiSelectMenuWithStaticOptionOptions struct {
	OptionGroups     bool
	InitialOptions   bool
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool
}

func (m MultiSelectMenuWithStaticOption) emptyAllFalseOptions() multiSelectMenuWithStaticOptionOptions {
	return multiSelectMenuWithStaticOptionOptions{
		OptionGroups:     false,
		InitialOptions:   false,
		Confirm:          false,
		MaxSelectedItems: false,
		FocusOnLoad:      false,
		Placeholder:      false,
	}
}

func NewMultiSelectMenuWithStaticOptions(actionId string) MultiSelectMenuWithStaticOption {
	return MultiSelectMenuWithStaticOption{
		slackType: element.MultiSelectMenuWithStaticOptions,
		actionID:  actionId,
		options:   []option.Option{},
		optionals: multiSelectMenuWithStaticOptionOptions{
			OptionGroups:     false,
			InitialOptions:   false,
			Confirm:          false,
			MaxSelectedItems: false,
			FocusOnLoad:      false,
			Placeholder:      false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenuWithStaticOption) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithStaticOption) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenuWithStaticOption) UpdateActionId(actionId string) MultiSelectMenuWithStaticOption {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// options

func (m *MultiSelectMenuWithStaticOption) setOptions(options []option.Option) {
	m.options = options
}

func (m *MultiSelectMenuWithStaticOption) addOption(option option.Option) {
	m.options = append(m.options, option)
}

func (m *MultiSelectMenuWithStaticOption) removeOptions() {
	m.options = []option.Option{}
}

// AddOption public add option
func (m MultiSelectMenuWithStaticOption) AddOption(option option.Option) MultiSelectMenuWithStaticOption {
	m.addOption(option)
	return m
}

// ClearOptions clear options
func (m MultiSelectMenuWithStaticOption) ClearOptions() MultiSelectMenuWithStaticOption {
	m.removeOptions()
	return m
}

func (m *MultiSelectMenuWithStaticOption) setOptionGroups(optionGroups []optiongroup.OptionGroup) {
	m.optionGroups = optionGroups
	m.optionals.OptionGroups = true
}

func (m *MultiSelectMenuWithStaticOption) removeOptionGroups() {
	m.optionals.OptionGroups = false
}

// ClearOptionGroups clear option groups
func (m MultiSelectMenuWithStaticOption) ClearOptionGroups() MultiSelectMenuWithStaticOption {
	m.removeOptionGroups()
	return m
}

// AddOptionGroup public add option group
func (m MultiSelectMenuWithStaticOption) AddOptionGroup(optionGroup optiongroup.OptionGroup) MultiSelectMenuWithStaticOption {
	m.setOptionGroups(append(m.optionGroups, optionGroup))
	return m
}

//////////////////////////////////////////////////
// all options

// ClearAllOptions clear all options
func (m MultiSelectMenuWithStaticOption) ClearAllOptions() MultiSelectMenuWithStaticOption {
	m.removeOptions()
	m.removeInitialOptions()
	return m
}

//////////////////////////////////////////////////
// initialOptions

func (m *MultiSelectMenuWithStaticOption) addInitialOption(initialOption option.Option) {
	m.addOption(initialOption)
	m.initialOptions = append(m.initialOptions, initialOption)
	m.optionals.InitialOptions = true
}

func (m *MultiSelectMenuWithStaticOption) removeInitialOptions() {
	m.optionals.InitialOptions = false
}

func (m *MultiSelectMenuWithStaticOption) setInitialOptions(initialOptions []option.Option) {
	m.initialOptions = initialOptions
	m.optionals.InitialOptions = true
}

// ClearInitialOptions clear initial options
func (m MultiSelectMenuWithStaticOption) ClearInitialOptions() MultiSelectMenuWithStaticOption {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m MultiSelectMenuWithStaticOption) AddInitialOption(initialOption option.Option) MultiSelectMenuWithStaticOption {
	m.addInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithStaticOption) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithStaticOption) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithStaticOption) AddConfirmDialog(confirm confirmationdialog.ConfirmationDialog) MultiSelectMenuWithStaticOption {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *MultiSelectMenuWithStaticOption) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenuWithStaticOption) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithStaticOption) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithStaticOption) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithStaticOption {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithStaticOption) UnsetMaxSelectedItems() MultiSelectMenuWithStaticOption {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenuWithStaticOption) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithStaticOption) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithStaticOption) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithStaticOption {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithStaticOption) UnsetFocusOnLoad() MultiSelectMenuWithStaticOption {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenuWithStaticOption) setPlaceholder(placeholder string) {
	m.placeholder = compositiontext.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithStaticOption) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithStaticOption) SetPlaceholder(placeholder string) MultiSelectMenuWithStaticOption {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithStaticOption) RemovePlaceholder() MultiSelectMenuWithStaticOption {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// abstract

// abstracted type
type multiSelectMenuWithStaticOptionAbstraction struct {
	Type             string
	ActionId         string
	Options          []option.Option
	OptionGroups     []optiongroup.OptionGroup
	InitialOptions   []option.Option
	Confirm          confirmationdialog.ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      compositiontext.CompositionText

	Optionals multiSelectMenuWithStaticOptionOptions
}

func (m MultiSelectMenuWithStaticOption) abstraction() multiSelectMenuWithStaticOptionAbstraction {
	return multiSelectMenuWithStaticOptionAbstraction{
		Type:             m.slackType.String(),
		ActionId:         m.actionID,
		Options:          m.options,
		OptionGroups:     m.optionGroups,
		InitialOptions:   m.initialOptions,
		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m multiSelectMenuWithStaticOptionAbstraction) Template() string {
	if m.Optionals.OptionGroups {
		return `{"action_id": "{{ .ActionId }}",
	
	"type": "{{ .Type }}",	
	"option_groups": [{{range $index, $option := .OptionGroups}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]

{{if .Optionals.InitialOptions}},
	"initial_options": [{{range $index, $option := .InitialOptions}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]
{{end}}

{{if .Optionals.Placeholder }},
	"placeholder": {{ .Placeholder.Render }}
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
}`
	}

	return `{"action_id": "{{ .ActionId }}",
	
	"type": "{{ .Type }}",	
	"options": [{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]

{{if .Optionals.InitialOptions}},
	"initial_options": [{{range $index, $option := .InitialOptions}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]
{{end}}

{{if .Optionals.Placeholder }},
	"placeholder": {{ .Placeholder.Render }}
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
}`
}

// Render returns json string
func (m MultiSelectMenuWithStaticOption) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

// ElementRender
func (m MultiSelectMenuWithStaticOption) ElementRender() {}

// SectionBlock public section block
func (m MultiSelectMenuWithStaticOption) Section() section.Section {
	s := section.NewSection("newSection").AddAccessory(m)
	return s
}
