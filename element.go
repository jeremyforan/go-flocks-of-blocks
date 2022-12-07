package flocksofblocks

import (
	"net/url"
	"strconv"
	"time"
)

type Element interface {
	ElementRender()
}

type ElementType string

const (
	ButtonElement                                  ElementType = "button"
	CheckboxesElement                              ElementType = "checkboxes"
	DatePickerElement                              ElementType = "datepicker"
	DateTimePickerElement                          ElementType = "datetimepicker"
	EmailInputElement                              ElementType = "email"
	ImageElementType                               ElementType = "image"
	MultiSelectMenuWithStaticOptionsElement        ElementType = "multi_static_select"
	MultiSelectMenuWithExternalDataSourceElement   ElementType = "multi_external_select"
	MultiSelectMenuWithUserListElement             ElementType = "multi_users_select"
	MultiSelectMenuWithConversationsListElement    ElementType = "multi_conversations_select"
	MultiSelectMenuWithPublicChannelsSelectElement ElementType = "multi_channels_select"
	NumberInputElement                             ElementType = "number_input"
	OverflowMenuElement                            ElementType = "overflow"
	PlainTextInputElement                          ElementType = "plain_text_input"
	RadioButtonElement                             ElementType = "radio_buttons"
	SelectMenuWithStaticOptionsElement             ElementType = "static_select"
	SelectMenuWithExternalDataSourceElement        ElementType = "external_select"
	SelectMenuWithUserListElement                  ElementType = "users_select"
	SelectMenuWithConversationsListElement         ElementType = "conversations_select"
	SelectMenuWithPublicChannelsSelectElement      ElementType = "channels_select"
	TimePickerElement                              ElementType = "timepicker"
	UrlInputElement                                ElementType = "url_text_input"
)

// stringer for ElementType
func (e ElementType) String() string {
	return string(e)
}

type InputElement interface {
	InputElement()
}

// Button as defined in slack
type Button struct {
	slackType ElementType
	text      CompositionText
	actionId  string

	// optionals
	url                string
	value              string
	style              ColorSchema
	confirm            ConfirmationDialog
	accessibilityLabel string

	// optionals help with the template rendering
	optionals buttonOptionals
}

// NewButton create a new button element for an action.
func NewButton(text string, actionId string) Button {
	button := Button{
		slackType: ButtonElement,
		text:      NewPlainText(text),
		actionId:  actionId,
		optionals: buttonOptionals{
			Url:                false,
			Value:              false,
			Style:              false,
			Confirm:            false,
			AccessibilityLabel: false,
		},
	}

	return button
}

type buttonAbstraction struct {
	Type     string          // required
	Text     CompositionText // required
	ActionId string          // required

	// optionals
	Url                string
	Value              string
	Style              string
	Confirm            ConfirmationDialog
	AccessibilityLabel string

	// optionals help with the template rendering
	Optionals buttonOptionals
}

func (b Button) Render() string {
	return Render(b.abstraction())
}

// setUrl sets the url for the button.
func (b *Button) setUrl(url *url.URL) {
	b.url = url.String()
	b.optionals.Url = true
}

// removeUrl removes the url from the button.
func (b *Button) removeUrl() {
	b.url = ""
	b.optionals.Url = false
}

// setValue sets the value for the button.
func (b *Button) setValue(value string) {
	b.value = value
	b.optionals.Value = true
}

func (b *Button) removeValue() {
	b.value = ""
	b.optionals.Value = true
}

func (b *Button) setStyle(style ColorSchema) {
	if style == StyleDefault {
		b.optionals.Style = false
	} else {
		b.style = style
		b.optionals.Style = true
	}
}

// setConfirmationDialog sets the confirmation dialog for the button.
func (b *Button) setConfirmationDialog(confirm ConfirmationDialog) {
	b.confirm = confirm
}

// setAccessibilityLabel sets the style for the button.
func (b *Button) setAccessibilityLabel(label string) {
	b.accessibilityLabel = label
	b.optionals.AccessibilityLabel = true
}

// todo: make primary / default / danger

