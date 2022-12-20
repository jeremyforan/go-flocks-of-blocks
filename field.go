package flocksofblocks

import "fmt"

// not sure how to think about this. I am struggling becuase I am sure there is some efficiant to be had here. All
// fields share properties. They have a name, a value, constraints and a set of compatibilities.
// If composition can be used here than naming conventions would be more consistent and the code would be easier to manage.
// The struggle is that as I write, I seem to keep producing the same amount of code, despite the level of abstraction.

// here are all the fields from all of the blocks, elements and compositions objects

//type
//value

//accessibility_label
//action_id
//alt_text
//author_name
//external_id
//initial_channel
//initial_conversation
//initial_date
//initial_time
//initial_user
//initial_value
//provider_name
//source
//style
//text

//url
//image_url
//title_url
//provider_icon_url
//thumbnail_url

//max_value
//min_value

//go:generate stringer-type=fieldKey
type fieldKey string

const (

	// String Fields
	fieldBlockId             fieldKey = "block_id"
	fieldActionId            fieldKey = "action_id"
	fieldAccessibilityLabel  fieldKey = "accessibility_label"
	fieldAltText             fieldKey = "alt_text"
	fieldAuthorName          fieldKey = "author_name"
	fieldExternalId          fieldKey = "external_id"
	fieldInitialChannel      fieldKey = "initial_channel"
	fieldInitialConversation fieldKey = "initial_conversation"
	fieldInitialDate         fieldKey = "initial_date"
	fieldInitialTime         fieldKey = "initial_time"
	fieldInitialUser         fieldKey = "initial_user"
	fieldInitialValue        fieldKey = "initial_value"
	fieldProviderName        fieldKey = "provider_name"
	fieldSource              fieldKey = "source"
	fieldStyle               fieldKey = "style"
	fieldText                fieldKey = "text"

	// Bool Fields
	fieldDefaultToCurrentConversation  fieldKey = "default_to_current_conversation"
	fieldDispatchAction                fieldKey = "dispatch_action"
	fieldEmoji                         fieldKey = "emoji"
	fieldExcludeBotUsers               fieldKey = "exclude_bot_users"
	fieldExcludeExternalSharedChannels fieldKey = "exclude_external_shared_channels"
	fieldFocusOnLoad                   fieldKey = "focus_on_load"
	fieldIsDecimalAllowed              fieldKey = "is_decimal_allowed"
	fieldMultiline                     fieldKey = "multiline"
	fieldOptionalInput                 fieldKey = "optional"
	fieldResponseUrlEnabled            fieldKey = "response_url_enabled"
	fieldVerbatim                      fieldKey = "verbatim"

	//int field
	fieldInitialDateTime  fieldKey = "initial_date_time"
	fieldMaxLength        fieldKey = "max_length"
	fieldMaxSelectedItems fieldKey = "max_selected_items"
	fieldMinLength        fieldKey = "min_length"
	fieldMinQueryLength   fieldKey = "min_query_length"
)

type stringField struct {
	name  fieldKey
	value string
	set   bool
}

func (s stringField) isSet() bool {
	return s.set
}

func (s stringField) render() string {
	return fmt.Sprintf("\"%s\":\"%s\"", s.name, s.value)
}

func (s *stringField) SetValue(value string) {
	s.value = value
	s.set = true
}

func (s *stringField) SetName(name fieldKey) {
	s.name = name
}

func (s *stringField) UnsetValue() {
	s.value = ""
	s.set = false
}

func (s stringField) String() string {
	return s.render()
}

type blockId struct {
	stringField
}
type accessibilityLabel struct {
	stringField
}
type actionId struct {
	stringField
}
type altText struct {
	stringField
}
type authorName struct {
	stringField
}
type externalId struct {
	stringField
}
type initialChannel struct {
	stringField
}
type initialConversation struct {
	stringField
}
type initialDate struct {
	stringField
}
type initialTime struct {
	stringField
}
type initialUser struct {
	stringField
}
type initialValue struct {
	stringField
}
type providerName struct {
	stringField
}
type source struct {
	stringField
}
type style struct {
	stringField
}
type text struct {
	stringField
}

func (f *accessibilityLabel) String() string {
	f.stringField.SetName(fieldAccessibilityLabel)
	return f.stringField.String()
}

func (f *actionId) String() string {
	f.stringField.SetName(fieldActionId)
	return f.stringField.String()
}

func (f *altText) String() string {
	f.stringField.SetName(fieldAltText)
	return f.stringField.String()
}

func (f *authorName) String() string {
	f.stringField.SetName(fieldAuthorName)
	return f.stringField.String()
}

func (f *blockId) String() string {
	f.stringField.SetName(fieldBlockId)
	return f.stringField.String()
}

func (f *externalId) String() string {
	f.stringField.SetName(fieldExternalId)
	return f.stringField.String()
}

func (f *initialChannel) String() string {
	f.stringField.SetName(fieldInitialChannel)
	return f.stringField.String()
}

