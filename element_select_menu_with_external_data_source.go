package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
)

// InputElement

type SelectMenuWithExternalDataSource struct {
	slackType ElementType
	actionID  string

	initialOption composition.Option
	confirm       composition.ConfirmationDialog
	focusOnLoad   bool
	placeholder   composition.CompositionText

	// External Options
	minQueryLength int

	optionals selectMenuWithExternalDataSourceOptions
}

type selectMenuWithExternalDataSourceOptions struct {
	InitialOption bool
	Confirm       bool
	FocusOnLoad   bool
	Placeholder   bool

	// External Options
	MinQueryLength bool
}

func NewSelectMenuWithExternalDataSource(actionId string) SelectMenuWithExternalDataSource {
	return SelectMenuWithExternalDataSource{
		slackType:     SelectMenuWithExternalDataSourceElement,
		actionID:      actionId,
		initialOption: composition.Option{},
		optionals: selectMenuWithExternalDataSourceOptions{
			InitialOption:  false,
			Confirm:        false,
			FocusOnLoad:    false,
			Placeholder:    false,
			MinQueryLength: false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *SelectMenuWithExternalDataSource) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *SelectMenuWithExternalDataSource) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m SelectMenuWithExternalDataSource) UpdateActionId(actionId string) SelectMenuWithExternalDataSource {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// initialOptions

func (m *SelectMenuWithExternalDataSource) removeInitialOption() {
	m.optionals.InitialOption = false
}

func (m *SelectMenuWithExternalDataSource) setInitialOption(initialOption composition.Option) {
	m.initialOption = initialOption
	m.optionals.InitialOption = true
}

// ClearInitialOptions clear initial options
func (m SelectMenuWithExternalDataSource) ClearInitialOption() SelectMenuWithExternalDataSource {
	m.removeInitialOption()
	return m
}

// AddInitialOption public add initial option
func (m SelectMenuWithExternalDataSource) AddInitialOption(initialOption composition.Option) SelectMenuWithExternalDataSource {
	m.setInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithExternalDataSource) setConfirm(confirm composition.ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithExternalDataSource) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithExternalDataSource) AddConfirmDialog(confirm composition.ConfirmationDialog) SelectMenuWithExternalDataSource {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m *SelectMenuWithExternalDataSource) RemoveConfirmDialog() {
	m.optionals.Confirm = false
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *SelectMenuWithExternalDataSource) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m SelectMenuWithExternalDataSource) FocusOnLoad() SelectMenuWithExternalDataSource {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m SelectMenuWithExternalDataSource) UnsetFocusOnLoad() SelectMenuWithExternalDataSource {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *SelectMenuWithExternalDataSource) setPlaceholder(placeholder string) {
	m.placeholder = composition.NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *SelectMenuWithExternalDataSource) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m SelectMenuWithExternalDataSource) AddPlaceholder(placeholder string) SelectMenuWithExternalDataSource {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m SelectMenuWithExternalDataSource) RemovePlaceholder() SelectMenuWithExternalDataSource {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// minQueryLength

func (m *SelectMenuWithExternalDataSource) setMinQueryLength(minQueryLength int) {
	m.minQueryLength = minQueryLength
	m.optionals.MinQueryLength = true
}

func (m *SelectMenuWithExternalDataSource) removeMinQueryLength() {
	m.optionals.MinQueryLength = false
}

// SetMinQueryLength public set min query length
func (m SelectMenuWithExternalDataSource) SetMinQueryLength(minQueryLength int) SelectMenuWithExternalDataSource {
	m.setMinQueryLength(minQueryLength)
	return m
}

// UnsetMinQueryLength public remove min query length
func (m SelectMenuWithExternalDataSource) UnsetMinQueryLength() SelectMenuWithExternalDataSource {
	m.removeMinQueryLength()
	return m
}

// ////////////////////////////////////////////////
// abstract
type selectMenuWithExternalDataSourceAbstraction struct {
	Type     string
	ActionId string

	InitialOption composition.Option
	Confirm       composition.ConfirmationDialog
	FocusOnLoad   bool
	Placeholder   composition.CompositionText

	// External Options
	MinQueryLength int

	Optionals selectMenuWithExternalDataSourceOptions
}

// abstraction
func (m SelectMenuWithExternalDataSource) abstraction() selectMenuWithExternalDataSourceAbstraction {
	return selectMenuWithExternalDataSourceAbstraction{
		Type:           m.slackType.String(),
		ActionId:       m.actionID,
		InitialOption:  m.initialOption,
		Confirm:        m.confirm,
		FocusOnLoad:    m.focusOnLoad,
		Placeholder:    m.placeholder,
		MinQueryLength: m.minQueryLength,
		Optionals:      m.optionals,
	}
}

// template
func (m selectMenuWithExternalDataSourceAbstraction) Template() string {
	return `{
"action_id": "{{ .ActionId }}",
"type": "{{ .Type }}"	

{{if .Optionals.InitialOption}},
	"initial_option": {{.InitialOption.Render}}
{{end}}

{{if .Optionals.MinQueryLength }},
	"min_query_length": {{ .MinQueryLength }}
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

func (m SelectMenuWithExternalDataSource) ElementRender() {}

func (m SelectMenuWithExternalDataSource) Render() string {
	raw := common.Render(m.abstraction())
	return common.Pretty(raw)
}

func (m SelectMenuWithExternalDataSource) Section() block.Section {
	s := block.NewSection("newSection").AddAccessory(m)
	return s
}
