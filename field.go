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
//block_id
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
	fieldBlockId  fieldKey = "block_id"
	fieldActionId fieldKey = "action_id"
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

// /////////////////////////////
// BlockId sets the block id of the block

type blockId struct {
	stringField
}

func (f blockId) String() string {
	f.stringField.SetName(fieldBlockId)
	return f.stringField.String()
}

type ActionId struct {
	stringField
}

func (f ActionId) String() string {
	f.stringField.SetName(fieldActionId)
	return f.stringField.String()
}
