package element

type Element interface {
	ElementRender()
}

type ElementType string

const (
	Button                                  ElementType = "button"
	Checkboxes                              ElementType = "checkboxes"
	DatePicker                              ElementType = "datepicker"
	DateTimePicker                          ElementType = "datetimepicker"
	EmailInput                              ElementType = "email"
	Image                                   ElementType = "image"
	MultiSelectMenuWithStaticOptions        ElementType = "multi_static_select"
	MultiSelectMenuWithExternalDataSource   ElementType = "multi_external_select"
	MultiSelectMenuWithUserList             ElementType = "multi_users_select"
	MultiSelectMenuWithConversationsList    ElementType = "multi_conversations_select"
	MultiSelectMenuWithPublicChannelsSelect ElementType = "multi_channels_select"
	NumberInput                             ElementType = "number_input"
	OverflowMenu                            ElementType = "overflow"
	PlainTextInput                          ElementType = "plain_text_input"
	RadioButton                             ElementType = "radio_buttons"
	SelectMenuWithStaticOptions             ElementType = "static_select"
	SelectMenuWithExternalDataSource        ElementType = "external_select"
	SelectMenuWithUserList                  ElementType = "users_select"
	SelectMenuWithConversationsList         ElementType = "conversations_select"
	SelectMenuWithPublicChannelsSelect      ElementType = "channels_select"
	TimePicker                              ElementType = "timepicker"
	UrlInput                                ElementType = "url_text_input"
)

// stringer for ElementType
func (e ElementType) String() string {
	return string(e)
}

type InputElement interface {
	InputElement()
}
