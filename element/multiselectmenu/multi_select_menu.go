package multiselectmenu

import (
	"go-flocks-of-blocks/block/section"
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/compositiontext"
	"go-flocks-of-blocks/composition/confirmationdialog"
	"go-flocks-of-blocks/composition/filter"
	"go-flocks-of-blocks/composition/option"
	"go-flocks-of-blocks/composition/optiongroup"
	"go-flocks-of-blocks/element"
)

// InputElement

type MultiSelectMenu struct {
	slackType element.ElementType
	actionID  string
	options   []option.Option

	optionGroups     []optiongroup.OptionGroup
	initialOptions   []option.Option
	confirm          confirmationdialog.ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      compositiontext.CompositionText

	// External Options
	minQueryLength int

	// User List
	initialUsers []string

	// Conversation
	defaultToCurrentConversation bool
	initialConversations         []string
	filter                       filter.Filter

	// Public Channel
	initialChannels []string

	optionals multiSelectMenuOptions
}

type multiSelectMenuOptions struct {
	OptionGroups     bool
	InitialOptions   bool
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool

	// External Options
	MinQueryLength bool

	// User List
	InitialUsers bool

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversations         bool
	Filter                       bool

	// Public Channel
	InitialChannels bool
}

// abstracted type
type multiSelectMenuAbstraction struct {
	Type             string
	ActionId         string
	Options          []option.Option
	OptionGroups     []optiongroup.OptionGroup
	InitialOptions   []option.Option
	Confirm          confirmationdialog.ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      compositiontext.CompositionText

	// External Options
	MinQueryLength int

	// User List
	InitialUsers []string

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversations         []string
	Filter                       filter.Filter

	// Public Channel
	InitialChannels []string

	Optionals multiSelectMenuOptions
}