// Template returns the template for the button.
func (b buttonAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"action_id": "{{.ActionId}}",
	"text": {{.Text.Render}}

{{if .Optionals.Url}},
	"url": "{{.Url}}"
{{end}}

{{if .Optionals.Value}},
	"value": "{{.Value}}"
{{end}}

{{if .Optionals.Style}},
	"style": "{{.Style}}"
{{end}}

{{if .Optionals.Confirm}},
	"confirm": {{.Confirm.Render}}
{{end}}

{{if .Optionals.AccessibilityLabel}},
	"accessibility_label": "{{.AccessibilityLabel}}"
{{end}}

}`
}

func (b Button) Section() {}
func (b Button) Action()  {}

func (b Button) ElementRender() {}

type buttonOptionals struct {
	Url                bool
	Value              bool
	Style              bool
	Confirm            bool
	AccessibilityLabel bool
}

// buttonConstructionOptions allows for optional parameters to be passed into the NewButton function.
type buttonConstructionOptions func(*Button)

////////////////////////////////////////////////////////////////////////////////////
// Button Abstraction

func (b *Button) abstraction() buttonAbstraction {
	url := ""
	if b.optionals.Url {
		url = b.url
	}
	return buttonAbstraction{
		Type:               b.slackType.String(),
		Text:               b.text,
		ActionId:           b.actionId,
		Url:                url,
		Value:              b.value,
		Style:              b.style.String(),
		Confirm:            b.confirm,
		AccessibilityLabel: b.accessibilityLabel,
		Optionals:          b.optionals,
	}
}

// AddUrl chain function to add url to an existing button
func (b Button) AddUrl(url *url.URL) Button {
	b.setUrl(url)
	return b
}

// RemoveUrl chain function to remove url from an existing button
func (b Button) RemoveUrl() Button {
	b.removeUrl()
	return b
}

// SetValue sets the value for the button.
func (b Button) SetValue(value string) Button {
	b.setValue(value)
	return b
}

func (b Button) RemoveValue() Button {
	b.removeValue()
	return b
}

// MakeStylePrimary chain method that sets the style of the button to primary.
func (b Button) MakeStylePrimary() Button {
	b.setStyle(StylePrimary)
	return b
}

// MakeStyleDanger invoke option sets the style of the button to primary.
func (b Button) MakeStyleDanger() Button {
	b.setStyle(StyleDanger)
	return b
}

// MakeStyleDefault invoke option sets the style of the button to primary.
func (b Button) MakeStyleDefault() Button {
	b.setStyle(StyleDefault)
	return b
}

// AddConfirmationDialog sets the confirmation dialog for the button.
func (b Button) AddConfirmationDialog(confirm ConfirmationDialog) Button {
	b.setConfirmationDialog(confirm)
	return b
}

// RemoveConfirmationDialog removes the confirmation dialog from the button.
func (b Button) RemoveConfirmationDialog() Button {
	b.confirm = ConfirmationDialog{}
	b.optionals.Confirm = false
	return b
}

// SetAccessibilityLabel sets the style for the button.
func (b Button) SetAccessibilityLabel(label string) Button {
	b.setAccessibilityLabel(label)
	return b
}

// RemoveAccessibilityLabel removes the style from the button.
func (b Button) RemoveAccessibilityLabel() Button {
	b.optionals.AccessibilityLabel = false
	return b
}

type Checkboxes struct {
	slackType ElementType // required
	actionId  string      // required
	options   []Option    // required

	initialOptions     []Option // optional
	confirmationDialog ConfirmationDialog
	focusOnLoad        bool

	optional checkboxOptional
}

// checkboxOptional is a struct to keep track of which optional fields are set.
type checkboxOptional struct {
	InitialOptions     bool
	ConfirmationDialog bool
	FocusOnLoad        bool
}

// NewCheckboxes creates a new checkbox element.
func NewCheckboxes(actionId string) Checkboxes {
	return Checkboxes{
		slackType: CheckboxesElement,
		actionId:  actionId,
		options:   []Option{},
		optional: checkboxOptional{
			InitialOptions:     false,
			ConfirmationDialog: false,
			FocusOnLoad:        false,
		},
	}
}

// AddOption adds an option to the checkboxes element.
func (c *Checkboxes) addOption(option Option) {
	c.options = append(c.options, option)
}

func (c *Checkboxes) addInitialOption(option Option) {
	c.initialOptions = append(c.initialOptions, option)
	c.optional.InitialOptions = true
}

// AddConfirmationDialog adds a confirmation dialog to the checkboxes element.
func (c *Checkboxes) addConfirmationDialog(confirmationDialog ConfirmationDialog) {
	c.confirmationDialog = confirmationDialog
	c.optional.ConfirmationDialog = true
}

// RemoveConfirmationDialog removes the confirmation dialog from the checkboxes element.
func (c *Checkboxes) removeConfirmationDialog() {
	c.optional.ConfirmationDialog = false
}

// SetFocusOnLoad sets the focus on load flag for the checkboxes element.
func (c *Checkboxes) setFocusOnLoad(focusOnLoad bool) {
	c.focusOnLoad = focusOnLoad
	c.optional.FocusOnLoad = true
}

// RemoveFocusOnLoad removes the focus on load flag from the checkboxes element.
func (c *Checkboxes) removeFocusOnLoad() {
	c.optional.FocusOnLoad = false
}

type abstractCheckboxes struct {
	Type               string
	ActionID           string
	Options            []Option
	InitialOptions     []Option
	ConfirmationDialog ConfirmationDialog
	FocusOnLoad        bool

	Optional checkboxOptional
}

// create a new abstract checkboxes element
func (c Checkboxes) abstraction() abstractCheckboxes {
	return abstractCheckboxes{
		Type:               c.slackType.String(),
		ActionID:           c.actionId,
		Options:            c.options,
		InitialOptions:     c.initialOptions,
		ConfirmationDialog: c.confirmationDialog,
		FocusOnLoad:        c.focusOnLoad,
		Optional:           c.optional,
	}
}

func (c abstractCheckboxes) Template() string {
	return `{
		"type": "{{.Type}}",
		"action_id": "{{.ActionID}}",
		"options": [
			{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}
		]{{if .Optional.InitialOptions}},
		"initial_options": [{{range $index, $option := .InitialOptions}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]{{end}}{{if .Optional.ConfirmationDialog}},
		"confirmation_dialog": {{.ConfirmationDialog.Render}}{{end}}{{if .Optional.FocusOnLoad}},
		"focus_on_load": "{{.FocusOnLoad}}"{{end}}
	}`
}

// Render renders the checkboxes element to JSON.
func (c Checkboxes) Render() string {
	return Render(c.abstraction())
}

// AddOption add option to checkboxes
func (c Checkboxes) AddOption(option Option) Checkboxes {
	c.addOption(option)
	return c
}

// AddInitialOption add initial option to checkboxes
func (c Checkboxes) AddInitialOption(option Option) Checkboxes {
	c.addInitialOption(option)
	c.addOption(option)
	return c
}

// AddConfirmationDialog add confirmation dialog to checkboxes
func (c Checkboxes) AddConfirmationDialog(confirmationDialog ConfirmationDialog) Checkboxes {
	c.addConfirmationDialog(confirmationDialog)
	return c
}

// RemoveConfirmationDialog remove confirmation dialog from checkboxes
func (c Checkboxes) RemoveConfirmationDialog() Checkboxes {
	c.removeConfirmationDialog()
	return c
}

// FocusOnLoad set focus on load to checkboxes
func (c Checkboxes) FocusOnLoad() Checkboxes {
	c.setFocusOnLoad(true)
	return c
}

// DisableFocusOnLoad remove focus on load from checkboxes
func (c Checkboxes) DisableFocusOnLoad() Checkboxes {
	c.setFocusOnLoad(false)
	return c
}

//InputElement

type DatePicker struct {
	slackType ElementType
	actionId  string

	initialDate time.Time
	confirm     ConfirmationDialog
	placeholder CompositionText
	focus       bool

	options optionalDatePicker
}

type optionalDatePicker struct {
	InitialDate bool
	Confirm     bool
	Placeholder bool
	Focus       bool
}

func NewDatePicker(actionId string) DatePicker {
	return DatePicker{
		slackType: DatePickerElement,
		actionId:  actionId,
		options: optionalDatePicker{
			InitialDate: false,
			Confirm:     false,
			Placeholder: false,
			Focus:       false,
		},
	}
}

// SetInitialDate sets the initial date for the date picker.
func (d *DatePicker) setInitialDate(initialDate time.Time) {
	// todo: implement a parser for the date format YYYY-MM-DD
	d.initialDate = initialDate
	d.options.InitialDate = true
}

// removeInitialDate removes the initial date for the date picker.
func (d *DatePicker) removeInitialDate() {
	d.options.InitialDate = false
}

// SetConfirm sets the confirmation dialog for the date picker.
func (d *DatePicker) setConfirm(confirm ConfirmationDialog) {
	d.confirm = confirm
	d.options.Confirm = true
}

// removeConfirm removes the confirmation dialog for the date picker.
func (d *DatePicker) removeConfirm() {
	d.options.Confirm = false
}

// SetPlaceholder sets the placeholder for the date picker.
func (d *DatePicker) setPlaceholder(placeholder string) {
	d.placeholder = NewPlainText(placeholder)
	d.options.Placeholder = true
}

// removePlaceholder removes the placeholder for the date picker.
func (d *DatePicker) removePlaceholder() {
	d.options.Placeholder = false
}

// SetFocus sets the focus for the date picker.
func (d *DatePicker) setFocus(focus bool) {
	d.focus = focus
	d.options.Focus = true
}

// removeFocus removes the focus for the date picker.
func (d *DatePicker) removeFocus() {
	d.options.Focus = false
}

type abstractDatePicker struct {
	Type        string
	ActionId    string
	InitalDate  string
	Confirm     ConfirmationDialog
	Placeholder CompositionText
	Focus       bool
	Optionals   optionalDatePicker
}

// abstraction
func (d DatePicker) abstraction() abstractDatePicker {
	return abstractDatePicker{
		Type:     d.slackType.String(),
		ActionId: d.actionId,

		InitalDate:  d.initialDate.Format("2006-01-02"),
		Confirm:     d.confirm,
		Placeholder: d.placeholder,
		Focus:       d.focus,

		Optionals: d.options,
	}
}

// Template
func (d abstractDatePicker) Template() string {
	return `{
		"type": "{{.Type}}",
		"action_id": "{{.ActionId}}"{{if .Optionals.InitialDate}},
		"initial_date": "{{.InitalDate}}"{{end}}{{if .Optionals.Confirm}}, 
		"confirm": {{.Confirm.Render}}{{end}}{{if .Optionals.Placeholder}},
		"placeholder": {{.Placeholder.Render}}{{end}}{{if .Optionals.Focus}},
		"initial_focus": {{.Focus}}{{end}}
	}`
}

// Render
func (d DatePicker) Render() string {
	return Render(d.abstraction())
}

// AddInitialDate chain function to add initial date to an existing date picker
func (d DatePicker) AddInitialDate(initialDate time.Time) DatePicker {
	d.setInitialDate(initialDate)
	return d
}

// RemoveInitialDate remove add initial date from date picker
func (d DatePicker) RemoveInitialDate() DatePicker {
	d.removeInitialDate()
	return d
}

// AddConfirm chain function to add confirm to an existing date picker
func (d DatePicker) AddConfirm(confirm ConfirmationDialog) DatePicker {
	d.setConfirm(confirm)
	return d
}

// RemoveConfirm remove add confirm from date picker
func (d DatePicker) RemoveConfirm() DatePicker {
	d.removeConfirm()
	return d
}

// AddPlaceholder chain function to add placeholder to an existing date picker
func (d DatePicker) AddPlaceholder(placeholder string) DatePicker {
	d.setPlaceholder(placeholder)
	return d
}

// RemovePlaceholder remove add placeholder from date picker
func (d DatePicker) RemovePlaceholder() DatePicker {
	d.removePlaceholder()
	return d
}

// MakeFocused chain function to add focus to an existing date picker
func (d DatePicker) MakeFocused() DatePicker {
	d.setFocus(true)
	return d
}

// RemoveInitialFocus remove add focus from date picker
func (d DatePicker) RemoveInitialFocus() DatePicker {
	d.setFocus(false)
	return d
}

type DateTimePicker struct {
	slackType ElementType
	actionId  string

	initialDateTime time.Time
	confirm         ConfirmationDialog
	focusOnLoad     bool

	options dateTimePickerOptions
}

type dateTimePickerOptions struct {
	InitialDateTime bool
	Confirm         bool
	FocusOnLoad     bool
}

// NewDateTimePicker creates a new date picker.
func NewDateTimePicker(actionId string) DateTimePicker {
	return DateTimePicker{
		slackType: DateTimePickerElement,
		actionId:  actionId,
		options: dateTimePickerOptions{
			InitialDateTime: false,
			Confirm:         false,
			FocusOnLoad:     false,
		},
	}
}

// SetInitialDateTime sets the initial date for the date picker.
func (d *DateTimePicker) setInitialDateTime(initialDateTime time.Time) {
	d.initialDateTime = initialDateTime
	d.options.InitialDateTime = true
}

func (d *DateTimePicker) removeInitialDateTime() {
	d.options.InitialDateTime = false
}

// SetConfirm sets the confirmation dialog for the date picker.
func (d *DateTimePicker) setConfirm(confirm ConfirmationDialog) {
	d.confirm = confirm
	d.options.Confirm = true
}

func (d *DateTimePicker) removeConfirm() {
	d.options.Confirm = false
}

// SetFocusOnLoad sets the focus on load for the date picker.
func (d *DateTimePicker) setFocusOnLoad(focusOnLoad bool) {
	d.focusOnLoad = focusOnLoad
	d.options.FocusOnLoad = true
}

func (d *DateTimePicker) removeFocusOnLoad() {
	d.options.FocusOnLoad = false
}

// AddInitialDateTime chain function to add initial date to an existing date picker
func (d DateTimePicker) AddInitialDateTime(initialDateTime time.Time) DateTimePicker {
	d.setInitialDateTime(initialDateTime)
	return d
}

// RemoveInitialDateTime remove add initial date from date picker
func (d DateTimePicker) RemoveInitialDateTime() DateTimePicker {
	d.removeInitialDateTime()
	return d
}

// AddConfirm chain function to add confirmation dialog to an existing date picker
func (d DateTimePicker) AddConfirmationDialog(confirm ConfirmationDialog) DateTimePicker {
	d.setConfirm(confirm)
	return d
}

// RemoveConfirm remove add confirmation dialog from date picker
func (d DateTimePicker) RemoveConfirmationDialog() DateTimePicker {
	d.removeConfirm()
	return d
}

// AddFocusOnLoad chain function to add focus on load to an existing date picker
func (d DateTimePicker) AddFocusOnLoad(focusOnLoad bool) DateTimePicker {
	d.setFocusOnLoad(focusOnLoad)
	return d
}

// RemoveFocusOnLoad remove add focus on load from date picker
func (d DateTimePicker) RemoveFocusOnLoad() DateTimePicker {
	d.removeFocusOnLoad()
	return d
}

// abstraction type
type abstractDateTimePicker struct {
	Type     string
	ActionId string

	InitialDateTime string
	Confirm         ConfirmationDialog
	FocusOnLoad     bool

	Optionals dateTimePickerOptions
}

// abstraction method
func (d DateTimePicker) abstraction() abstractDateTimePicker {
	unixString := strconv.FormatInt(d.initialDateTime.Unix(), 10)

	return abstractDateTimePicker{
		Type:     d.slackType.String(),
		ActionId: d.actionId,

		InitialDateTime: unixString,
		Confirm:         d.confirm,
		FocusOnLoad:     d.focusOnLoad,

		Optionals: d.options,
	}
}

// Render renders the date picker to a JSON string.
func (d DateTimePicker) Render() string {
	return Render(d.abstraction())
}

// Template function
func (d abstractDateTimePicker) Template() string {
	return `{
		"type": "{{.Type}}",
		"action_id": "{{.ActionId}}"{{if .Optionals.InitialDateTime}},
		"initial_date": "{{.InitialDateTime}}"{{end}}{{if .Optionals.Confirm}},
		"confirm": {{.Confirm.Render}}{{end}}{{if .Optionals.FocusOnLoad}},
		"focus_on_load": {{.FocusOnLoad}},{{end}}
	}`
}

type EmailInput struct {
	slackType ElementType
	actionId  string

	initialEmail         string
	dispatchActionConfig DispatchActionConfig
	focusOnLoad          bool
	placeholder          string

	options EmailInputOptions
}

type EmailInputOptions struct {
	InitialEmail         bool
	DispatchActionConfig bool
	FocusOnLoad          bool
	Placeholder          bool
}

func NewEmailInput(actionId string) EmailInput {
	return EmailInput{
		slackType:            EmailInputElement,
		actionId:             actionId,
		dispatchActionConfig: NewDispatchActionConfig(),

		options: EmailInputOptions{
			InitialEmail:         false,
			DispatchActionConfig: false,
			FocusOnLoad:          false,
			Placeholder:          false,
		},
	}
}

// SetInitialEmail sets the initial email for the email input.
func (e *EmailInput) setInitialEmail(initialEmail string) {
	e.initialEmail = initialEmail
	e.options.InitialEmail = true
}

// RemoveInitialEmail removes the initial email for the email input.
func (e *EmailInput) removeInitialEmail() {
	e.options.InitialEmail = false
}

// todo: email input not implemented yet

type ImageElement struct {
	slackType ElementType
	imageUrl  *url.URL
	altText   string
}

func NewImageElement(imageUrl *url.URL, altText string) ImageElement {
	return ImageElement{
		slackType: ImageElementType,
		imageUrl:  imageUrl,
		altText:   altText,
	}
}

// SetImageUrl sets the image url for the image.
func (i *ImageElement) setImageUrl(imageUrl *url.URL) {
	i.imageUrl = imageUrl
}

// SetAltText sets the alt text for the image.
func (i *ImageElement) setAltText(altText string) {
	i.altText = altText
}

// abstractImage struct
type abstractImage struct {
	Type     string
	ImageUrl string
	AltText  string
}

// abstraction method
func (i *ImageElement) abstraction() abstractImage {
	return abstractImage{
		Type:     i.slackType.String(),
		ImageUrl: i.imageUrl.String(),
		AltText:  i.altText,
	}
}

// Template method
func (i abstractImage) Template() string {
	return `{
		"type": "{{.Type}}",
		"image_url": "{{.ImageUrl}}",
		"alt_text": "{{.AltText}}"
	}`
}

// Render method
func (i ImageElement) Render() string {
	return Render(i.abstraction())
}

// InputElement

type MultiSelectMenuWithConversationsList struct {
	slackType ElementType
	actionID  string

	confirm          ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      CompositionText

	// Conversation
	defaultToCurrentConversation bool
	initialConversations         []string
	filter                       Filter

	optionals multiSelectMenuWithConversationsListOptions
}

type multiSelectMenuWithConversationsListOptions struct {
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversations         bool
	Filter                       bool
}

// abstracted type
type multiSelectMenuWithConversationsListAbstraction struct {
	Type     string
	ActionId string

	Confirm          ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      CompositionText

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversations         []string
	Filter                       Filter

	Optionals multiSelectMenuWithConversationsListOptions
}

func NewMultiSelectMenuWithConversationsList(actionId string) MultiSelectMenuWithConversationsList {
	return MultiSelectMenuWithConversationsList{
		slackType: MultiSelectMenuWithConversationsListElement,
		actionID:  actionId,
		optionals: multiSelectMenuWithConversationsListOptions{
			Confirm:                      false,
			MaxSelectedItems:             false,
			FocusOnLoad:                  false,
			Placeholder:                  false,
			DefaultToCurrentConversation: false,
			InitialConversations:         false,
			Filter:                       false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenuWithConversationsList) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithConversationsList) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenuWithConversationsList) UpdateActionId(actionId string) MultiSelectMenuWithConversationsList {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithConversationsList) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithConversationsList) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithConversationsList) AddConfirmDialog(confirm ConfirmationDialog) MultiSelectMenuWithConversationsList {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m MultiSelectMenuWithConversationsList) RemoveConfirmDialog() MultiSelectMenuWithConversationsList {
	m.removeConfirm()
	return m
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenuWithConversationsList) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithConversationsList) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithConversationsList) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithConversationsList {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithConversationsList) UnsetMaxSelectedItems() MultiSelectMenuWithConversationsList {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenuWithConversationsList) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithConversationsList) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithConversationsList) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithConversationsList {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithConversationsList) UnsetFocusOnLoad() MultiSelectMenuWithConversationsList {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenuWithConversationsList) setPlaceholder(placeholder string) {
	m.placeholder = NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithConversationsList) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithConversationsList) AddPlaceholder(placeholder string) MultiSelectMenuWithConversationsList {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithConversationsList) RemovePlaceholder() MultiSelectMenuWithConversationsList {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// defaultToCurrentConversation

// setDefaultToCurrentConversation public set default to current conversation
func (m *MultiSelectMenuWithConversationsList) setDefaultToCurrentConversation(defaultToCurrentConversation bool) {
	m.defaultToCurrentConversation = defaultToCurrentConversation
	m.optionals.DefaultToCurrentConversation = defaultToCurrentConversation

}

// unsetDefaultToCurrentConversation public remove default to current conversation
func (m *MultiSelectMenuWithConversationsList) unsetDefaultToCurrentConversation() {
	m.setDefaultToCurrentConversation(false)
}

// DefaultToCurrentConversation public set default to current conversation
func (m MultiSelectMenuWithConversationsList) DefaultToCurrentConversation() MultiSelectMenuWithConversationsList {
	m.setDefaultToCurrentConversation(true)
	return m
}

// UnsetDefaultToCurrentConversation public remove default to current conversation
func (m MultiSelectMenuWithConversationsList) UnsetDefaultToCurrentConversation() MultiSelectMenuWithConversationsList {
	m.unsetDefaultToCurrentConversation()
	return m
}

//////////////////////////////////////////////////
// initialConversations

// addInitialConversation private add initial conversation
func (m *MultiSelectMenuWithConversationsList) addInitialConversation(initialConversation string) {
	m.initialConversations = append(m.initialConversations, initialConversation)
	m.optionals.InitialConversations = true
}

// removeInitialConversations private remove initial conversations
func (m *MultiSelectMenuWithConversationsList) removeInitialConversations() {
	m.initialConversations = []string{}
	m.optionals.InitialConversations = false
}

// AddInitialConversation public add initial conversation
func (m MultiSelectMenuWithConversationsList) AddInitialConversation(initialConversation string) MultiSelectMenuWithConversationsList {
	m.addInitialConversation(initialConversation)
	return m
}

// ClearInitialConversations clear initial conversations
func (m MultiSelectMenuWithConversationsList) ClearInitialConversations() MultiSelectMenuWithConversationsList {
	m.removeInitialConversations()
	return m
}

//////////////////////////////////////////////////
// filter

func (m *MultiSelectMenuWithConversationsList) setFilter(filter Filter) {
	m.filter = filter
	m.optionals.Filter = true
}

func (m *MultiSelectMenuWithConversationsList) removeFilter() {
	m.optionals.Filter = false
}

// AddFilter public set filter
func (m MultiSelectMenuWithConversationsList) AddFilter(filter Filter) MultiSelectMenuWithConversationsList {
	m.setFilter(filter)
	return m
}

// RemoveFilter public remove filter
func (m MultiSelectMenuWithConversationsList) RemoveFilter() MultiSelectMenuWithConversationsList {
	m.removeFilter()
	return m
}

// create abstract
func (m MultiSelectMenuWithConversationsList) abstraction() multiSelectMenuWithConversationsListAbstraction {
	return multiSelectMenuWithConversationsListAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,

		// Conversation
		DefaultToCurrentConversation: m.defaultToCurrentConversation,
		InitialConversations:         removeDuplicateString(m.initialConversations),
		Filter:                       m.filter,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m multiSelectMenuWithConversationsListAbstraction) Template() string {
	return `{
	"action_id": "{{ .ActionId }}",

	"type": "{{ .Type }}"

{{if .Optionals.InitialConversations}},
	"initial_conversations": [{{range $index, $conversations := .InitialConversations}}{{if $index}},{{end}}"{{ $conversations}}"{{end}}]
{{end}}

{{if .Optionals.DefaultToCurrentConversation}},
	"default_to_current_conversation": {{ .DefaultToCurrentConversation }}
{{end}}

{{if .Optionals.Confirm }},
	"confirm": {{ .Confirm.Render }}
{{end}}

{{if .Optionals.MaxSelectedItems }},
	"max_selected_items": {{ .MaxSelectedItems }}
{{end}}

{{if .Optionals.Filter }},
	{{ .Filter.Render }}
{{end}}

{{if .Optionals.FocusOnLoad }},
	"focus_on_load": {{ .FocusOnLoad }}
{{end}}

{{if .Optionals.Placeholder }},
	"placeholder": {{ .Placeholder.Render }}
{{end}}

}`
}

func (m MultiSelectMenuWithConversationsList) ElementRender() {}

func (m MultiSelectMenuWithConversationsList) Render() string {
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m MultiSelectMenuWithConversationsList) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

// InputElement

type MultiSelectMenuWithExternalDataSource struct {
	slackType ElementType
	actionID  string

	initialOptions   []Option
	confirm          ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      CompositionText

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
		slackType:      MultiSelectMenuWithExternalDataSourceElement,
		actionID:       actionId,
		initialOptions: []Option{},
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

func (m *MultiSelectMenuWithExternalDataSource) addInitialOption(initialOption Option) {
	m.initialOptions = append(m.initialOptions, initialOption)
	m.optionals.InitialOptions = true
}

func (m *MultiSelectMenuWithExternalDataSource) removeInitialOptions() {
	m.optionals.InitialOptions = false
}

func (m *MultiSelectMenuWithExternalDataSource) setInitialOptions(initialOptions []Option) {
	m.initialOptions = initialOptions
	m.optionals.InitialOptions = true
}

// ClearInitialOptions clear initial options
func (m MultiSelectMenuWithExternalDataSource) ClearInitialOptions() MultiSelectMenuWithExternalDataSource {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m MultiSelectMenuWithExternalDataSource) AddInitialOption(initialOption Option) MultiSelectMenuWithExternalDataSource {
	m.addInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithExternalDataSource) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithExternalDataSource) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithExternalDataSource) AddConfirmDialog(confirm ConfirmationDialog) MultiSelectMenuWithExternalDataSource {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m MultiSelectMenuWithExternalDataSource) RemoveConfirmDialog() MultiSelectMenuWithExternalDataSource {
	m.removeConfirm()
	return m
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
	m.placeholder = NewPlainText(placeholder)
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

	InitialOptions   []Option
	Confirm          ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      CompositionText

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
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m MultiSelectMenuWithExternalDataSource) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

// InputElement

type MultiSelectMenuWithPublicChannelsSelect struct {
	slackType ElementType
	actionID  string

	confirm          ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      CompositionText

	// Public Channel
	initialChannels []string

	optionals multiSelectMenuWithPublicChannelsSelectOptions
}

type multiSelectMenuWithPublicChannelsSelectOptions struct {
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool

	// Public Channel
	InitialChannels bool
}

// abstracted type
type multiSelectMenuWithPublicChannelsSelectAbstraction struct {
	Type     string
	ActionId string

	Confirm          ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      CompositionText

	// Public Channel
	InitialChannels []string

	Optionals multiSelectMenuWithPublicChannelsSelectOptions
}

func NewMultiSelectMenuWithPublicChannelsSelect(actionId string) MultiSelectMenuWithPublicChannelsSelect {
	return MultiSelectMenuWithPublicChannelsSelect{
		slackType: MultiSelectMenuWithPublicChannelsSelectElement,
		actionID:  actionId,

		optionals: multiSelectMenuWithPublicChannelsSelectOptions{
			Confirm:          false,
			MaxSelectedItems: false,
			FocusOnLoad:      false,
			Placeholder:      false,
			InitialChannels:  false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenuWithPublicChannelsSelect) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenuWithPublicChannelsSelect) UpdateActionId(actionId string) MultiSelectMenuWithPublicChannelsSelect {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithPublicChannelsSelect) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithPublicChannelsSelect) AddConfirmDialog(confirm ConfirmationDialog) MultiSelectMenuWithPublicChannelsSelect {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m MultiSelectMenuWithPublicChannelsSelect) RemoveConfirmDialog() MultiSelectMenuWithPublicChannelsSelect {
	m.removeConfirm()
	return m
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenuWithPublicChannelsSelect) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithPublicChannelsSelect) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithPublicChannelsSelect {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithPublicChannelsSelect) UnsetMaxSelectedItems() MultiSelectMenuWithPublicChannelsSelect {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenuWithPublicChannelsSelect) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithPublicChannelsSelect) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithPublicChannelsSelect {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithPublicChannelsSelect) UnsetFocusOnLoad() MultiSelectMenuWithPublicChannelsSelect {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenuWithPublicChannelsSelect) setPlaceholder(placeholder string) {
	m.placeholder = NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithPublicChannelsSelect) AddPlaceholder(placeholder string) MultiSelectMenuWithPublicChannelsSelect {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithPublicChannelsSelect) RemovePlaceholder() MultiSelectMenuWithPublicChannelsSelect {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// initialChannels

func (m *MultiSelectMenuWithPublicChannelsSelect) setInitialChannels(initialChannels []string) {
	m.initialChannels = initialChannels
	m.optionals.InitialChannels = true
}

func (m *MultiSelectMenuWithPublicChannelsSelect) removeInitialChannels() {
	m.initialChannels = []string{}
	m.optionals.InitialChannels = false
}

// addInitialChannel private set initial channels
func (m *MultiSelectMenuWithPublicChannelsSelect) addInitialChannel(initialChannel string) {
	m.initialChannels = append(m.initialChannels, initialChannel)
	m.optionals.InitialChannels = true
}

// removeInitialChannel private remove initial channels
func (m *MultiSelectMenuWithPublicChannelsSelect) removeInitialChannel(initialChannel string) {
	for i, v := range m.initialChannels {
		if v == initialChannel {
			m.initialChannels = append(m.initialChannels[:i], m.initialChannels[i+1:]...)
		}
	}
}

// AddInitialChannels public set initial channels
func (m MultiSelectMenuWithPublicChannelsSelect) AddInitialChannels(initialChannel string) MultiSelectMenuWithPublicChannelsSelect {
	m.addInitialChannel(initialChannel)
	return m
}

// RemoveInitialChannels public remove initial channels
func (m MultiSelectMenuWithPublicChannelsSelect) RemoveInitialChannels(initialChannel string) MultiSelectMenuWithPublicChannelsSelect {
	m.removeInitialChannel(initialChannel)
	return m
}

// create abstract
func (m MultiSelectMenuWithPublicChannelsSelect) abstraction() multiSelectMenuWithPublicChannelsSelectAbstraction {
	return multiSelectMenuWithPublicChannelsSelectAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,

		// Public Channel
		InitialChannels: removeDuplicateString(m.initialChannels),

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m multiSelectMenuWithPublicChannelsSelectAbstraction) Template() string {
	return `{
	"action_id": "{{ .ActionId }}",
	"type": "{{ .Type }}"

{{if .Optionals.InitialChannels}},
	"initial_channels": [{{range $index, $channel := .InitialChannels}}{{if $index}},{{end}}"{{ $channel}}"{{end}}]
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

}`
}

func (m MultiSelectMenuWithPublicChannelsSelect) ElementRender() {}

func (m MultiSelectMenuWithPublicChannelsSelect) Render() string {
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m MultiSelectMenuWithPublicChannelsSelect) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

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
func (m MultiSelectMenuWithStaticOption) RemoveConfirmDialog() MultiSelectMenuWithStaticOption {
	m.removeConfirm()
	return m
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

// InputElement

type MultiSelectMenuWithUserList struct {
	slackType ElementType
	actionID  string

	confirm          ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      CompositionText

	// User List
	initialUsers []string

	optionals multiSelectMenuWithUserListOptions
}

type multiSelectMenuWithUserListOptions struct {
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool

	// User List
	InitialUsers bool
}

func NewMultiSelectMenuWithUserList(actionId string) MultiSelectMenuWithUserList {
	return MultiSelectMenuWithUserList{
		slackType: MultiSelectMenuWithUserListElement,
		actionID:  actionId,
		optionals: multiSelectMenuWithUserListOptions{
			Confirm:          false,
			MaxSelectedItems: false,
			FocusOnLoad:      false,
			Placeholder:      false,
			InitialUsers:     false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *MultiSelectMenuWithUserList) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *MultiSelectMenuWithUserList) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m MultiSelectMenuWithUserList) UpdateActionId(actionId string) MultiSelectMenuWithUserList {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *MultiSelectMenuWithUserList) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *MultiSelectMenuWithUserList) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m MultiSelectMenuWithUserList) AddConfirmDialog(confirm ConfirmationDialog) MultiSelectMenuWithUserList {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m MultiSelectMenuWithUserList) RemoveConfirmDialog() MultiSelectMenuWithUserList {
	m.removeConfirm()
	return m
}

//////////////////////////////////////////////////
// maxSelectedItems

func (m *MultiSelectMenuWithUserList) setMaxSelectedItems(maxSelectedItems int) {
	m.maxSelectedItems = maxSelectedItems
	m.optionals.MaxSelectedItems = true
}

func (m *MultiSelectMenuWithUserList) removeMaxSelectedItems() {
	m.optionals.MaxSelectedItems = false
}

// MaxSelectedItems public set max selected items
func (m MultiSelectMenuWithUserList) MaxSelectedItems(maxSelectedItems int) MultiSelectMenuWithUserList {
	m.setMaxSelectedItems(maxSelectedItems)
	m.optionals.MaxSelectedItems = true
	return m
}

// UnsetMaxSelectedItems public remove max selected items
func (m MultiSelectMenuWithUserList) UnsetMaxSelectedItems() MultiSelectMenuWithUserList {
	m.optionals.MaxSelectedItems = false
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *MultiSelectMenuWithUserList) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *MultiSelectMenuWithUserList) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// FocusOnLoad public set focus on load
func (m MultiSelectMenuWithUserList) FocusOnLoad(focusOnLoad bool) MultiSelectMenuWithUserList {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m MultiSelectMenuWithUserList) UnsetFocusOnLoad() MultiSelectMenuWithUserList {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *MultiSelectMenuWithUserList) setPlaceholder(placeholder string) {
	m.placeholder = NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *MultiSelectMenuWithUserList) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m MultiSelectMenuWithUserList) AddPlaceholder(placeholder string) MultiSelectMenuWithUserList {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m MultiSelectMenuWithUserList) RemovePlaceholder() MultiSelectMenuWithUserList {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// initialUsers

func (m *MultiSelectMenuWithUserList) addInitialUser(initialUser string) {
	m.initialUsers = append(m.initialUsers, initialUser)
	m.optionals.InitialUsers = true
}

func (m *MultiSelectMenuWithUserList) removeInitialUsers() {
	m.initialUsers = []string{}
	m.optionals.InitialUsers = false
}

// AddInitialUser public add initial user
func (m MultiSelectMenuWithUserList) AddInitialUser(initialUser string) MultiSelectMenuWithUserList {
	m.addInitialUser(initialUser)
	return m
}

// ClearInitialUsers clear initial users
func (m MultiSelectMenuWithUserList) ClearInitialUsers() MultiSelectMenuWithUserList {
	m.removeInitialUsers()
	return m
}

// ////////////////////////////////////////////////
// abstract

// abstracted type
type multiSelectMenuWithUserListAbstraction struct {
	Type     string
	ActionId string

	Confirm          ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      CompositionText

	// User List
	InitialUsers []string

	Optionals multiSelectMenuWithUserListOptions
}

// create abstract
func (m MultiSelectMenuWithUserList) abstraction() multiSelectMenuWithUserListAbstraction {
	return multiSelectMenuWithUserListAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm:          m.confirm,
		MaxSelectedItems: m.maxSelectedItems,
		FocusOnLoad:      m.focusOnLoad,
		Placeholder:      m.placeholder,

		// User List
		InitialUsers: removeDuplicateString(m.initialUsers),

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m multiSelectMenuWithUserListAbstraction) Template() string {
	return `{
"action_id": "{{ .ActionId }}",
		
"type": "{{ .Type }}"	

{{if .Optionals.InitialUsers}},
	"initial_users": [{{range $index, $user := .InitialUsers}}{{if $index}},{{end}}"{{ $user}}"{{end}}]
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

}`
}

func (m MultiSelectMenuWithUserList) ElementRender() {}

func (m MultiSelectMenuWithUserList) Render() string {
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m MultiSelectMenuWithUserList) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

type NumberInput struct {
	slackType        ElementType
	actionID         string
	isDecimalAllowed bool

	initialValue int
	minValue     int
	maxValue     int

	dispatchActionConfig DispatchActionConfig

	focusOnLoad bool
	placeholder CompositionText

	optionals numberInputOptions
}

// optionals struct
type numberInputOptions struct {
	MinValue             bool
	MaxValue             bool
	InitialValue         bool
	DispatchActionConfig bool
	FocusOnLoad          bool
	Placeholder          bool
}

func NewNumberInput(actionId string) NumberInput {
	return NumberInput{
		slackType:        NumberInputElement,
		actionID:         actionId,
		isDecimalAllowed: false,

		optionals: numberInputOptions{
			MinValue:             false,
			MaxValue:             false,
			InitialValue:         false,
			DispatchActionConfig: false,
			FocusOnLoad:          false,
			Placeholder:          false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (n *NumberInput) setActionId(actionId string) {
	n.actionID = actionId
}

func (n *NumberInput) removeActionId() {
	n.actionID = ""
}

// UpdateActionId public update action id
func (n NumberInput) UpdateActionId(actionId string) NumberInput {
	n.setActionId(actionId)
	return n
}

//////////////////////////////////////////////////
// isDecimalAllowed

func (n *NumberInput) setIsDecimalAllowed(isDecimalAllowed bool) {
	n.isDecimalAllowed = isDecimalAllowed
}

func (n *NumberInput) removeIsDecimalAllowed() {
	n.isDecimalAllowed = false
}

// UpdateIsDecimalAllowed public update isDecimalAllowed
func (n NumberInput) DecimalAllowed() NumberInput {
	n.setIsDecimalAllowed(true)
	return n
}

func (n NumberInput) UnsetDecimalAllowed() NumberInput {
	n.setIsDecimalAllowed(false)
	return n
}

//////////////////////////////////////////////////
// initialValue

func (n *NumberInput) setInitialValue(initialValue int) {
	n.initialValue = initialValue
	n.optionals.InitialValue = true
}

func (n *NumberInput) removeInitialValue() {
	n.optionals.InitialValue = false
}

// InitialValue public update initialValue
func (n NumberInput) InitialValue(initialValue int) NumberInput {
	n.setInitialValue(initialValue)
	return n
}

//////////////////////////////////////////////////
// minValue

func (n *NumberInput) setMinValue(minValue int) {
	n.minValue = minValue
	n.optionals.MinValue = true
}

func (n *NumberInput) removeMinValue() {
	n.optionals.MinValue = false
}

// UpdateMinValue public update minValue
func (n NumberInput) MinValue(minValue int) NumberInput {
	n.setMinValue(minValue)
	return n
}

//////////////////////////////////////////////////
// maxValue

func (n *NumberInput) setMaxValue(maxValue int) {
	n.maxValue = maxValue
	n.optionals.MaxValue = true
}

func (n *NumberInput) removeMaxValue() {
	n.optionals.MaxValue = false
}

// UpdateMaxValue public update maxValue
func (n NumberInput) MaxValue(maxValue int) NumberInput {
	n.setMaxValue(maxValue)
	return n
}

//////////////////////////////////////////////////
// dispatchActionConfig

func (n *NumberInput) setDispatchActionConfig(dispatchActionConfig DispatchActionConfig) {
	n.dispatchActionConfig = dispatchActionConfig
	n.optionals.DispatchActionConfig = true
}

func (n *NumberInput) removeDispatchActionConfig() {
	n.optionals.DispatchActionConfig = false
}

// DispatchAction public update dispatchActionConfig
func (n NumberInput) DispatchAction(dispatchActionConfig DispatchActionConfig) NumberInput {
	n.setDispatchActionConfig(dispatchActionConfig)
	return n
}

//////////////////////////////////////////////////
// focusOnLoad

func (n *NumberInput) setFocusOnLoad(focusOnLoad bool) {
	n.focusOnLoad = focusOnLoad
	n.optionals.FocusOnLoad = true
}

func (n *NumberInput) removeFocusOnLoad() {
	n.optionals.FocusOnLoad = false
}

// FocusOnLoad public update focusOnLoad
func (n NumberInput) FocusOnLoad() NumberInput {
	n.setFocusOnLoad(true)
	return n
}

func (n NumberInput) UnsetFocusOnLoad() NumberInput {
	n.setFocusOnLoad(false)
	return n
}

//////////////////////////////////////////////////
// placeholder

func (n *NumberInput) setPlaceholder(placeholder string) {
	n.placeholder = NewPlainText(placeholder)
	n.optionals.Placeholder = true
}

func (n *NumberInput) removePlaceholder() {
	n.optionals.Placeholder = false
}

// Placeholder public update placeholder
func (n NumberInput) Placeholder(placeholder string) NumberInput {
	n.setPlaceholder(placeholder)
	return n
}

//////////////////////////////////////////////////
// abstraction

type numberInputAbstraction struct {
	Type                 string
	ActionID             string
	IsDecimalAllowed     bool
	InitialValue         int
	MinValue             int
	MaxValue             int
	DispatchActionConfig DispatchActionConfig
	FocusOnLoad          bool
	Placeholder          CompositionText

	Optionals numberInputOptions
}

func (n NumberInput) abstraction() numberInputAbstraction {
	return numberInputAbstraction{
		Type:                 n.slackType.String(),
		ActionID:             n.actionID,
		IsDecimalAllowed:     n.isDecimalAllowed,
		InitialValue:         n.initialValue,
		MinValue:             n.minValue,
		MaxValue:             n.maxValue,
		DispatchActionConfig: n.dispatchActionConfig,
		FocusOnLoad:          n.focusOnLoad,
		Placeholder:          n.placeholder,

		Optionals: n.optionals,
	}
}

//////////////////////////////////////////////////
// template

func (n numberInputAbstraction) Template() string {
	return `{
"type": "{{.Type}}",
"action_id": "{{.ActionID}}",
"is_decimal_allowed": {{.IsDecimalAllowed}}

{{if .Optionals.InitialValue}},
	"initial_value": "{{.InitialValue}}"
{{end}}

{{if .Optionals.MinValue}},
	"min_value": "{{.MinValue}}"
{{end}}

{{if .Optionals.MaxValue}},
	"max_value": "{{.MaxValue}}"
{{end}}

{{if .Optionals.DispatchActionConfig}},
	"dispatch_action_config": {{.DispatchActionConfig.Render}}
{{end}}

{{if .Optionals.FocusOnLoad}},
	"focus_on_load": {{.FocusOnLoad}}
{{end}}

{{if .Optionals.Placeholder}},
	"placeholder": {{.Placeholder.Render}}
{{end}}
}`
}

// render public render
func (n NumberInput) Render() string {
	raw := Render(n.abstraction())
	return Pretty(raw)
}

// element interface
func (n NumberInput) InputElement() {}

// Input
func (n NumberInput) Input(label string) Input {
	return NewInput(label, n)
}

type OverflowMenu struct {
	slackType ElementType
	actionID  string

	options []Option
	confirm ConfirmationDialog

	optionals overflowMenuOptions
}

type overflowMenuOptions struct {
	Confirm bool
}

// NewOverflowMenu creates a new OverflowMenu
func NewOverflowMenu(actionId string) OverflowMenu {
	return OverflowMenu{
		slackType: OverflowMenuElement,
		actionID:  actionId,
		options:   []Option{},
		optionals: overflowMenuOptions{
			Confirm: false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (o *OverflowMenu) setActionId(actionId string) {
	o.actionID = actionId
}

func (o *OverflowMenu) removeActionId() {
	o.actionID = ""
}

// UpdateActionId public update action id
func (o OverflowMenu) UpdateActionId(actionId string) OverflowMenu {
	o.setActionId(actionId)
	return o
}

//////////////////////////////////////////////////
// options

// AddOption adds an option to the OverflowMenu
func (o *OverflowMenu) addOption(option Option) {
	o.options = append(o.options, option)
}

// RemoveOption removes an option from the OverflowMenu
func (o *OverflowMenu) removeOption(option Option) {
	for i, v := range o.options {
		if v == option {
			o.options = append(o.options[:i], o.options[i+1:]...)
		}
	}
}

// AddOption public update options
func (o OverflowMenu) AddOption(options Option) OverflowMenu {
	o.addOption(options)
	return o
}

//////////////////////////////////////////////////
// confirm

func (o *OverflowMenu) setConfirm(confirm ConfirmationDialog) {
	o.confirm = confirm
	o.optionals.Confirm = true
}

func (o *OverflowMenu) removeConfirm() {
	o.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (o OverflowMenu) AddConfirmDialog(confirm ConfirmationDialog) OverflowMenu {
	o.setConfirm(confirm)
	o.optionals.Confirm = true
	return o
}

// RemoveConfirmDialog public remove confirm
func (o OverflowMenu) RemoveConfirmDialog() OverflowMenu {
	o.removeConfirm()
	return o
}

//////////////////////////////////////////////////
// abstraction

type overflowMenuAbstraction struct {
	Type     string
	ActionID string
	Options  []Option
	Confirm  ConfirmationDialog

	Optionals overflowMenuOptions
}

// abstractOverflowMenu abstracts the OverflowMenu
func (o OverflowMenu) abstractOverflowMenu() overflowMenuAbstraction {
	return overflowMenuAbstraction{
		Type:     o.slackType.String(),
		ActionID: o.actionID,
		Options:  o.options,
		Confirm:  o.confirm,

		Optionals: o.optionals,
	}
}

// Template returns the template for the OverflowMenu
func (o overflowMenuAbstraction) Template() string {
	return `{
"type": "{{.Type}}",
"action_id": "{{.ActionID}}",
"options": [{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]

{{if .Optionals.Confirm}},
	"confirm": {{.Confirm.Render}}
{{end}}

}`
}

// Render
func (o OverflowMenu) Render() string {
	raw := Render(o.abstractOverflowMenu())
	return Pretty(raw)
}

//Plain TextInput
// https://api.slack.com/reference/block-kit/block-elements#input

type PlainTextInput struct {
	slackType ElementType
	actionID  string

	initialValue         string
	multiline            bool
	minLength            int
	maxLength            int
	dispatchActionConfig DispatchActionConfig
	focusOnLoad          bool
	placeholder          CompositionText

	optionals plainTextInputOptions
}

// optionals
type plainTextInputOptions struct {
	InitialValue         bool
	Multiline            bool
	MinLength            bool
	MaxLength            bool
	DispatchActionConfig bool
	FocusOnLoad          bool
	Placeholder          bool
}

//////////////////////////////////////////////////

func NewPlainTextInput(actionId string) PlainTextInput {
	return PlainTextInput{
		slackType: PlainTextInputElement,
		actionID:  actionId,
		optionals: plainTextInputOptions{
			InitialValue:         false,
			Multiline:            false,
			MinLength:            false,
			MaxLength:            false,
			DispatchActionConfig: false,
			FocusOnLoad:          false,
			Placeholder:          false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *PlainTextInput) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *PlainTextInput) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m PlainTextInput) UpdateActionId(actionId string) PlainTextInput {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// initialValue

func (m *PlainTextInput) setInitialValue(initialValue string) {
	m.initialValue = initialValue
	m.optionals.InitialValue = true
}

func (m *PlainTextInput) removeInitialValue() {
	m.initialValue = ""
	m.optionals.InitialValue = false
}

// UpdateInitialValue public update initial value
func (m PlainTextInput) UpdateInitialValue(initialValue string) PlainTextInput {
	m.setInitialValue(initialValue)
	return m
}

//////////////////////////////////////////////////
// multiline

func (m *PlainTextInput) setMultiline(multiline bool) {
	m.multiline = multiline
	m.optionals.Multiline = true
}

func (m *PlainTextInput) removeMultiline() {
	m.multiline = false
	m.optionals.Multiline = false
}

// EnableMultiline public update multiline
func (m PlainTextInput) EnableMultiline() PlainTextInput {
	m.setMultiline(true)
	return m
}

func (m PlainTextInput) DisableMultiline() PlainTextInput {
	m.setMultiline(false)
	return m
}

//////////////////////////////////////////////////
// minLength

func (m *PlainTextInput) setMinLength(minLength int) {
	m.minLength = minLength
	m.optionals.MinLength = true
}

func (m *PlainTextInput) removeMinLength() {
	m.minLength = 0
	m.optionals.MinLength = false
}

// UpdateMinLength public update min length
func (m PlainTextInput) SetMinLength(minLength int) PlainTextInput {
	m.setMinLength(minLength)
	return m
}

//////////////////////////////////////////////////
// maxLength

func (m *PlainTextInput) setMaxLength(maxLength int) {
	m.maxLength = maxLength
	m.optionals.MaxLength = true
}

func (m *PlainTextInput) removeMaxLength() {
	m.maxLength = 0
	m.optionals.MaxLength = false
}

// UpdateMaxLength public update max length
func (m PlainTextInput) SetMaxLength(maxLength int) PlainTextInput {
	m.setMaxLength(maxLength)
	return m
}

//////////////////////////////////////////////////
// dispatchActionConfig

func (m *PlainTextInput) setDispatchActionConfig(dispatchActionConfig DispatchActionConfig) {
	m.dispatchActionConfig = dispatchActionConfig
	m.optionals.DispatchActionConfig = true
}

func (m *PlainTextInput) removeDispatchActionConfig() {
	m.dispatchActionConfig = DispatchActionConfig{}
	m.optionals.DispatchActionConfig = false
}

// AddDispatchActionConfig public update dispatch action config
func (m PlainTextInput) AddDispatchActionConfig(dispatchActionConfig DispatchActionConfig) PlainTextInput {
	m.setDispatchActionConfig(dispatchActionConfig)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *PlainTextInput) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m PlainTextInput) FocusOnLoad() PlainTextInput {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m PlainTextInput) UnsetFocusOnLoad() PlainTextInput {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *PlainTextInput) setPlaceholder(placeholder string) {
	m.placeholder = NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *PlainTextInput) removePlaceholder() {
	m.optionals.Placeholder = false
}

// AddPlaceholder public set placeholder
func (m PlainTextInput) AddPlaceholder(placeholder string) PlainTextInput {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m PlainTextInput) RemovePlaceholder() PlainTextInput {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// abstraction

type plainTextInputAbstraction struct {
	Type     string
	ActionID string

	InitialValue         string
	Multiline            bool
	MinLength            int
	MaxLength            int
	DispatchActionConfig DispatchActionConfig
	FocusOnLoad          bool
	Placeholder          CompositionText

	Optionals plainTextInputOptions
}

// abstraction
func (m PlainTextInput) abstraction() plainTextInputAbstraction {
	return plainTextInputAbstraction{
		Type:     m.slackType.String(),
		ActionID: m.actionID,

		InitialValue:         m.initialValue,
		Multiline:            m.multiline,
		MinLength:            m.minLength,
		MaxLength:            m.maxLength,
		DispatchActionConfig: m.dispatchActionConfig,
		FocusOnLoad:          m.focusOnLoad,
		Placeholder:          m.placeholder,

		Optionals: m.optionals,
	}
}

// template

func (m plainTextInputAbstraction) Template() string {
	return `{
"type": "{{.Type}}",
"action_id": "{{.ActionID}}"

{{if .Optionals.InitialValue}},
	"initial_value": "{{.InitialValue}}"
{{end}}

{{if .Optionals.Multiline}},
	"multiline": {{.Multiline}}
{{end}}

{{if .Optionals.MinLength}},
	"min_length": {{.MinLength}}
{{end}}
	
{{if .Optionals.MaxLength}},
	"max_length": {{.MaxLength}}
{{end}}

{{if .Optionals.DispatchActionConfig}},
	"dispatch_action_config": {{.DispatchActionConfig.Render}}
{{end}}

{{if .Optionals.FocusOnLoad}},
	"focus_on_load": {{.FocusOnLoad}}
{{end}}

{{if .Optionals.Placeholder}},
	"placeholder": {{.Placeholder.Render}}
{{end}}

}`
}

// Render
func (m PlainTextInput) Render() string {
	raw := Render(m.abstraction())
	return Pretty(raw)
}

type RadioButton struct {
	slackType ElementType
	actionID  string

	options       []Option
	initialOption Option
	ConfirmationDialog
	focusOnLoad bool

	optionals radioButtonOptions
}

type radioButtonOptions struct {
	InitialOption bool
	Confirm       bool
	FocusOnLoad   bool
}

// NewRadioButton public constructor
func NewRadioButton(actionId string) RadioButton {
	return RadioButton{
		slackType: RadioButtonElement,
		actionID:  actionId,
		options:   []Option{},
		optionals: radioButtonOptions{
			InitialOption: false,
			Confirm:       false,
			FocusOnLoad:   false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *RadioButton) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *RadioButton) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m RadioButton) UpdateActionId(actionId string) RadioButton {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// options

// AddOption adds an option to the RadioButton
func (m *RadioButton) addOption(option Option) {
	m.options = append(m.options, option)
}

// AddOptions adds multiple options to the RadioButton
func (m *RadioButton) addOptions(options []Option) {
	for _, option := range options {
		m.addOption(option)
	}
}

// RemoveOption removes an option from the RadioButton
func (m *RadioButton) removeOption(option Option) {
	for i, v := range m.options {
		if v == option {
			m.options = append(m.options[:i], m.options[i+1:]...)
		}
	}
}

// RemoveOptions removes multiple options from the RadioButton
func (m *RadioButton) removeOptions(options []Option) {
	for _, option := range options {
		m.removeOption(option)
	}
}

// AddOption public update options
func (m RadioButton) AddOption(option Option) RadioButton {
	m.addOption(option)
	return m
}

// AddOptions public update options
func (m RadioButton) AddOptions(options []Option) RadioButton {
	m.addOptions(options)
	return m
}

// RemoveOption public update options
func (m RadioButton) RemoveOption(option Option) RadioButton {
	m.removeOption(option)
	return m
}

// RemoveOptions public update options
func (m RadioButton) RemoveOptions(options []Option) RadioButton {
	m.removeOptions(options)
	return m
}

//////////////////////////////////////////////////
// initialOption

func (m *RadioButton) setInitialOption(option Option) {
	m.initialOption = option
	m.optionals.InitialOption = true
}

func (m *RadioButton) removeInitialOption() {
	m.initialOption = Option{}
	m.optionals.InitialOption = false
}

// UpdateInitialOption public update initialOption
func (m RadioButton) UpdateInitialOption(option Option) RadioButton {
	m.setInitialOption(option)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *RadioButton) setConfirm(confirm ConfirmationDialog) {
	m.ConfirmationDialog = confirm
	m.optionals.Confirm = true
}

func (m *RadioButton) removeConfirm() {
	m.ConfirmationDialog = ConfirmationDialog{}
	m.optionals.Confirm = false
}

// UpdateConfirm public update confirm
func (m RadioButton) AddConfirmationDialog(confirm ConfirmationDialog) RadioButton {
	m.setConfirm(confirm)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *RadioButton) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

func (m *RadioButton) removeFocusOnLoad() {
	m.focusOnLoad = false
	m.optionals.FocusOnLoad = false
}

// UpdateFocusOnLoad public update focusOnLoad
func (m RadioButton) FocusOnLoad() RadioButton {
	m.setFocusOnLoad(true)
	return m
}

//////////////////////////////////////////////////
// abstraction

type radioButtonAbstraction struct {
	Type     string
	ActionID string

	Options       []Option
	InitialOption Option
	Confirm       ConfirmationDialog
	FocusOnLoad   bool

	Optionals radioButtonOptions
}

func (m RadioButton) abstraction() radioButtonAbstraction {
	return radioButtonAbstraction{
		Type:     m.slackType.String(),
		ActionID: m.actionID,

		Options:       m.options,
		InitialOption: m.initialOption,
		Confirm:       m.ConfirmationDialog,
		FocusOnLoad:   m.focusOnLoad,

		Optionals: m.optionals,
	}
}

// Render
func (m RadioButton) Render() string {
	output := Render(m.abstraction())
	return Pretty(output)
}

// template

func (m radioButtonAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"action_id": "{{.ActionID}}",
	"options": [{{range $index, $option := .Options}}{{if $index}},{{end}}{{ $option.Render}}{{end}}]

{{if .Optionals.InitialOption}}
	"initial_option": {{.InitialOption.Render}}
{{end}}

{{if .Optionals.Confirm}}
	"confirm": {{.Confirm.Render}}
{{end}}

{{if .Optionals.FocusOnLoad}}
	"focus_on_load": {{.FocusOnLoad}}
{{end}}
}`
}

// InputElement

type SelectMenuWithConversationsList struct {
	slackType ElementType
	actionID  string

	confirm            ConfirmationDialog
	responseUrlEnabled bool
	focusOnLoad        bool
	placeholder        CompositionText

	// Conversation
	defaultToCurrentConversation bool
	initialConversation          string
	filter                       Filter

	optionals selectMenuWithConversationsListOptions
}

type selectMenuWithConversationsListOptions struct {
	Confirm bool

	FocusOnLoad bool
	Placeholder bool

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversation          bool
	Filter                       bool
	ResponseUrlEnabled           bool
}

// abstracted type
type selectMenuWithConversationsListAbstraction struct {
	Type     string
	ActionId string

	Confirm ConfirmationDialog

	FocusOnLoad bool
	Placeholder CompositionText

	// Conversation
	DefaultToCurrentConversation bool
	InitialConversation          string
	Filter                       Filter
	ResponseUrlEnabled           bool

	Optionals selectMenuWithConversationsListOptions
}

func NewSelectMenuWithConversationsList(actionId string) SelectMenuWithConversationsList {
	return SelectMenuWithConversationsList{
		slackType: SelectMenuWithConversationsListElement,
		actionID:  actionId,
		optionals: selectMenuWithConversationsListOptions{
			Confirm:                      false,
			FocusOnLoad:                  false,
			Placeholder:                  false,
			DefaultToCurrentConversation: false,
			InitialConversation:          false,
			Filter:                       false,
			ResponseUrlEnabled:           false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *SelectMenuWithConversationsList) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *SelectMenuWithConversationsList) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m SelectMenuWithConversationsList) UpdateActionId(actionId string) SelectMenuWithConversationsList {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithConversationsList) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithConversationsList) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithConversationsList) AddConfirmDialog(confirm ConfirmationDialog) SelectMenuWithConversationsList {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m SelectMenuWithConversationsList) RemoveConfirmDialog() SelectMenuWithConversationsList {
	m.removeConfirm()
	return m
}

