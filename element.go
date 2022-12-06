package flocksofblocks

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
	ImageElement                                   ElementType = "image"
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