func NewMultiSelectMenuWithConversationsList(actionId string) MultiSelectMenu {
	return MultiSelectMenu{
		slackType: element.MultiSelectMenuWithConversationsList,
		actionID:  actionId,
		options:   []option.Option{},
		optionals: multiSelectMenuOptions{},
	}
}
func NewMultiSelectMenuWithPublicChannelsSelect(actionId string) MultiSelectMenu {
	return MultiSelectMenu{
		slackType: element.MultiSelectMenuWithPublicChannelsSelect,
		actionID:  actionId,
		options:   []option.Option{},
		optionals: multiSelectMenuOptions{},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenu) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenu) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenu) UpdateActionId(actionId string) MultiSelectMenu {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// options

func (m *MultiSelectMenu) setOptions(options []option.Option) {
	m.options = options
}

func (m *MultiSelectMenu) addOption(option option.Option) {
	m.options = append(m.options, option)
}

func (m *MultiSelectMenu) removeOptions() {
	m.options = []option.Option{}
}

// AddOption public add option
func (m MultiSelectMenu) AddOption(option option.Option) MultiSelectMenu {
	m.addOption(option)
	return m
}

// ClearOptions clear options
func (m MultiSelectMenu) ClearOptions() MultiSelectMenu {
	m.removeOptions()
	return m
}

func (m *MultiSelectMenu) setOptionGroups(optionGroups []optiongroup.OptionGroup) {
	m.optionGroups = optionGroups
	m.optionals.OptionGroups = true
}

func (m *MultiSelectMenu) removeOptionGroups() {
	m.optionals.OptionGroups = false
}

// ClearOptionGroups clear option groups
func (m MultiSelectMenu) ClearOptionGroups() MultiSelectMenu {
	m.removeOptionGroups()
	return m
}

// AddOptionGroup public add option group
func (m MultiSelectMenu) AddOptionGroup(optionGroup optiongroup.OptionGroup) MultiSelectMenu {
	m.setOptionGroups(append(m.optionGroups, optionGroup))
	return m
}

//////////////////////////////////////////////////
// all options

// ClearAllOptions clear all options
func (m MultiSelectMenu) ClearAllOptions() MultiSelectMenu {
	m.removeOptions()
	m.removeInitialOptions()
	return m
}

//////////////////////////////////////////////////
// initialOptions

func (m *MultiSelectMenu) addInitialOption(initialOption option.Option) {
	m.addOption(initialOption)
	m.initialOptions = append(m.initialOptions, initialOption)
	m.optionals.InitialOptions = true
}

func (m *MultiSelectMenu) removeInitialOptions() {
	m.optionals.InitialOptions = false
}

func (m *MultiSelectMenu) setInitialOptions(initialOptions []option.Option) {
	m.initialOptions = initialOptions
	m.optionals.InitialOptions = true
}

// ClearInitialOptions clear initial options
func (m MultiSelectMenu) ClearInitialOptions() MultiSelectMenu {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m MultiSelectMenu) AddInitialOption(initialOption option.Option) MultiSelectMenu {
	m.addInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenu) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenu) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenu) AddConfirmDialog(confirm confirmationdialog.ConfirmationDialog) MultiSelectMenu {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *MultiSelectMenu) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenu) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenu) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenu) MaxSelectedItems(maxSelectedItems int) MultiSelectMenu {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenu) UnsetMaxSelectedItems() MultiSelectMenu {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenu) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenu) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenu) FocusOnLoad(focusOnLoad bool) MultiSelectMenu {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenu) UnsetFocusOnLoad() MultiSelectMenu {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenu) setPlaceholder(placeholder string) {
	m.placeholder = compositiontext.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenu) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenu) AddPlaceholder(placeholder string) MultiSelectMenu {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenu) RemovePlaceholder() MultiSelectMenu {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// minQueryLength

func (m *MultiSelectMenu) setMinQueryLength(minQueryLength int) {
	m.minQueryLength = minQueryLength
	m.optionals.MinQueryLength = true
}

func (m *MultiSelectMenu) removeMinQueryLength() {
	m.optionals.MinQueryLength = false
}

// MinQueryLength public set min query length
func (m MultiSelectMenu) MinQueryLength(minQueryLength int) MultiSelectMenu {
	m.setMinQueryLength(minQueryLength)
	return m
}

// UnsetMinQueryLength public remove min query length
func (m MultiSelectMenu) UnsetMinQueryLength() MultiSelectMenu {
	m.removeMinQueryLength()
	return m
}

//////////////////////////////////////////////////
// initialUsers

func (m *MultiSelectMenu) addInitialUser(initialUser string) {
	m.initialUsers = append(m.initialUsers, initialUser)
	m.optionals.InitialUsers = true
}

func (m *MultiSelectMenu) removeInitialUsers() {
	m.initialUsers = []string{}
	m.optionals.InitialUsers = false
}

// AddInitialUser public add initial user
func (m MultiSelectMenu) AddInitialUser(initialUser string) MultiSelectMenu {
	m.addInitialUser(initialUser)
	return m
}

// ClearInitialUsers clear initial users
func (m MultiSelectMenu) ClearInitialUsers() MultiSelectMenu {
	m.removeInitialUsers()
	return m
}

//////////////////////////////////////////////////
// defaultToCurrentConversation

// setDefaultToCurrentConversation public set default to current conversation
func (m *MultiSelectMenu) setDefaultToCurrentConversation(defaultToCurrentConversation bool) {
	m.defaultToCurrentConversation = defaultToCurrentConversation
	m.optionals.DefaultToCurrentConversation = defaultToCurrentConversation

}

// unsetDefaultToCurrentConversation public remove default to current conversation
func (m *MultiSelectMenu) unsetDefaultToCurrentConversation() {
	m.setDefaultToCurrentConversation(false)
}

// DefaultToCurrentConversation public set default to current conversation
func (m MultiSelectMenu) DefaultToCurrentConversation() MultiSelectMenu {
	m.setDefaultToCurrentConversation(true)
	return m
}

// UnsetDefaultToCurrentConversation public remove default to current conversation
func (m MultiSelectMenu) UnsetDefaultToCurrentConversation() MultiSelectMenu {
	m.unsetDefaultToCurrentConversation()
	return m
}

//////////////////////////////////////////////////
// initialConversations

// addInitialConversation private add initial conversation
func (m *MultiSelectMenu) addInitialConversation(initialConversation string) {
	m.initialConversations = append(m.initialConversations, initialConversation)
	m.optionals.InitialConversations = true
}

// removeInitialConversations private remove initial conversations
func (m *MultiSelectMenu) removeInitialConversations() {
	m.initialConversations = []string{}
	m.optionals.InitialConversations = false
}

// AddInitialConversation public add initial conversation
func (m MultiSelectMenu) AddInitialConversation(initialConversation string) MultiSelectMenu {
	m.addInitialConversation(initialConversation)
	return m
}

// ClearInitialConversations clear initial conversations
func (m MultiSelectMenu) ClearInitialConversations() MultiSelectMenu {
	m.removeInitialConversations()
	return m
}

//////////////////////////////////////////////////
// filter

func (m *MultiSelectMenu) setFilter(filter filter.Filter) {
	m.filter = filter
	m.optionals.Filter = true
}

func (m *MultiSelectMenu) removeFilter() {
	m.optionals.Filter = false
}

// AddFilter public set filter
func (m MultiSelectMenu) AddFilter(filter filter.Filter) MultiSelectMenu {
	m.setFilter(filter)
	return m
}

// RemoveFilter public remove filter
func (m MultiSelectMenu) RemoveFilter() MultiSelectMenu {
	m.removeFilter()
	return m
}

//////////////////////////////////////////////////
// initialChannels

// create abstract
func (m MultiSelectMenu) abstraction() multiSelectMenuAbstraction {
	return multiSelectMenuAbstraction{
		Type:             m.slackType.String(),
		ActionId:         m.actionID,
		Options:          m.options,
		OptionGroups:     m.optionGroups,
		InitialOptions:   m.initialOptions,
		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,

		// External Options
		MinQueryLength: m.minQueryLength,

		// User List
		InitialUsers: m.initialUsers,

		// Conversation
		DefaultToCurrentConversation: m.defaultToCurrentConversation,
		InitialConversations:         m.initialConversations,
		Filter:                       m.filter,

		// Public Channel
		InitialChannels: m.initialChannels,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m multiSelectMenuAbstraction) Template() string {
	return `"action_id": "{{ .ActionId }}",
"type": "{{ .Type }}"{{if 

	}`
}

func (m MultiSelectMenu) ElementRender() {}

func (m MultiSelectMenu) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

func (m MultiSelectMenu) Section() section.Section {
	s := section.NewSection("newSection").AddAccessory(m)
	return s
}