//////////////////////////////////////////////////
// responseUrlEnabled

func (m *SelectMenuWithConversationsList) setResponseUrlEnabled(responseUrlEnabled bool) {
	m.responseUrlEnabled = responseUrlEnabled
	m.optionals.ResponseUrlEnabled = responseUrlEnabled
}

func (m *SelectMenuWithConversationsList) unsetResponseUrlEnabled() {
	m.optionals.ResponseUrlEnabled = false
}

func (m SelectMenuWithConversationsList) EnableResponseUrlEnabled() SelectMenuWithConversationsList {
	m.setResponseUrlEnabled(true)
	return m
}

func (m SelectMenuWithConversationsList) DisableResponseUrlEnabled() SelectMenuWithConversationsList {
	m.setResponseUrlEnabled(false)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *SelectMenuWithConversationsList) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m SelectMenuWithConversationsList) FocusOnLoad() SelectMenuWithConversationsList {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m SelectMenuWithConversationsList) UnsetFocusOnLoad() SelectMenuWithConversationsList {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *SelectMenuWithConversationsList) setPlaceholder(placeholder string) {
	m.placeholder = NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *SelectMenuWithConversationsList) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m SelectMenuWithConversationsList) AddPlaceholder(placeholder string) SelectMenuWithConversationsList {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m SelectMenuWithConversationsList) RemovePlaceholder() SelectMenuWithConversationsList {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// defaultToCurrentConversation

// setDefaultToCurrentConversation public set default to current conversation
func (m *SelectMenuWithConversationsList) setDefaultToCurrentConversation(defaultToCurrentConversation bool) {
	m.defaultToCurrentConversation = defaultToCurrentConversation
	m.optionals.DefaultToCurrentConversation = defaultToCurrentConversation

}

// unsetDefaultToCurrentConversation public remove default to current conversation
func (m *SelectMenuWithConversationsList) unsetDefaultToCurrentConversation() {
	m.setDefaultToCurrentConversation(false)
}

// DefaultToCurrentConversation public set default to current conversation
func (m SelectMenuWithConversationsList) DefaultToCurrentConversation() SelectMenuWithConversationsList {
	m.setDefaultToCurrentConversation(true)
	return m
}

// UnsetDefaultToCurrentConversation public remove default to current conversation
func (m SelectMenuWithConversationsList) UnsetDefaultToCurrentConversation() SelectMenuWithConversationsList {
	m.unsetDefaultToCurrentConversation()
	return m
}

//////////////////////////////////////////////////
// initialConversations

// addInitialConversation private add initial conversation
func (m *SelectMenuWithConversationsList) addInitialConversation(initialConversation string) {
	m.initialConversation = initialConversation
	m.optionals.InitialConversation = true
}

// removeInitialConversations private remove initial conversations
func (m *SelectMenuWithConversationsList) removeInitialConversation() {
	m.optionals.InitialConversation = false
}

// SetInitialConversation public add initial conversation
func (m SelectMenuWithConversationsList) SetInitialConversation(initialConversation string) SelectMenuWithConversationsList {
	m.addInitialConversation(initialConversation)
	return m
}

// ClearInitialConversations clear initial conversations
func (m SelectMenuWithConversationsList) UnsetInitialConversation() SelectMenuWithConversationsList {
	m.removeInitialConversation()
	return m
}

//////////////////////////////////////////////////
// filter

func (m *SelectMenuWithConversationsList) setFilter(filter Filter) {
	m.filter = filter
	m.optionals.Filter = true
}

func (m *SelectMenuWithConversationsList) removeFilter() {
	m.optionals.Filter = false
}

// AddFilter public set filter
func (m SelectMenuWithConversationsList) AddFilter(filter Filter) SelectMenuWithConversationsList {
	m.setFilter(filter)
	return m
}

// RemoveFilter public remove filter
func (m SelectMenuWithConversationsList) RemoveFilter() SelectMenuWithConversationsList {
	m.removeFilter()
	return m
}

// create abstract
func (m SelectMenuWithConversationsList) abstraction() selectMenuWithConversationsListAbstraction {
	return selectMenuWithConversationsListAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm: m.confirm,

		FocusOnLoad: m.focusOnLoad,
		Placeholder: m.placeholder,

		// Conversation
		DefaultToCurrentConversation: m.defaultToCurrentConversation,
		InitialConversation:          m.initialConversation,
		Filter:                       m.filter,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m selectMenuWithConversationsListAbstraction) Template() string {
	return `{	
"type": "{{ .Type }}",
"action_id": "{{ .ActionId }}"

{{if .Optionals.InitialConversation}},
	"initial_conversation":  "{{ .InitialConversation }}"
{{end}}

{{if .Optionals.DefaultToCurrentConversation}},
	"default_to_current_conversation": {{ .DefaultToCurrentConversation }}
{{end}}

{{if .Optionals.Confirm }},
	"confirm": {{ .Confirm.Render }}
{{end}}

{{ if .Optionals.ResponseUrlEnabled }},
	"response_url_enabled": {{ .ResponseUrlEnabled }}
{{end}}

{{if .Optionals.Filter }},
	{{ .Filter.Render }}
{{end}}

{{if .Optionals.FocusOnLoad }},
	"focus_on_load": {{ .FocusOnLoad }}
{{end}}

{{if .Optionals.Placeholder }},
	"placeholder": {{ .Placeholder.Render }}
{{end}}

}`
}

func (m SelectMenuWithConversationsList) ElementRender() {}

func (m SelectMenuWithConversationsList) Render() string {
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m SelectMenuWithConversationsList) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

// InputElement

type SelectMenuWithExternalDataSource struct {
	slackType ElementType
	actionID  string

	initialOption Option
	confirm       ConfirmationDialog
	focusOnLoad   bool
	placeholder   CompositionText

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
		initialOption: Option{},
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

func (m *SelectMenuWithExternalDataSource) setInitialOption(initialOption Option) {
	m.initialOption = initialOption
	m.optionals.InitialOption = true
}

// ClearInitialOptions clear initial options
func (m SelectMenuWithExternalDataSource) ClearInitialOption() SelectMenuWithExternalDataSource {
	m.removeInitialOption()
	return m
}

// AddInitialOption public add initial option
func (m SelectMenuWithExternalDataSource) AddInitialOption(initialOption Option) SelectMenuWithExternalDataSource {
	m.setInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithExternalDataSource) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithExternalDataSource) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithExternalDataSource) AddConfirmDialog(confirm ConfirmationDialog) SelectMenuWithExternalDataSource {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m SelectMenuWithExternalDataSource) RemoveConfirmDialog() SelectMenuWithExternalDataSource {
	m.removeConfirm()
	return m
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
	m.placeholder = NewPlainText(placeholder)
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

	InitialOption Option
	Confirm       ConfirmationDialog
	FocusOnLoad   bool
	Placeholder   CompositionText

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
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m SelectMenuWithExternalDataSource) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

