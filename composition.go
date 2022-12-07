package flocksofblocks

import "net/url"

type Composition interface {
	compositeRender() string
}

type CompositionType string

// stringer
func (c CompositionType) String() string {
	return string(c)
}

const (
	PlainText CompositionType = "plain_text"
	Mrkdwn    CompositionType = "mrkdwn"
)

///////////////////////////////
// CompositionText

type CompositionText struct {
	slackType CompositionType
	text      string
	emoji     bool
	verbatim  bool
}

// NewPlainText Create a new markdown object
func NewPlainText(text string) CompositionText {
	return newText("plain_text", text)
}

// NewMrkdwnText Create a new markdown compositiontext object
func NewMrkdwnText(text string) CompositionText {
	return newText("mrkdwn", text)
}

func newText(slackType CompositionType, text string) CompositionText {
	return CompositionText{
		slackType: slackType,
		text:      text,
	}
}

// SetEmoji set the emoji flag
func (m *CompositionText) setEmoji(emoji bool) {
	m.emoji = emoji
}

// EnableEmoji enable the emoji flag
func (m CompositionText) EnableEmoji() CompositionText {
	m.setEmoji(true)
	return m
}

// SetVerbatim set the verbatim flag
func (m *CompositionText) SetVerbatim(verbatim bool) {
	m.verbatim = verbatim
}

// Stringer that returns the compositiontext
func (m CompositionText) String() string {
	return m.text
}

// compositionTextAbstraction is used to render the composition compositiontext
type compositionTextAbstraction struct {
	Type     string
	Text     string
	Emoji    bool
	Verbatim bool
}

// abstraction for the text
func (m CompositionText) abstraction() compositionTextAbstraction {
	return compositionTextAbstraction{
		Type:     m.slackType.String(),
		Text:     m.text,
		Emoji:    m.emoji,
		Verbatim: m.verbatim,
	}
}

