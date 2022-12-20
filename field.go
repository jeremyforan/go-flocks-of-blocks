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
)

type stringField struct {
	name  fieldKey
	value string
	set   bool
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

func (s stringField) render() string {
	return fmt.Sprintf("\"%s\":\"%s\"", s.name, s.value)
}

// Render into a string
func (s stringField) String() string {
	return s.render()
}

func (s stringField) isSet() bool {
	return s.set
}

// /////////////////////////////
// BlockId sets the block id of the block

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

func (f accessibilityLabel) String() string {
	f.stringField.SetName(fieldAccessibilityLabel)
	return f.stringField.String()
}

func (f actionId) String() string {
	f.stringField.SetName(fieldActionId)
	return f.stringField.String()
}

func (f altText) String() string {
	f.stringField.SetName(fieldAltText)
	return f.stringField.String()
}

func (f authorName) String() string {
	f.stringField.SetName(fieldAuthorName)
	return f.stringField.String()
}

func (f externalId) String() string {
	f.stringField.SetName(fieldExternalId)
	return f.stringField.String()
}

func (f initialChannel) String() string {
	f.stringField.SetName(fieldInitialChannel)
	return f.stringField.String()
}

func (f initialConversation) String() string {
	f.stringField.SetName(fieldInitialConversation)
	return f.stringField.String()
}

func (f initialDate) String() string {
	f.stringField.SetName(fieldInitialDate)
	return f.stringField.String()
}

func (f initialTime) String() string {
	f.stringField.SetName(fieldInitialTime)
	return f.stringField.String()
}

func (f initialUser) String() string {
	f.stringField.SetName(fieldInitialUser)
	return f.stringField.String()
}

func (f initialValue) String() string {
	f.stringField.SetName(fieldInitialValue)
	return f.stringField.String()
}

func (f providerName) String() string {
	f.stringField.SetName(fieldProviderName)
	return f.stringField.String()
}

func (f source) String() string {
	f.stringField.SetName(fieldSource)
	return f.stringField.String()
}

func (f style) String() string {
	f.stringField.SetName(fieldStyle)
	return f.stringField.String()
}

func (f text) String() string {
	f.stringField.SetName(fieldText)
	return f.stringField.String()
}