// InputElement

type SelectMenuWithPublicChannelsSelect struct {
	slackType ElementType
	actionID  string

	confirm            ConfirmationDialog
	responseUrlEnabled bool
	focusOnLoad        bool
	placeholder        CompositionText

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

	Confirm            ConfirmationDialog
	ResponseUrlEnabled bool
	FocusOnLoad        bool
	Placeholder        CompositionText

	// Public Channel
	InitialChannel string

	Optionals SelectMenuWithPublicChannelsSelectOptions
}

func NewSelectMenuWithPublicChannelsSelect(actionId string) SelectMenuWithPublicChannelsSelect {
	return SelectMenuWithPublicChannelsSelect{
		slackType: SelectMenuWithPublicChannelsSelectElement,
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

func (m *SelectMenuWithPublicChannelsSelect) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithPublicChannelsSelect) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithPublicChannelsSelect) AddConfirmDialog(confirm ConfirmationDialog) SelectMenuWithPublicChannelsSelect {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m SelectMenuWithPublicChannelsSelect) RemoveConfirmDialog() SelectMenuWithPublicChannelsSelect {
	m.removeConfirm()
	return m
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
	m.placeholder = NewPlainText(placeholder)
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
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m SelectMenuWithPublicChannelsSelect) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

// InputElement

