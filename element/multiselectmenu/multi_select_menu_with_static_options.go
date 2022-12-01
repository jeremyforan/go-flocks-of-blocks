package multiselectmenu

import (
	"go-flocks-of-blocks/composition/compositiontext"
	"go-flocks-of-blocks/composition/confirmationdialog"
	"go-flocks-of-blocks/composition/option"
	"go-flocks-of-blocks/composition/optiongroup"
	"go-flocks-of-blocks/element"
)

// InputElement

type MultiSelectMenuWithStaticOptions struct {
	slackType element.ElementType
	actionID  string
	options   []option.Option

	optionGroups     []optiongroup.OptionGroup
	initialOptions   []option.Option
	confirm          confirmationdialog.ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      compositiontext.CompositionText

	optionals multiSelectMenuWithStaticOptionsOptions
}

type multiSelectMenuWithStaticOptionsOptions struct {
	OptionGroups     bool
	InitialOptions   bool
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool
}

// abstracted type
type multiSelectMenuWithStaticOptionsAbstraction struct {
	Type             string
	ActionId         string
	Options          []option.Option
	OptionGroups     []optiongroup.OptionGroup
	InitialOptions   []option.Option
	Confirm          confirmationdialog.ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      compositiontext.CompositionText

	Optionals multiSelectMenuWithStaticOptionsOptions
}

func NewMultiSelectMenuWithStaticOptions(actionId string) MultiSelectMenuWithStaticOptions {
	return MultiSelectMenuWithStaticOptions{
		slackType: element.MultiSelectMenuWithStaticOptions,
		actionID:  actionId,
		options:   []option.Option{},
		optionals: multiSelectMenuWithStaticOptionsOptions{
			OptionGroups:     false,
			InitialOptions:   false,
			Confirm:          false,
			MaxSelectedItems: false,
			FocusOnLoad:      false,
			Placeholder:      false,
		},
	}
}

// action_id methods

func (m *MultiSelectMenuWithStaticOptions) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithStaticOptions) removeActionId() {
	m.actionID = ""
}

// private options methods

func (m *MultiSelectMenuWithStaticOptions) setOptions(options []option.Option) {
	m.options = options
}

func (m *MultiSelectMenuWithStaticOptions) addOption(option option.Option) {
	// todo: check if option is already in the list
	m.options = append(m.options, option)
}

func (m *MultiSelectMenuWithStaticOptions) removeOptions() {
	m.options = []option.Option{}
}

// ClearOptions clear options
func (m MultiSelectMenuWithStaticOptions) ClearOptions() MultiSelectMenuWithStaticOptions {
	m.removeOptions()
	return m
}

// AddOption public add option
func (m MultiSelectMenuWithStaticOptions) AddOption(option option.Option) MultiSelectMenuWithStaticOptions {
	m.addOption(option)
	return m
}

// private options group methods
func (m *MultiSelectMenuWithStaticOptions) setOptionGroups(optionGroups []optiongroup.OptionGroup) {
	m.optionGroups = optionGroups
	m.optionals.OptionGroups = true
}

func (m *MultiSelectMenuWithStaticOptions) removeOptionGroups() {
	m.optionals.OptionGroups = false
}

func (m MultiSelectMenuWithStaticOptions) ClearOptionGroups() MultiSelectMenuWithStaticOptions {
	m.removeOptionGroups()
	return m
}

// AddOptionGroup public add option group
func (m MultiSelectMenuWithStaticOptions) AddOptionGroup(optionGroup optiongroup.OptionGroup) MultiSelectMenuWithStaticOptions {
	m.setOptionGroups(append(m.optionGroups, optionGroup))
	return m
}

// private initial options methods

func (m *MultiSelectMenuWithStaticOptions) addInitialOption(initialOption option.Option) {
	m.addOption(initialOption)
	m.initialOptions = append(m.initialOptions, initialOption)
	m.optionals.InitialOptions = true
}

func (m *MultiSelectMenuWithStaticOptions) removeInitialOptions() {
	m.optionals.InitialOptions = false
}

func (m *MultiSelectMenuWithStaticOptions) setInitialOptions(initialOptions []option.Option) {
	m.initialOptions = initialOptions
	m.optionals.InitialOptions = true
}

// ClearInitialOptions clear initial options
func (m MultiSelectMenuWithStaticOptions) ClearInitialOptions() MultiSelectMenuWithStaticOptions {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m MultiSelectMenuWithStaticOptions) AddInitialOption(initialOption option.Option) MultiSelectMenuWithStaticOptions {
	m.addInitialOption(initialOption)
	return m
}

// private confirm methods

func (m *MultiSelectMenuWithStaticOptions) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithStaticOptions) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithStaticOptions) AddConfirmDialog(confirm confirmationdialog.ConfirmationDialog) MultiSelectMenuWithStaticOptions {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *MultiSelectMenuWithStaticOptions) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

// private max selected items methods

func (m *MultiSelectMenuWithStaticOptions) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithStaticOptions) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithStaticOptions) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithStaticOptions {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithStaticOptions) UnsetMaxSelectedItems() MultiSelectMenuWithStaticOptions {
	m.optionals.MaxSelectedItems = false
	return m
}

// private max selected items methods

func (m *MultiSelectMenuWithStaticOptions) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithStaticOptions) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithStaticOptions) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithStaticOptions {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithStaticOptions) UnsetFocusOnLoad() MultiSelectMenuWithStaticOptions {
	m.removeFocusOnLoad()
	return m
}

// private placeholder methods

func (m *MultiSelectMenuWithStaticOptions) setPlaceholder(placeholder string) {
	m.placeholder = compositiontext.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithStaticOptions) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithStaticOptions) SetPlaceholder(placeholder string) MultiSelectMenuWithStaticOptions {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithStaticOptions) RemovePlaceholder() MultiSelectMenuWithStaticOptions {
	m.optionals.Placeholder = false
	return m
}

// ClearAllOptions clear all options
func (m MultiSelectMenuWithStaticOptions) ClearAllOptions() MultiSelectMenuWithStaticOptions {
	m.removeOptions()
	m.removeInitialOptions()
	return m
}

// create abstract
func (m MultiSelectMenuWithStaticOptions) abstraction() multiSelectMenuWithStaticOptionsAbstraction {
	return multiSelectMenuWithStaticOptionsAbstraction{
		Type:             m.slackType.String(),
		ActionId:         m.actionID,
		Options:          m.options,
		OptionGroups:     m.optionGroups,
		InitialOptions:   m.initialOptions,
		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,
		Optionals:        m.optionals,
	}
}

// Template returns template string
func (m MultiSelectMenuWithStaticOptions) Template() string {
	return `"action_id": "{{ .ActionId }}",
"type": "{{ .Type }}"{{if 

	}`
}
