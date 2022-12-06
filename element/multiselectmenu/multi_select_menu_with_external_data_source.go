package multiselectmenu

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/confirmationdialog"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/option"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

// InputElement

type MultiSelectMenuWithExternalDataSource struct {
	slackType element.ElementType
	actionID  string

	initialOptions   []option.Option
	confirm          confirmationdialog.ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      compositiontext.CompositionText

	// External Options
	minQueryLength int

	optionals multiSelectMenuWithExternalDataSourceOptions
}

type multiSelectMenuWithExternalDataSourceOptions struct {
	InitialOptions   bool
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool

	// External Options
	MinQueryLength bool
}

func NewMultiSelectMenuWithExternalDataSource(actionId string) MultiSelectMenuWithExternalDataSource {
	return MultiSelectMenuWithExternalDataSource{
		slackType:      element.MultiSelectMenuWithExternalDataSource,
		actionID:       actionId,
		initialOptions: []option.Option{},
		optionals: multiSelectMenuWithExternalDataSourceOptions{
			InitialOptions:   false,
			Confirm:          false,
			MaxSelectedItems: false,
			FocusOnLoad:      false,
			Placeholder:      false,
			MinQueryLength:   false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenuWithExternalDataSource) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithExternalDataSource) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenuWithExternalDataSource) UpdateActionId(actionId string) MultiSelectMenuWithExternalDataSource {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// initialOptions

func (m *MultiSelectMenuWithExternalDataSource) addInitialOption(initialOption option.Option) {
	m.initialOptions = append(m.initialOptions, initialOption)
	m.optionals.InitialOptions = true
}

func (m *MultiSelectMenuWithExternalDataSource) removeInitialOptions() {
	m.optionals.InitialOptions = false
}

func (m *MultiSelectMenuWithExternalDataSource) setInitialOptions(initialOptions []option.Option) {
	m.initialOptions = initialOptions
	m.optionals.InitialOptions = true
}

// ClearInitialOptions clear initial options
func (m MultiSelectMenuWithExternalDataSource) ClearInitialOptions() MultiSelectMenuWithExternalDataSource {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m MultiSelectMenuWithExternalDataSource) AddInitialOption(initialOption option.Option) MultiSelectMenuWithExternalDataSource {
	m.addInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithExternalDataSource) setConfirm(confirm confirmationdialog.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithExternalDataSource) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithExternalDataSource) AddConfirmDialog(confirm confirmationdialog.ConfirmationDialog) MultiSelectMenuWithExternalDataSource {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *MultiSelectMenuWithExternalDataSource) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenuWithExternalDataSource) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithExternalDataSource) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithExternalDataSource) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithExternalDataSource {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithExternalDataSource) UnsetMaxSelectedItems() MultiSelectMenuWithExternalDataSource {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenuWithExternalDataSource) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithExternalDataSource) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithExternalDataSource) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithExternalDataSource {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithExternalDataSource) UnsetFocusOnLoad() MultiSelectMenuWithExternalDataSource {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenuWithExternalDataSource) setPlaceholder(placeholder string) {
	m.placeholder = compositiontext.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithExternalDataSource) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithExternalDataSource) AddPlaceholder(placeholder string) MultiSelectMenuWithExternalDataSource {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithExternalDataSource) RemovePlaceholder() MultiSelectMenuWithExternalDataSource {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// minQueryLength

func (m *MultiSelectMenuWithExternalDataSource) setMinQueryLength(minQueryLength int) {
	m.minQueryLength = minQueryLength
	m.optionals.MinQueryLength = true
}

func (m *MultiSelectMenuWithExternalDataSource) removeMinQueryLength() {
	m.optionals.MinQueryLength = false
}

// MinQueryLength public set min query length
func (m MultiSelectMenuWithExternalDataSource) MinQueryLength(minQueryLength int) MultiSelectMenuWithExternalDataSource {
	m.setMinQueryLength(minQueryLength)
	return m
}

// UnsetMinQueryLength public remove min query length
func (m MultiSelectMenuWithExternalDataSource) UnsetMinQueryLength() MultiSelectMenuWithExternalDataSource {
	m.removeMinQueryLength()
	return m
}

// ////////////////////////////////////////////////
// abstract
type multiSelectMenuWithExternalDataSourceAbstraction struct {
	Type     string
	ActionId string

	InitialOptions   []option.Option
	Confirm          confirmationdialog.ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      compositiontext.CompositionText

	// External Options
	MinQueryLength int

	Optionals multiSelectMenuWithExternalDataSourceOptions
}

// abstraction
func (m MultiSelectMenuWithExternalDataSource) abstraction() multiSelectMenuWithExternalDataSourceAbstraction {
	return multiSelectMenuWithExternalDataSourceAbstraction{
		Type:             m.slackType.String(),
		ActionId:         m.actionID,
		InitialOptions:   m.initialOptions,
		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,
		MinQueryLength:   m.minQueryLength,
		Optionals:        m.optionals,
	}
}

// template
func (m multiSelectMenuWithExternalDataSourceAbstraction) Template() string {
	return `{
"action_id": "{{ .ActionId }}",
		
"type": "{{ .Type }}"	

{{if .Optionals.InitialOptions}},
	"initial_options": [{{range $index, $option := .InitialOptions}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]
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

{{if .Optionals.MinQueryLength }},
	"min_query_length": {{ .MinQueryLength }}
{{end}}
}`
}

func (m MultiSelectMenuWithExternalDataSource) ElementRender() {}

func (m MultiSelectMenuWithExternalDataSource) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

func (m MultiSelectMenuWithExternalDataSource) Section() block.Section {
	s := block.NewSection("newSection").AddAccessory(m)
	return s
}