type SelectMenuWithStaticOption struct {
	slackType ElementType
	actionID  string
	options   []Option

	optionGroups  []OptionGroup
	initialOption Option
	confirm       ConfirmationDialog

	focusOnLoad bool
	placeholder CompositionText

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
		options:   []Option{},
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

func (m *SelectMenuWithStaticOption) setOptions(options []Option) {
	m.options = options
}

func (m *SelectMenuWithStaticOption) addOption(option Option) {
	m.options = append(m.options, option)
}

func (m *SelectMenuWithStaticOption) removeOptions() {
	m.options = []Option{}
}

// AddOption public add option
func (m SelectMenuWithStaticOption) AddOption(option Option) SelectMenuWithStaticOption {
	m.addOption(option)
	return m
}

// ClearOptions clear options
func (m SelectMenuWithStaticOption) ClearOptions() SelectMenuWithStaticOption {
	m.removeOptions()
	return m
}

func (m *SelectMenuWithStaticOption) setOptionGroups(optionGroups []OptionGroup) {
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
func (m SelectMenuWithStaticOption) AddOptionGroup(optionGroup OptionGroup) SelectMenuWithStaticOption {
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

func (m *SelectMenuWithStaticOption) addInitialOption(initialOption Option) {
	m.addOption(initialOption)
	m.initialOption = initialOption
	m.optionals.InitialOption = true
}

func (m *SelectMenuWithStaticOption) removeInitialOptions() {
	m.optionals.InitialOption = false
}

func (m *SelectMenuWithStaticOption) setInitialOptions(initialOption Option) {
	m.initialOption = initialOption
	m.optionals.InitialOption = true
}

// ClearInitialOptions clear initial options
func (m SelectMenuWithStaticOption) ClearInitialOptions() SelectMenuWithStaticOption {
	m.removeInitialOptions()
	return m
}

// AddInitialOption public add initial option
func (m SelectMenuWithStaticOption) AddInitialOption(initialOption Option) SelectMenuWithStaticOption {
	m.addInitialOption(initialOption)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithStaticOption) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithStaticOption) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithStaticOption) AddConfirmDialog(confirm ConfirmationDialog) SelectMenuWithStaticOption {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m SelectMenuWithStaticOption) RemoveConfirmDialog() SelectMenuWithStaticOption {
	m.removeConfirm()
	return m
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
	m.placeholder = NewPlainText(placeholder)
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
	Options          []Option
	OptionGroups     []OptionGroup
	InitialOption    Option
	Confirm          ConfirmationDialog
	MaxSelectedItems int
	FocusOnLoad      bool
	Placeholder      CompositionText

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
	raw := Render(m.abstraction())
	return Pretty(raw)
}

// ElementRender
func (m SelectMenuWithStaticOption) ElementRender() {}

// SectionBlock public section block
func (m SelectMenuWithStaticOption) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

// InputElement

type SelectMenuWithUserList struct {
	slackType ElementType
	actionID  string

	confirm     ConfirmationDialog
	focusOnLoad bool
	placeholder CompositionText

	// User List
	initialUser string

	optionals selectMenuWithUserListOptions
}

type selectMenuWithUserListOptions struct {
	Confirm     bool
	FocusOnLoad bool
	Placeholder bool

	// User List
	InitialUser bool
}

func NewSelectMenuWithUserList(actionId string) SelectMenuWithUserList {
	return SelectMenuWithUserList{
		slackType: SelectMenuWithUserListElement,
		actionID:  actionId,
		optionals: selectMenuWithUserListOptions{
			Confirm:     false,
			FocusOnLoad: false,
			Placeholder: false,
			InitialUser: false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *SelectMenuWithUserList) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *SelectMenuWithUserList) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m SelectMenuWithUserList) UpdateActionId(actionId string) SelectMenuWithUserList {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *SelectMenuWithUserList) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *SelectMenuWithUserList) removeConfirm() {
	m.optionals.Confirm = false
}

