package section

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

type Section struct {
	slackType block.BlockType
	text      compositiontext.CompositionText

	accessory element.Element
	blockId   string

	fields []compositiontext.CompositionText

	optional sectionOptions
}

type sectionOptions struct {
	Accessory bool
	BlockId   bool
	Field     bool
}

// NewSection creates a new section.
func NewSection(text string) Section {
	return Section{
		slackType: block.Section,
		text:      compositiontext.NewPlainText(text),
		optional: sectionOptions{
			Accessory: false,
			BlockId:   false,
			Field:     false,
		},
	}
}

// SetAccessory sets the accessory for the section.
func (s *Section) setAccessory(accessory element.Element) {
	s.accessory = accessory
	s.optional.Accessory = true
}

func (s *Section) removeAccessory() {
	s.optional.Accessory = false
}

// addAccessory adds an accessory to the section.
func (s Section) AddAccessory(accessory element.Element) Section {
	s.setAccessory(accessory)
	return s
}

// setBlockId sets the block id for the section.
func (s *Section) setBlockId(blockId string) {
	s.blockId = blockId
	s.optional.BlockId = true
}

func (s *Section) removeBlockId() {
	s.optional.BlockId = false
}

// AddField adds a field to the section.
func (s *Section) addPlainField(field string) {
	s.fields = append(s.fields, compositiontext.NewPlainText(field))
	s.optional.Field = true
}

func (s *Section) addMarkdownField(field string) {
	s.fields = append(s.fields, compositiontext.NewMrkdwnText(field))
	s.optional.Field = true
}

func (s *Section) removeField() {
	s.optional.Field = false
}

// abstraction for the section block
type sectionAbstraction struct {
	Type string
	Text compositiontext.CompositionText

	Accessory element.Element
	BlockId   string

	Fields []compositiontext.CompositionText

	Optional sectionOptions
}

// abstraction
func (s Section) abstraction() sectionAbstraction {
	return sectionAbstraction{
		Type: s.slackType.String(),
		Text: s.text,

		Accessory: s.accessory,
		BlockId:   s.blockId,

		Fields: s.fields,

		Optional: s.optional,
	}
}

// Template returns the template for the section.
func (s sectionAbstraction) Template() string {
	return `{
		"type": "{{.Type}}",
		"text": {{.Text.Render}}
{{if .Optional.Accessory}},
		"accessory": {{.Accessory.Render}}
{{end}}

{{if .Optional.BlockId}},
		"block_id": "{{.BlockId}}"
{{end}}

{{if .Optional.Field}},
		"fields": [
			{{range $index, $field := .Fields}}{{if $index}},{{end}}{{ $field.Render}}{{end}}
		]{{end}}
	}`
}

// Render renders the section to a string
func (s Section) Render() string {
	output := common.Render(s.abstraction())
	return common.Pretty(output)
}

type SectionType interface {
	Section()
}

func (s Section) String() string {
	return s.Render()
}
