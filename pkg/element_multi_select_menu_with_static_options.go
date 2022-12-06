package flocksofblocks

// InputElement

type MultiSelectMenuWithStaticOption struct {
	slackType ElementType
	actionID  string
	options   []Option

	optionGroups     []OptionGroup
	initialOptions   []Option
	confirm          ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      CompositionText

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
		slackType: MultiSelectMenuWithStaticOptionsElement,
		actionID:  actionId,
		options:   []Option{},
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

func (m *MultiSelectMenuWithStaticOption) setOptions(options []Option) {
	m.options = options
}

func (m *MultiSelectMenuWithStaticOption) addOption(option Option) {
	m.options = append(m.options, option)
}

func (m *MultiSelectMenuWithStaticOption) removeOptions() {
	m.options = []Option{}
}

// AddOption public add option
func (m MultiSelectMenuWithStaticOption) AddOption(option Option) MultiSelectMenuWithStaticOption {
	m.addOption(option)
	return m
}

// ClearOptions clear options
func (m MultiSelectMenuWithStaticOption) ClearOptions() MultiSelectMenuWithStaticOption {
	m.removeOptions()
	return m
}

func (m *MultiSelectMenuWithStaticOption) setOptionGroups(optionGroups []OptionGroup) {
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
func (m MultiSelectMenuWithStaticOption) AddOptionGroup(optionGroup OptionGroup) MultiSelectMenuWithStaticOption {
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

func (m *MultiSelectMenuWithStaticOption) addInitialOption(initialOption Option) {
	m.addOption(initialOption)
	m.initialOptions = append(m.initialOptions, initialOption)
	m.optionals.InitialOptions = true
}

func (m *MultiSelectMenuWithStaticOption) removeInitialOptions() {
	m.optionals.InitialOptions = false
}

func (m *MultiSelectMenuWithStaticOption) setInitialOptions(initialOptions []Option) {
	m.initialOptions = initialOptions
	m.optionals.InitialOptions = true
}

// ClearInitialOptions clear initial options
func (m MultiSelectMenuWithStaticOption) ClearInitialOptions() MultiSelectMenuWithStaticOption {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m MultiSelectMenuWithStaticOption) AddInitialOption(initialOption Option) MultiSelectMenuWithStaticOption {
	m.addInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithStaticOption) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithStaticOption) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithStaticOption) AddConfirmDialog(confirm ConfirmationDialog) MultiSelectMenuWithStaticOption {
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
	m.placeholder = NewPlainText(placeholder)
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
	Options          []Option
	OptionGroups     []OptionGroup
	InitialOptions   []Option
	Confirm          ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      CompositionText

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
	return `{
"type": "{{ .Type }}",
"action_id": "{{ .ActionId }}",
	
{{if .Optionals.OptionGroups }}	
	"option_groups": [{{range $index, $option := .OptionGroups}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]
{{else}}
	"options": [{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]
{{end}}

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
	raw := Render(m.abstraction())
	return Pretty(raw)
}

// ElementRender interface implementation
func (m MultiSelectMenuWithStaticOption) ElementRender() {}

// Section public section block
func (m MultiSelectMenuWithStaticOption) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}