// AddConfirmDialog public set confirm
func (m SelectMenuWithUserList) AddConfirmDialog(confirm ConfirmationDialog) SelectMenuWithUserList {
	m.setConfirm(confirm)
	m.optionals.Confirm = true
	return m
}

// RemoveConfirmDialog public remove confirm
func (m SelectMenuWithUserList) RemoveConfirmDialog() SelectMenuWithUserList {
	m.removeConfirm()
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *SelectMenuWithUserList) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = focusOnLoad
}

// FocusOnLoad public set focus on load
func (m SelectMenuWithUserList) FocusOnLoad() SelectMenuWithUserList {
	m.setFocusOnLoad(true)
	return m
}

// UnsetFocusOnLoad public remove focus on load
func (m SelectMenuWithUserList) UnsetFocusOnLoad() SelectMenuWithUserList {
	m.setFocusOnLoad(false)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *SelectMenuWithUserList) setPlaceholder(placeholder string) {
	m.placeholder = NewPlainText(placeholder)
	m.optionals.Placeholder = true
}

func (m *SelectMenuWithUserList) removePlaceholder() {
	m.optionals.Placeholder = false
}

// SetPlaceholder public set placeholder
func (m SelectMenuWithUserList) AddPlaceholder(placeholder string) SelectMenuWithUserList {
	m.setPlaceholder(placeholder)
	return m
}