func (m compositionTextAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"text": "{{.Text}}"{{if .Emoji}},
	"emoji": {{.Emoji}}{{end}}{{if .Verbatim}}, 
	"verbatim": {{.Verbatim}}{{end}}
}`
}

// Render the compositiontext
func (m CompositionText) Render() string {
	return Render(m.abstraction())
}

///////////////////////////////
// ConfirmationDialog

type ConfirmationDialog struct {
	title CompositionText

	// text is a CompositionText object which can be either a PlainText or a MarkdownText
	text    CompositionText
	confirm CompositionText
	deny    CompositionText

	style ColorSchema

	optionals ConfirmationDialogOptions
}

// todo make abstraction for this

// ConfirmationDialogOptions struct
type ConfirmationDialogOptions struct {
	Style bool
}

// NewConfirmationDialog creates a new confirmation dialog
// todo: might consider making better input names
func NewConfirmationDialog(title string, text string, confirm string, deny string) ConfirmationDialog {

	return ConfirmationDialog{
		title:   NewPlainText(title),
		text:    NewPlainText(text),
		confirm: NewPlainText(confirm),
		deny:    NewPlainText(deny),
		optionals: ConfirmationDialogOptions{
			Style: false,
		},
	}
}

// set the style
func (c *ConfirmationDialog) setStyle(style ColorSchema) {
	c.style = style
	c.optionals.Style = true
}

// set the style public
func (c ConfirmationDialog) SetStyle(style ColorSchema) ConfirmationDialog {
	c.setStyle(style)
	return c
}

// confirmationDialogAbstraction is used to render the confirmation dialog
type confirmationDialogAbstraction struct {
	Title   CompositionText
	Text    CompositionText
	Confirm CompositionText
	Deny    CompositionText
	Style   string

	Optional ConfirmationDialogOptions
}

// create an abstraction for the template
func (c ConfirmationDialog) abstraction() confirmationDialogAbstraction {
	return confirmationDialogAbstraction{
		Title:    c.title,
		Text:     c.text,
		Confirm:  c.confirm,
		Deny:     c.deny,
		Style:    c.style.String(),
		Optional: c.optionals,
	}
}

// create the template
func (c confirmationDialogAbstraction) Template() string {
	return `{
	"title": {{.Title.Render}},
	"text": {{.Text.Render}},
	"confirm": {{.Confirm.Render}},
	"deny": {{.Deny.Render}}
{{if .Optionals.Style}},	
	"style": "{{.Style}}"
{{end}}
}`
}

// Render the template
func (c ConfirmationDialog) Render() string {
	raw := Render(c.abstraction())
	return Pretty(raw)
}

///////////////////////////////
// DispatchActionTypes

type DispatchActionTypes string

const (
	OnEnterPressed     DispatchActionTypes = "on_enter_pressed"
	OnCharacterEntered DispatchActionTypes = "on_character_entered"
)

type DispatchActionConfig struct {
	triggerActionsOn []DispatchActionTypes
}

type abstractionDispatchActionConfig struct {
	TriggerActionsOn []string
}

func NewDispatchActionConfig() DispatchActionConfig {
	return DispatchActionConfig{
		triggerActionsOn: []DispatchActionTypes{},
	}
}

func (d *DispatchActionConfig) setTriggerActionsOn(triggerActionsOn DispatchActionTypes) {
	d.triggerActionsOn = append(d.triggerActionsOn, triggerActionsOn)
}

func (d *DispatchActionConfig) removeTriggerActionsOn() {
	d.triggerActionsOn = []DispatchActionTypes{}
}

// OnEnterPressed chain function to add on_enter_pressed to an existing dispatch action config
func (d DispatchActionConfig) OnEnterPressed() DispatchActionConfig {
	d.setTriggerActionsOn(OnEnterPressed)
	return d
}

// OnCharacterEntered chain function to add on_character_entered to an existing dispatch action config
func (d DispatchActionConfig) OnCharacterEntered() DispatchActionConfig {
	d.setTriggerActionsOn(OnCharacterEntered)
	return d
}

// RemoveTriggerActionsOn remove add trigger actions on from dispatch action config
func (d DispatchActionConfig) RemoveTriggerActions() DispatchActionConfig {
	d.removeTriggerActionsOn()
	return d
}

// Template generates the template for the block
func (d abstractionDispatchActionConfig) Template() string {
	return `{
"trigger_actions_on": [{{range $index, $element := .TriggerActionsOn}}{{if $index}}, {{end}}"{{$element}}"{{end}}]
}`
}

// abstraction
func (d DispatchActionConfig) abstraction() abstractionDispatchActionConfig {
	return abstractionDispatchActionConfig{
		TriggerActionsOn: removeDuplicateStr(d.triggerActionsOn),
	}
}

// Render the block
func (d DispatchActionConfig) Render() string {
	output := Render(d.abstraction())
	return Pretty(output)
}

func removeDuplicateStr(strSlice []DispatchActionTypes) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		item := string(item)
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

///////////////////////////////
// Filter

type Filter struct {
	include                       []string
	excludeExternalSharedChannels bool
	excludeBotUsers               bool

	optionals filterOptions
}

type filterOptions struct {
	Include                       bool
	ExcludeExternalSharedChannels bool
	ExcludeBotUsers               bool
}

func NewFilter() Filter {
	return Filter{
		include:                       []string{},
		excludeExternalSharedChannels: false,
		excludeBotUsers:               false,
		optionals: filterOptions{
			Include:                       false,
			ExcludeExternalSharedChannels: false,
			ExcludeBotUsers:               false,
		},
	}
}

// abstracted type
type filterAbstraction struct {
	Include                       []string
	ExcludeExternalSharedChannels bool
	ExcludeBotUsers               bool

	Optionals filterOptions
}

// add include to filter
func (f *Filter) addInclude(include string) {
	f.include = append(f.include, include)
	f.optionals.Include = true
}

// AddInclude add filter string
func (f Filter) AddInclude(include string) Filter {
	f.addInclude(include)
	return f
}

// IncludeIM Add IM to include filter
func (f Filter) IncludeIM() Filter {
	return f.AddInclude("im")
}

// IncludeMPIM Add MPIM to include filter
func (f Filter) IncludeMPIM() Filter {
	return f.AddInclude("mpim")
}

// IncludePrivate Add private to include filter
func (f Filter) IncludePrivate() Filter {
	return f.AddInclude("private")
}

// IncludePublic Add public to include filter
func (f Filter) IncludePublic() Filter {
	return f.AddInclude("public")
}

// ClearInclude clear include
func (f *Filter) clearInclude() {
	f.include = []string{}
	f.optionals.Include = false
}

// ClearInclude clear include
func (f Filter) ClearInclude() Filter {
	f.clearInclude()
	return f
}

// set exclude external shared channels
func (f *Filter) setExcludeExternalSharedChannels(excludeExternalSharedChannels bool) {
	f.excludeExternalSharedChannels = excludeExternalSharedChannels
	f.optionals.ExcludeExternalSharedChannels = excludeExternalSharedChannels
}

// ExcludeExternalSharedChannels exclude external shared channels
func (f Filter) ExcludeExternalSharedChannels() Filter {
	f.setExcludeExternalSharedChannels(true)
	return f
}

// UnsetExcludeExternalSharedChannels unset exclude external shared channels
func (f Filter) UnsetExcludeExternalSharedChannels() Filter {
	f.setExcludeExternalSharedChannels(false)
	return f
}

// set exclude bot users
func (f *Filter) setExcludeBotUsers(excludeBotUsers bool) {
	f.excludeBotUsers = excludeBotUsers
	f.optionals.ExcludeBotUsers = excludeBotUsers
}

// ExcludeBotUsers exclude bot users
func (f Filter) ExcludeBotUsers() Filter {
	f.setExcludeBotUsers(true)
	return f
}

// UnsetExcludeBotUsers unset exclude bot users
func (f Filter) UnsetExcludeBotUsers() Filter {
	f.setExcludeBotUsers(false)
	return f
}

// abstraction
func (f Filter) abstraction() filterAbstraction {
	return filterAbstraction{
		Include:                       removeDuplicateString(f.include),
		ExcludeExternalSharedChannels: f.excludeExternalSharedChannels,
		ExcludeBotUsers:               f.excludeBotUsers,
		Optionals:                     f.optionals,
	}
}

func (f filterAbstraction) empty() bool {
	if f.Optionals.Include {
		return false
	}
	if f.Optionals.ExcludeExternalSharedChannels {
		return false
	}
	if f.Optionals.ExcludeBotUsers {
		return false
	}
	return true
}

// template
func (f filterAbstraction) Template() string {
	if f.empty() {
		return ""
	}
	return `"filter": {
	{{if .Optionals.Include}}"include": [{{range $index, $include := .Include}}{{if $index}},{{end}}"{{ $include}}"{{end}}]{{if .Optionals.ExcludeExternalSharedChannels}},{{end}}{{end}}{{if .Optionals.ExcludeExternalSharedChannels}}
	"exclude_external_shared_channels": {{.ExcludeExternalSharedChannels}}{{if .Optionals.ExcludeBotUsers }},{{end}}{{end}}{{if .Optionals.ExcludeBotUsers }}
	"exclude_bot_users": {{.ExcludeBotUsers}}{{end}}
}`
}

// Render method
func (f Filter) Render() string {
	return Render(f.abstraction())
}

// im, mpim, private, and public.

type Option struct {
	// Required
	text  CompositionText
	value string

	// Optionals
	description CompositionText
	url         url.URL

	optionals optional
}

type optional struct {
	Description bool
	Url         bool
}

type optionOption func(*Option)

// NewOption creates a new option.
func NewOption(text string, value string) Option {
	return Option{
		text:  NewPlainText(text),
		value: value,
	}
}

func (o *Option) setDescription(description CompositionText) {
	o.description = description
	o.optionals.Description = true
}

func (o *Option) setUrl(url url.URL) {
	o.url = url
	o.optionals.Url = true
}

func (o Option) SetDescription(description CompositionText) Option {
	o.setDescription(description)
	return o
}

func (o *Option) SetUrl(url url.URL) *Option {
	o.setUrl(url)
	return o
}

func (o Option) RemoveDescription() Option {
	o.optionals.Description = false
	return o
}

func (o Option) RemoveUrl() Option {
	o.optionals.Url = false
	return o
}

// optionAbstraction is used to render the option
type optionAbstraction struct {
	Text        CompositionText
	Value       string
	Description CompositionText
	Url         string
	Optionals   optional
}

// create an option abstraction for rendering
func (o Option) abstraction() optionAbstraction {
	url := ""
	if o.optionals.Url {
		url = o.url.String()
	}
	return optionAbstraction{
		Text:        o.text,
		Value:       o.value,
		Description: o.description,
		Url:         url,
		Optionals:   o.optionals,
	}
}

func (o optionAbstraction) Template() string {
	return `{
	"text": {{ .Text.Render}},
	"value": "{{.Value}}"{{if .Optionals.Description}},
	"description": {{.Description.Render}}{{end}}{{if .Optionals.Url}},
	"url": "{{.Url}}"{{end}}	
}`
}

// Render renders the option to a string.
func (o Option) Render() string {
	return Render(o.abstraction())
}

/////////////////////////
// OptionGroup

type OptionGroup struct {
	label   CompositionText
	options []Option
}

func NewOptionGroup(label string) OptionGroup {
	return OptionGroup{
		label:   NewPlainText(label),
		options: []Option{},
	}
}

// setLabel sets the label for the block.
func (o *OptionGroup) setLabel(label CompositionText) {
	o.label = label
}

// SetLabel sets the label for the block.
func (o OptionGroup) SetLabel(label string) OptionGroup {
	o.setLabel(NewPlainText(label))
	return o
}

func (o OptionGroup) AddOption(option Option) OptionGroup {
	o.options = append(o.options, option)
	return o
}

// compositionOptionAbstraction is used to render the composition option
type optionGroupAbstraction struct {
	Label   CompositionText
	Options []Option
}

// generate the abstraction for the block
func (o OptionGroup) abstraction() optionGroupAbstraction {
	return optionGroupAbstraction{
		Label:   o.label,
		Options: o.options,
	}
}

// Render renders the block to a string.
func (o OptionGroup) Render() string {
	return Render(o.abstraction())
}

// Template returns the template for the block.
func (o optionGroupAbstraction) Template() string {
	return `{
		"label": {{.Label.Render}},
		"options": [
			{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}
		]
	}`
}