func (f *initialConversation) String() string {
	f.stringField.SetName(fieldInitialConversation)
	return f.stringField.String()
}

func (f *initialDate) String() string {
	f.stringField.SetName(fieldInitialDate)
	return f.stringField.String()
}

func (f *initialTime) String() string {
	f.stringField.SetName(fieldInitialTime)
	return f.stringField.String()
}

func (f *initialUser) String() string {
	f.stringField.SetName(fieldInitialUser)
	return f.stringField.String()
}

func (f *initialValue) String() string {
	f.stringField.SetName(fieldInitialValue)
	return f.stringField.String()
}

func (f *providerName) String() string {
	f.stringField.SetName(fieldProviderName)
	return f.stringField.String()
}

func (f *source) String() string {
	f.stringField.SetName(fieldSource)
	return f.stringField.String()
}

func (f *style) String() string {
	f.stringField.SetName(fieldStyle)
	return f.stringField.String()
}

func (f *text) String() string {
	f.stringField.SetName(fieldText)
	return f.stringField.String()
}

// //////////////////////////////////////
// boolField

type boolField struct {
	name  fieldKey
	value bool
	set   bool
}

func (f boolField) render() string {
	return fmt.Sprintf("\"%s\":%v", f.name, f.value)
}

func (f boolField) isSet() bool {
	return f.set
}

func (f *boolField) SetValue(value bool) {
	f.value = value
	f.set = true
}

func (f *boolField) SetName(name fieldKey) {
	f.name = name
}

func (f *boolField) UnsetValue() {
	f.value = false
	f.set = false
}

func (f boolField) String() string {
	return f.render()
}

type defaultToCurrentConversation struct {
	*boolField
}
type dispatchAction struct {
	*boolField
}
type emoji struct {
	*boolField
}
type excludeBotUsers struct {
	*boolField
}
type excludeExternalSharedChannels struct {
	*boolField
}
type focusOnLoad struct {
	*boolField
}
type isDecimalAllowed struct {
	*boolField
}
type multiline struct {
	*boolField
}
type optionalInput struct {
	*boolField
}
type responseUrlEnabled struct {
	*boolField
}
type verbatim struct {
	*boolField
}

func (f *defaultToCurrentConversation) String() string {
	f.boolField.SetName(fieldDefaultToCurrentConversation)
	return f.boolField.String()
}
func (f *dispatchAction) String() string {
	f.boolField.SetName(fieldDispatchAction)
	return f.boolField.String()
}
func (f *emoji) String() string {
	f.boolField.SetName(fieldEmoji)
	return f.boolField.String()
}
func (f *excludeBotUsers) String() string {
	f.boolField.SetName(fieldExcludeBotUsers)
	return f.boolField.String()
}
func (f *excludeExternalSharedChannels) String() string {
	f.boolField.SetName(fieldExcludeExternalSharedChannels)
	return f.boolField.String()
}
func (f *focusOnLoad) String() string {
	f.boolField.SetName(fieldFocusOnLoad)
	return f.boolField.String()
}
func (f *isDecimalAllowed) String() string {
	f.boolField.SetName(fieldIsDecimalAllowed)
	return f.boolField.String()
}
func (f *multiline) String() string {
	f.boolField.SetName(fieldMultiline)
	return f.boolField.String()
}
func (f *optionalInput) String() string {
	f.boolField.SetName(fieldOptionalInput)
	return f.boolField.String()
}
func (f *responseUrlEnabled) String() string {
	f.boolField.SetName(fieldResponseUrlEnabled)
	return f.boolField.String()
}
func (f *verbatim) String() string {
	f.boolField.SetName(fieldVerbatim)
	return f.boolField.String()
}

// Int field
type intField struct {
	name  fieldKey
	value int
	set   bool
}

func (f intField) render() string {
	return fmt.Sprintf("\"%s\":%d", f.name, f.value)
}

func (f intField) isSet() bool {
	return f.set
}

func (f *intField) SetValue(value int) {
	f.value = value
	f.set = true
}

func (f *intField) SetName(name fieldKey) {
	f.name = name
}

func (f *intField) UnsetValue() {
	f.value = 1
	f.set = false
}

func (f intField) String() string {
	return f.render()
}

type initialDateTime struct {
	*intField
}
type maxLength struct {
	*intField
}
type maxSelectedItems struct {
	*intField
}
type minLength struct {
	*intField
}
type minQueryLength struct {
	*intField
}

func (f *initialDateTime) String() string {
	f.intField.SetName(fieldInitialDateTime)
	return f.intField.String()
}
func (f *maxLength) String() string {
	f.intField.SetName(fieldMaxLength)
	return f.intField.String()
}
func (f *maxSelectedItems) String() string {
	f.intField.SetName(fieldMaxSelectedItems)
	return f.intField.String()
}
func (f *minLength) String() string {
	f.intField.SetName(fieldMinLength)
	return f.intField.String()
}
func (f *minQueryLength) String() string {
	f.intField.SetName(fieldMinQueryLength)
	return f.intField.String()
}