// RemovePlaceholder public remove placeholder
func (m SelectMenuWithUserList) RemovePlaceholder() SelectMenuWithUserList {
	m.optionals.Placeholder = false
	return m
}

//////////////////////////////////////////////////
// initialUsers

func (m *SelectMenuWithUserList) setInitialUser(initialUser string) {
	m.initialUser = initialUser
	m.optionals.InitialUser = true
}

func (m *SelectMenuWithUserList) removeInitialUser() {
	m.optionals.InitialUser = false
}

// AddInitialUser public add initial user
func (m SelectMenuWithUserList) SetInitialUser(initialUser string) SelectMenuWithUserList {
	m.setInitialUser(initialUser)
	return m
}

// ClearInitialUsers clear initial users
func (m SelectMenuWithUserList) ClearInitialUsers() SelectMenuWithUserList {
	m.removeInitialUser()
	return m
}

// ////////////////////////////////////////////////
// abstract

// abstracted type
type selectMenuWithUserListAbstraction struct {
	Type     string
	ActionId string

	Confirm     ConfirmationDialog
	FocusOnLoad bool
	Placeholder CompositionText

	// User List
	InitialUser string

	Optionals selectMenuWithUserListOptions
}

// create abstract
func (m SelectMenuWithUserList) abstraction() selectMenuWithUserListAbstraction {
	return selectMenuWithUserListAbstraction{
		Type:     m.slackType.String(),
		ActionId: m.actionID,

		Confirm:     m.confirm,
		FocusOnLoad: m.focusOnLoad,
		Placeholder: m.placeholder,

		// User List
		InitialUser: m.initialUser,

		Optionals: m.optionals,
	}
}

// Template returns template string
func (m selectMenuWithUserListAbstraction) Template() string {
	return `{
"type": "{{ .Type }}",
"action_id": "{{ .ActionId }}"

{{if .Optionals.InitialUser}},
	"initial_user": "{{ .InitialUser}}"
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

func (m SelectMenuWithUserList) ElementRender() {}

func (m SelectMenuWithUserList) Render() string {
	raw := Render(m.abstraction())
	return Pretty(raw)
}

func (m SelectMenuWithUserList) Section() Section {
	s := NewSection("newSection").AddAccessory(m)
	return s
}

type TimePicker struct {
	slackType ElementType
	actionID  string

	initialTime string
	confirm     ConfirmationDialog
	focusOnLoad bool
	placeholder CompositionText

	optionals timePickerOptions
}

type timePickerOptions struct {
	InitialTime bool
	Confirm     bool
	FocusOnLoad bool
	Placeholder bool
}

// NewTimePicker public constructor
func NewTimePicker(actionId string) TimePicker {
	return TimePicker{
		slackType: TimePickerElement,
		actionID:  actionId,
		optionals: timePickerOptions{
			InitialTime: false,
			Confirm:     false,
			FocusOnLoad: false,
			Placeholder: false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *TimePicker) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *TimePicker) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m TimePicker) UpdateActionId(actionId string) TimePicker {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// initialTime

func (m *TimePicker) setInitialTime(initialTime string) {
	m.initialTime = initialTime
}

func (m *TimePicker) removeInitialTime() {
	m.initialTime = ""
}

// UpdateInitialTime public update initialTime
func (m TimePicker) UpdateInitialTime(initialTime string) TimePicker {
	m.setInitialTime(initialTime)
	return m
}

//////////////////////////////////////////////////
// confirm

func (m *TimePicker) setConfirm(confirm ConfirmationDialog) {
	m.confirm = confirm
	m.optionals.Confirm = true
}

func (m *TimePicker) removeConfirm() {
	m.optionals.Confirm = false
}

// UpdateConfirm public update confirm
func (m TimePicker) UpdateConfirm(confirm ConfirmationDialog) TimePicker {
	m.setConfirm(confirm)
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *TimePicker) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *TimePicker) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// UpdateFocusOnLoad public update focusOnLoad
func (m TimePicker) UpdateFocusOnLoad(focusOnLoad bool) TimePicker {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *TimePicker) setPlaceholder(placeholder CompositionText) {
	m.placeholder = placeholder
	m.optionals.Placeholder = true
}

func (m *TimePicker) removePlaceholder() {
	m.optionals.Placeholder = false
}

// UpdatePlaceholder public update placeholder
func (m TimePicker) UpdatePlaceholder(placeholder string) TimePicker {
	m.setPlaceholder(NewPlainText(placeholder))
	return m
}

//////////////////////////////////////////////////
// abstract

type timePickerAbstract struct {
	Type        string
	ActionID    string
	InitialTime string
	Confirm     ConfirmationDialog
	FocusOnLoad bool
	Placeholder CompositionText

	Optionals timePickerOptions
}

// abstract public method
func (m TimePicker) abstraction() timePickerAbstract {
	return timePickerAbstract{
		Type:        m.slackType.String(),
		ActionID:    m.actionID,
		InitialTime: m.initialTime,
		Confirm:     m.confirm,
		FocusOnLoad: m.focusOnLoad,
		Placeholder: m.placeholder,

		Optionals: m.optionals,
	}
}

//////////////////////////////////////////////////
// template

// Template public method
func (m timePickerAbstract) Template() string {
	return `{
"type": "{{.Type}}",
"action_id": "{{.ActionID}}"

{{if .Optionals.InitialTime}},
	"initial_time": "{{.InitialTime}}"
{{end}}

{{if .Optionals.Confirm}},
	"confirm": {{.Confirm.Render}}
{{end}}
		
{{if .Optionals.FocusOnLoad}},
	"focus_on_load": {{.FocusOnLoad}}
{{end}}
		
{{if .Optionals.Placeholder}},
	"placeholder": {{.Placeholder.Render}}
{{end}}

	}`
}

//////////////////////////////////////////////////
// render

// Render public method
func (m TimePicker) Render() string {
	output := Render(m.abstraction())
	return Pretty(output)
}

type URLInput struct {
	slackType ElementType
	actionID  string

	initialValue         *url.URL
	dispatchActionConfig DispatchActionConfig
	focusOnLoad          bool
	placeholder          CompositionText

	optionals urlInputOptions
}

type urlInputOptions struct {
	InitialValue         bool
	DispatchActionConfig bool
	FocusOnLoad          bool
	Placeholder          bool
}

func NewURLInput(actionId string) URLInput {
	return URLInput{
		slackType: UrlInputElement,
		actionID:  actionId,
		optionals: urlInputOptions{
			InitialValue:         false,
			DispatchActionConfig: false,
			FocusOnLoad:          false,
			Placeholder:          false,
		},
	}
}

//////////////////////////////////////////////////
// actionID

func (m *URLInput) setActionId(actionId string) {
	m.actionID = actionId
}

func (m *URLInput) removeActionId() {
	m.actionID = ""
}

// UpdateActionId public update action id
func (m URLInput) UpdateActionId(actionId string) URLInput {
	m.setActionId(actionId)
	return m
}

//////////////////////////////////////////////////
// initialValue

func (m *URLInput) setInitialValue(initialValue *url.URL) {
	m.initialValue = initialValue
	m.optionals.InitialValue = true
}

func (m *URLInput) removeInitialValue() {
	m.optionals.InitialValue = false
}

// UpdateInitialValue public update initial value
func (m URLInput) UpdateInitialValue(initialValue *url.URL) URLInput {
	m.setInitialValue(initialValue)
	return m
}

// RemoveInitialValue public remove initial value
func (m URLInput) RemoveInitialValue() URLInput {
	m.removeInitialValue()
	return m
}

//////////////////////////////////////////////////
// dispatchActionConfig

func (m *URLInput) setDispatchActionConfig(dispatchActionConfig DispatchActionConfig) {
	m.dispatchActionConfig = dispatchActionConfig
	m.optionals.DispatchActionConfig = true
}

func (m *URLInput) removeDispatchActionConfig() {
	m.optionals.DispatchActionConfig = false
}

// UpdateDispatchActionConfig public update dispatch action config
func (m URLInput) UpdateDispatchActionConfig(dispatchActionConfig DispatchActionConfig) URLInput {
	m.setDispatchActionConfig(dispatchActionConfig)
	return m
}

// RemoveDispatchActionConfig public remove dispatch action config
func (m URLInput) RemoveDispatchActionConfig() URLInput {
	m.removeDispatchActionConfig()
	return m
}

//////////////////////////////////////////////////
// focusOnLoad

func (m *URLInput) setFocusOnLoad(focusOnLoad bool) {
	m.focusOnLoad = focusOnLoad
	m.optionals.FocusOnLoad = true
}

func (m *URLInput) removeFocusOnLoad() {
	m.optionals.FocusOnLoad = false
}

// UpdateFocusOnLoad public update focus on load
func (m URLInput) UpdateFocusOnLoad(focusOnLoad bool) URLInput {
	m.setFocusOnLoad(focusOnLoad)
	return m
}

// RemoveFocusOnLoad public remove focus on load
func (m URLInput) RemoveFocusOnLoad() URLInput {
	m.removeFocusOnLoad()
	return m
}

//////////////////////////////////////////////////
// placeholder

func (m *URLInput) setPlaceholder(placeholder CompositionText) {
	m.placeholder = placeholder
	m.optionals.Placeholder = true
}

func (m *URLInput) removePlaceholder() {
	m.optionals.Placeholder = false
}

// UpdatePlaceholder public update placeholder
func (m URLInput) UpdatePlaceholder(placeholder string) URLInput {
	m.setPlaceholder(NewPlainText(placeholder))
	return m
}

// RemovePlaceholder public remove placeholder
func (m URLInput) RemovePlaceholder() URLInput {
	m.removePlaceholder()
	return m
}

//////////////////////////////////////////////////
// abstraction

type urlInputAbstraction struct {
	Type     string
	ActionID string

	InitialValue         string
	DispatchActionConfig DispatchActionConfig
	FocusOnLoad          bool
	Placeholder          CompositionText

	Optionals urlInputOptions
}

// abstract method
func (m URLInput) abstraction() urlInputAbstraction {
	url := ""
	if m.optionals.InitialValue {
		url = m.initialValue.String()
	}

	return urlInputAbstraction{
		Type:     m.slackType.String(),
		ActionID: m.actionID,

		InitialValue:         url,
		DispatchActionConfig: m.dispatchActionConfig,
		FocusOnLoad:          m.focusOnLoad,
		Placeholder:          m.placeholder,

		Optionals: m.optionals,
	}
}

//////////////////////////////////////////////////
// render

// Render method
func (m URLInput) Render() string {
	output := Render(m.abstraction())
	return Pretty(output)
}

//////////////////////////////////////////////////
// template

// Template method
func (m urlInputAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"action_id": "{{.ActionID}}"

{{if .Optionals.InitialValue}},
	"initial_value": "{{.InitialValue}}"
{{end}}

{{if .Optionals.DispatchActionConfig}},
	"dispatch_action_config": {{.DispatchActionConfig.Render}}
{{end}}

{{if .Optionals.FocusOnLoad}},
	"focus_on_load": {{.FocusOnLoad}}
{{end}}

{{if .Optionals.Placeholder}},
	"placeholder": {{.Placeholder.Render}}
{{end}}
}
`
}
