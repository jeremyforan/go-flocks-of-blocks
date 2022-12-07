package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/pkg"
	"net/url"
)

type BlockType string

const (
	ActionsBlock BlockType = "actions"
	ContextBlock BlockType = "context"
	DividerBlock BlockType = "divider"
	FileBlock    BlockType = "file"
	HeaderBlock  BlockType = "header"
	ImageBlock   BlockType = "image"
	InputBlock   BlockType = "input"
	SectionBlock BlockType = "section"
	VideoBlock   BlockType = "video"
)

// stringer for BlockType
func (b BlockType) String() string {
	return string(b)
}

type Block interface {
	BlockRender()
}

type Action struct {
	slackType BlockType
	elements  []flocksofblocks.Element
	blockId   string

	optionals actionOptions
}

func (a Action) BlockRender() {}

func NewAction(blockId string) Action {
	return Action{
		slackType: ActionsBlock,
		elements:  []flocksofblocks.Element{},
		blockId:   blockId,
		optionals: actionOptions{
			blockId: false,
		},
	}
}

// SetBlockId sets the block id for the block.
func (a *Action) setBlockId(blockId string) {
	a.blockId = blockId
	a.optionals.blockId = true
}

func (a *Action) removeBlockId() {
	a.blockId = ""
	a.optionals.blockId = false
}

func (a *Action) addElement(element flocksofblocks.Element) {
	a.elements = append(a.elements, element)
}

type actionOptions struct {
	blockId bool
}

type actionAbstraction struct {
	Type     string
	Elements []flocksofblocks.Element
	BlockId  string

	Optional actionOptions
}

// AddBlockId chain function to add block id to an existing action block
func (a Action) AddBlockId(blockId string) Action {
	a.setBlockId(blockId)
	return a
}

// RemoveBlockId remove add block id from action block
func (a Action) RemoveBlockId() Action {
	a.removeBlockId()
	return a
}

// AddElement Add element to existing action block.
func (a Action) AddElement(element flocksofblocks.Element) Action {
	a.addElement(element)
	return a
}

// generate abstraction from action
func (a Action) abstraction() actionAbstraction {
	return actionAbstraction{
		Type:     string(a.slackType),
		Elements: a.elements,
		BlockId:  a.blockId,

		Optional: a.optionals,
	}
}

func (a actionAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	
	"elements": [{{range $index, $element := .Elements}}{{if $index}},{{end}}{{$element.Render}}{{end}}]

{{if .BlockId}},
	"block_id": "{{.BlockId}}"
{{end}}
}`
}

// Render the block
func (a Action) Render() string {
	output := Render(a.abstraction())
	return Pretty(output)
}

type ActionType interface {
	Action()
}

///////////////////////////////////////////
// Context

type Context struct {
	slackType BlockType
	elements  []flocksofblocks.Element

	blockId string

	optionals contextOptions
}

type contextOptions struct {
	BlockId bool
}

// NewContext creates a new context.
func NewContext() Context {
	return Context{
		slackType: ContextBlock,
		elements:  []flocksofblocks.Element{},
		optionals: contextOptions{
			BlockId: false,
		},
	}
}

// setBlockId sets the block id for the block.
func (c *Context) setBlockId(blockId string) {
	c.blockId = blockId
	c.optionals.BlockId = true
}

// removeBlockId removes the block id from the context.
func (c *Context) removeBlockId() {
	c.optionals.BlockId = false
}

// AddBlockId chain function to add block id to an existing context
func (c Context) AddBlockId(blockId string) Context {
	c.setBlockId(blockId)
	return c
}

// RemoveBlockId remove add block id from context
func (c Context) RemoveBlockId() Context {
	c.removeBlockId()
	return c
}

// addElement adds an element to the context.
func (c *Context) addElement(element flocksofblocks.Element) {
	c.elements = append(c.elements, element)
}

// AddElement chain function to add an element to an existing context
func (c Context) AddElement(element flocksofblocks.Element) Context {
	c.addElement(element)
	return c
}

// ContextAbstraction is the abstraction of the context block.
type ContextAbstraction struct {
	Type     string
	BlockId  string
	Elements []flocksofblocks.Element
}

// BlockRender is the implementation of the BlockRender interface.
func (c Context) BlockRender() {}

// BlockRenderAbstraction is the implementation of the BlockRenderAbstraction interface.
func (c Context) abstraction() ContextAbstraction {
	return ContextAbstraction{
		Type:     c.slackType.String(),
		BlockId:  c.blockId,
		Elements: c.elements,
	}
}

// Template is the template for the context block.
func (c ContextAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"elements": [{{range $index, $element := .Elements}}{{$element.Render}}{{if not $last}},{{end}}{{end}}]

{{if .Optionals.BlockId}},
	"block_id": "{{.BlockId}}"
{{end}}
	}`
}

///////////////////////////////////////////
// Divider

type Divider struct {
	slackType BlockType // required

	blockId string // optional

	optionals dividerOptionals
}

func (d Divider) BlockRender() {}

func NewDividerBlock() Divider {
	divider := Divider{
		slackType: DividerBlock,
		optionals: dividerOptionals{
			BlockId: false,
		},
	}
	return divider
}

// SetBlockId sets the block id for the block.
func (d *Divider) setBlockId(blockId string) {
	d.blockId = blockId
	d.optionals.BlockId = true
}

func (d *Divider) removeBlockId() {
	d.blockId = ""
	d.optionals.BlockId = false
}

// Render renders the block to a string.
func (d Divider) Render() string {
	return Render(d.abstraction())
}

// SlackType return slack type
func (d Divider) Type() string {
	return string(d.slackType)
}

func (d dividerAbstraction) Template() string {
	return `{
	"type": "{{.Type}}"
	{{if .Optionals.BlockId}},"block_id": "{{.BlockId}}"{{end}}
}`
}

type dividerAbstraction struct {
	Type      string
	BlockId   string
	Optionals dividerOptionals
}

// create an abstraction of the divider block
func (d Divider) abstraction() dividerAbstraction {
	return dividerAbstraction{
		Type:      d.Type(),
		BlockId:   d.blockId,
		Optionals: d.optionals,
	}
}

type dividerOptionals struct {
	BlockId bool
}

// SetBlockId sets the block id for the block.
func (d Divider) SetBlockId(blockId string) Divider {
	d.setBlockId(blockId)
	return d
}

// RemoveBlockId removes the block id from the block.
func (d Divider) RemoveBlockId() Divider {
	d.blockId = ""
	d.optionals.BlockId = false
	return d
}

///////////////////////////////////////////
// File

type File struct {
	slackType  BlockType
	externalId string
	source     string
	blockId    string

	optional fileOptions
}

type fileOptions struct {
	BlockId bool
}

// NewFile creates a new file.
func NewFile(externalId string, source string) File {
	return File{
		slackType:  FileBlock,
		externalId: externalId,
		source:     source,
		optional: fileOptions{
			BlockId: false,
		},
	}
}

// SetBlockId sets the block id for the block.
func (f *File) setBlockId(blockId string) {
	f.blockId = blockId
	f.optional.BlockId = true
}

func (f *File) removeBlockId() {
	f.blockId = ""
	f.optional.BlockId = false
}

// AddBlockId chain function to add block id to an existing file
func (f File) AddBlockId(blockId string) File {
	f.setBlockId(blockId)
	return f
}

// RemoveBlockId remove add block id from file
func (f File) RemoveBlockId() File {
	f.removeBlockId()
	return f
}

// BlockRender is a dummy function to satisfy the Block interface.
func (f File) BlockRender() {}

// fileAbstraction is a helper struct to generate the abstraction for the file.
type fileAbstraction struct {
	Type       string
	ExternalId string
	Source     string
	BlockId    string
	Optional   fileOptions
}

// abstraction is a helper function to generate the abstraction for the file.
func (f File) abstraction() fileAbstraction {
	return fileAbstraction{
		Type:       string(f.slackType),
		ExternalId: f.externalId,
		Source:     f.source,
		BlockId:    f.blockId,

		Optional: f.optional,
	}
}

// template is a helper function to generate the template for the file.
func (f fileAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"external_id": "{{.ExternalId}}",
	"source": "{{.Source}}"{{if .Optional.BlockId}},
	"block_id": "{{.BlockId}}"{{end}}
}`
}

// render is a helper function to generate the json for the file.
func (f File) Render() string {
	return Render(f.abstraction())
}

///////////////////////////////////////////
// Header

type Header struct {
	slackType BlockType
	text      flocksofblocks.CompositionText

	blockId string

	optional headerOptions
}

type headerOptions struct {
	BlockId bool
}

// NewHeader creates a new header.
func NewHeader(text string) Header {
	return Header{
		slackType: HeaderBlock,
		text:      flocksofblocks.NewPlainText(text),
		optional: headerOptions{
			BlockId: false,
		},
	}
}

// SetBlockId sets the block id for the block.
func (h *Header) setBlockId(blockId string) {
	h.blockId = blockId
	h.optional.BlockId = true
}

func (h *Header) removeBlockId() {
	h.blockId = ""
	h.optional.BlockId = false
}

// AddBlockId chain function to add block id to an existing header
func (h Header) AddBlockId(blockId string) Header {
	h.setBlockId(blockId)
	return h
}

// RemoveBlockId remove add block id from header
func (h Header) RemoveBlockId() Header {
	h.removeBlockId()
	return h
}

// abstraction is a helper struct to generate the abstraction for the header.
type headerAbstraction struct {
	Type     string
	Text     flocksofblocks.CompositionText
	BlockId  string
	Optional headerOptions
}

// generate abstraction from header
func (h Header) abstraction() headerAbstraction {
	return headerAbstraction{
		Type:    string(h.slackType),
		Text:    h.text,
		BlockId: h.blockId,

		Optional: h.optional,
	}
}

// template is a helper struct to generate the template for the header.
func (f headerAbstraction) Template() string {
	return `{
		"type": "{{.Type}}",
		"text": {{.Text.Render}}{{if .Optional.BlockId}},
		"block_id": "{{.BlockId}}"
		{{end}}
	}`
}

// render is a helper function to generate the json for the file.
func (h Header) Render() string {
	return Render(h.abstraction())
}

// BlockRender is a dummy function to satisfy the Block interface.

type Image struct {
	slackType BlockType
	imageUrl  *url.URL // todo: make this a string set using a url
	altText   string

	title   flocksofblocks.CompositionText
	blockId string

	optional imageOptions
}

type imageOptions struct {
	Title   bool
	BlockId bool
}

// NewImage creates a new image.
func NewImage(imageUrl *url.URL, altText string) Image {
	return Image{
		slackType: ImageBlock,
		imageUrl:  imageUrl,
		altText:   altText,
		optional: imageOptions{
			Title:   false,
			BlockId: false,
		},
	}
}

// SetTitle sets the title for the image.
func (i *Image) setTitle(title string) {
	i.title = flocksofblocks.NewPlainText(title)
	i.optional.Title = true
}

func (i *Image) removeTitle() {
	i.optional.Title = false
}

// AddTitle chain function to add title to an existing image
func (i Image) AddTitle(title string) Image {
	i.setTitle(title)
	return i
}

// RemoveTitle remove add title from image
func (i Image) RemoveTitle() Image {
	i.removeTitle()
	return i
}

// SetBlockId sets the block id for the block.
func (i *Image) setBlockId(blockId string) {
	i.blockId = blockId
	i.optional.BlockId = true
}

func (i *Image) removeBlockId() {
	i.optional.BlockId = false
}

// AddBlockId chain function to add block id to an existing image
func (i Image) AddBlockId(blockId string) Image {
	i.setBlockId(blockId)
	return i
}

// RemoveBlockId remove add block id from image
func (i Image) RemoveBlockId() Image {
	i.removeBlockId()
	return i
}

// abstraction is a helper struct to generate the abstraction for the image.
type imageAbstraction struct {
	Type     string
	ImageUrl string
	AltText  string
	Title    flocksofblocks.CompositionText
	BlockId  string
	Optional imageOptions
}

// generate abstraction from image
func (i Image) abstraction() imageAbstraction {
	return imageAbstraction{
		Type:     i.slackType.String(),
		ImageUrl: i.imageUrl.String(),
		AltText:  i.altText,
		Title:    i.title,
		BlockId:  i.blockId,
		Optional: i.optional,
	}
}

// Template returns the template for the image.
func (i imageAbstraction) Template() string {
	return `{
		"type": "{{.Type}}",
		"image_url": "{{.ImageUrl}}",
		"alt_text": "{{.AltText}}"{{if .Optional.Title}},
		"title": {{.Title.Render}}{{end}}{{if .Optional.BlockId}},
		"block_id": "{{.BlockId}}"{{end}}
	}`
}

// Render renders the image to a string.
func (i Image) Render() string {
	return Render(i.abstraction())
}

///////////////////////////////////////////
// Input

type Input struct {
	slackType BlockType
	label     flocksofblocks.CompositionText
	element   flocksofblocks.InputElement

	dispatchAction bool
	blockID        string
	hint           flocksofblocks.CompositionText
	slackOptional  bool

	optionals inputOptional
}

type inputOptional struct {
	DispatchAction bool
	BlockID        bool
	Hint           bool
	SlackOptional  bool
}

func NewInput(label string, element flocksofblocks.InputElement) Input {
	return Input{
		slackType: InputBlock,
		label:     flocksofblocks.NewPlainText(label),
		element:   element,
	}
}

// set dispatch action
func (i *Input) setDispatchAction(dispatchAction bool) {
	i.dispatchAction = dispatchAction
	i.optionals.DispatchAction = true
}

// set block id
func (i *Input) setBlockID(blockID string) {
	i.blockID = blockID
	i.optionals.BlockID = true
}

// remove block id
func (i *Input) removeBlockID() {
	i.optionals.BlockID = false
}

// setHint
func (i *Input) setHint(s string) {
	i.hint = flocksofblocks.NewPlainText(s)
	i.optionals.Hint = true
}

// removeHint remove hint
func (i *Input) removeHint() {
	i.optionals.Hint = false
}

// set label
func (i *Input) setLabel(label string) {
	i.label = flocksofblocks.NewPlainText(label)
}

// SetSlackOptional set slack optional
func (i *Input) setSlackOptional() {
	i.slackOptional = true
	i.optionals.SlackOptional = true
}

// RemoveSlackOptional remove slack optional
func (i *Input) removeSlackOptional() {
	i.optionals.SlackOptional = false
}

type abstractionInput struct {
	Type           string
	Label          flocksofblocks.CompositionText
	Element        flocksofblocks.InputElement
	DispatchAction bool
	BlockID        string
	Hint           flocksofblocks.CompositionText
	SlackOptional  bool
	Optionals      inputOptional
}

// create abstraction input
func (i Input) abstraction() abstractionInput {
	return abstractionInput{
		Type:           i.slackType.String(),
		Label:          i.label,
		Element:        i.element,
		DispatchAction: i.dispatchAction,
		BlockID:        i.blockID,
		Hint:           i.hint,
		SlackOptional:  i.slackOptional,
		Optionals:      i.optionals,
	}
}

// Template for input
func (i abstractionInput) Template() string {
	return `{
"type": "{{.Type}}",
"label": {{.Label.Render}},
	
"element": {{.Element.Render}}

{{if .Optionals.DispatchAction}},
		"dispatch_action": "{{.DispatchAction}}"
{{end}}

{{if .Optionals.BlockID}},
	"block_id": "{{.BlockID}}"
{{end}}

{{if .Optionals.Hint}},
	"hint": {{.Hint.Render}}
{{end}}

{{if .Optionals.SlackOptional}},
	"optional": "{{.SlackOptional}}"
{{end}}
	}`
}

// Render render input
func (i Input) Render() string {
	raw := Render(i.abstraction())
	return Pretty(raw)
}

type InputType interface {
	Input()
}

func (i Input) SetLabel(label string) Input {
	i.setLabel(label)
	return i
}

// EnableDispatchAction bool as chain method
func (i Input) EnableDispatchAction() Input {
	i.setDispatchAction(true)
	return i
}

// DisableDispatchAction bool as chain method
func (i Input) DisableDispatchAction() Input {
	i.setDispatchAction(false)
	return i
}

// SetBlockID BlockID string as chain method
func (i Input) SetBlockID(blockID string) Input {
	i.setBlockID(blockID)
	return i
}

// RemoveBlockID string as chain method
func (i Input) RemoveBlockID() Input {
	i.removeBlockID()
	return i
}

// SetHint as chain method
func (i Input) SetHint(s string) Input {
	i.setHint(s)
	return i
}

// RemoveHint as chain method
func (i Input) RemoveHint() Input {
	i.removeHint()
	return i
}

// SetSlackOptional as chain method
func (i Input) MakeOptional() Input {
	i.setSlackOptional()
	return i
}

func (i Input) RemoveOptional() Input {
	i.removeSlackOptional()
	return i
}

///////////////////////////////////////////
// Section

type Section struct {
	slackType BlockType
	text      flocksofblocks.CompositionText

	accessory flocksofblocks.Element
	blockId   string

	fields []flocksofblocks.CompositionText

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
		slackType: SectionBlock,
		text:      flocksofblocks.NewPlainText(text),
		optional: sectionOptions{
			Accessory: false,
			BlockId:   false,
			Field:     false,
		},
	}
}

// SetAccessory sets the accessory for the section.
func (s *Section) setAccessory(accessory flocksofblocks.Element) {
	s.accessory = accessory
	s.optional.Accessory = true
}

func (s *Section) removeAccessory() {
	s.optional.Accessory = false
}

// addAccessory adds an accessory to the section.
func (s Section) AddAccessory(accessory flocksofblocks.Element) Section {
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
	s.fields = append(s.fields, flocksofblocks.NewPlainText(field))
	s.optional.Field = true
}

func (s *Section) addMarkdownField(field string) {
	s.fields = append(s.fields, flocksofblocks.NewMrkdwnText(field))
	s.optional.Field = true
}

func (s *Section) removeField() {
	s.optional.Field = false
}

// abstraction for the section block
type sectionAbstraction struct {
	Type string
	Text flocksofblocks.CompositionText

	Accessory flocksofblocks.Element
	BlockId   string

	Fields []flocksofblocks.CompositionText

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
	output := Render(s.abstraction())
	return Pretty(output)
}

type SectionType interface {
	Section()
}

func (s Section) String() string {
	return s.Render()
}

// BlockRender
func (s Section) BlockRender() {}

///////////////////////////////////////////
// Video

type Video struct {
	slackType    BlockType
	title        flocksofblocks.CompositionText
	thumbnailUrl *url.URL
	videoUrl     *url.URL
	altText      string

	// optionals
	authorName   string
	providerName string

	description flocksofblocks.CompositionText

	providerIconUrl *url.URL
	titleUrl        *url.URL
	blockId         string

	// optional
	optional optionalVideo
}

type optionalVideo struct {
	AuthorName      bool
	ProviderName    bool
	Description     bool
	ProviderIconUrl bool
	TitleUrl        bool
	BlockId         bool
}

func NewVideo(title string, thumbnailUrl *url.URL, videoUrl *url.URL, altText string) Video {
	return Video{
		slackType: VideoBlock,
		title:     flocksofblocks.NewPlainText(title),

		thumbnailUrl: thumbnailUrl,
		videoUrl:     videoUrl,
		altText:      altText,

		optional: optionalVideo{
			AuthorName:      false,
			ProviderName:    false,
			Description:     false,
			ProviderIconUrl: false,
			TitleUrl:        false,
			BlockId:         false,
		},
	}
}

// setAuthorName
func (v *Video) setAuthorName(authorName string) {
	v.authorName = authorName
	v.optional.AuthorName = true
}

// RemoveAuthorName
func (v *Video) removeAuthorName() {
	v.optional.AuthorName = false
}

// setProviderName
func (v *Video) setProviderName(providerName string) {
	v.providerName = providerName
	v.optional.ProviderName = true
}

// removeProviderName
func (v *Video) removeProviderName() {
	v.optional.ProviderName = false
}

// setDescription
func (v *Video) setDescription(description string) {
	v.description = flocksofblocks.NewPlainText(description).EnableEmoji()
	v.optional.Description = true
}

// removeDescription
func (v *Video) removeDescription() {
	v.optional.Description = false
}

// setProviderIconUrl
func (v *Video) setProviderIconUrl(providerIconUrl *url.URL) {
	v.providerIconUrl = providerIconUrl
	v.optional.ProviderIconUrl = true
}

// removeProviderIconUrl
func (v *Video) removeProviderIconUrl() {
	v.optional.ProviderIconUrl = false
}

// setTitleUrl
func (v *Video) setTitleUrl(titleUrl *url.URL) {
	v.titleUrl = titleUrl
	v.optional.TitleUrl = true
}

// removeTitleUrl
func (v *Video) removeTitleUrl() {
	v.optional.TitleUrl = false
}

// setBlockId
func (v *Video) setBlockId(blockId string) {
	v.blockId = blockId
	v.optional.BlockId = true
}

// removeBlockId
func (v *Video) removeBlockId() {
	v.optional.BlockId = false
}

// abstraction structure
type abstractionVideo struct {
	Type         string
	Title        flocksofblocks.CompositionText
	ThumbnailUrl string
	VideoUrl     string
	AltText      string

	AuthorName      string
	ProviderName    string
	Description     flocksofblocks.CompositionText
	ProviderIconUrl string
	TitleUrl        string
	BlockId         string

	Optional optionalVideo
}

// abstract
func (v Video) abstraction() abstractionVideo {
	providerIconUrl := ""
	if v.optional.ProviderIconUrl {
		providerIconUrl = v.providerIconUrl.String()
	}

	titleUrl := ""
	if v.optional.TitleUrl {
		titleUrl = v.titleUrl.String()
	}

	return abstractionVideo{
		Type:         string(v.slackType),
		Title:        v.title,
		ThumbnailUrl: v.thumbnailUrl.String(),
		VideoUrl:     v.videoUrl.String(),
		AltText:      v.altText,
		AuthorName:   v.authorName,
		ProviderName: v.providerName,
		Description:  v.description,

		ProviderIconUrl: providerIconUrl,
		TitleUrl:        titleUrl,
		BlockId:         v.blockId,
		Optional:        v.optional,
	}
}

// AddAuthorName chain function to add author name to an existing video
func (v Video) AddAuthorName(authorName string) Video {
	v.setAuthorName(authorName)
	return v
}

// RemoveAuthorName remove add author name from video
func (v Video) RemoveAuthorName() Video {
	v.removeAuthorName()
	return v
}

// AddProviderName chain function to add provider name to an existing video
func (v Video) AddProviderName(providerName string) Video {
	v.setProviderName(providerName)
	return v
}

// RemoveProviderName remove add provider name from video
func (v Video) RemoveProviderName() Video {
	v.removeProviderName()
	return v
}

// AddDescription chain function to add description to an existing video
func (v Video) AddDescription(description string) Video {
	v.setDescription(description)
	return v
}

// RemoveDescription remove add description from video
func (v Video) RemoveDescription() Video {
	v.removeDescription()
	return v
}

// AddProviderIconUrl chain function to add provider icon url to an existing video
func (v Video) AddProviderIconUrl(providerIconUrl *url.URL) Video {
	v.setProviderIconUrl(providerIconUrl)
	return v
}

// RemoveProviderIconUrl remove add provider icon url from video
func (v Video) RemoveProviderIconUrl() Video {
	v.removeProviderIconUrl()
	return v
}

// AddTitleUrl chain function to add title url to an existing video
func (v Video) AddTitleUrl(titleUrl *url.URL) Video {
	v.setTitleUrl(titleUrl)
	return v
}

// RemoveTitleUrl remove add title url from video
func (v Video) RemoveTitleUrl() Video {
	v.removeTitleUrl()
	return v
}

// AddBlockId chain function to add block id to an existing video
func (v Video) AddBlockId(blockId string) Video {
	v.setBlockId(blockId)
	return v
}

// RemoveBlockId remove add block id from video
func (v Video) RemoveBlockId() Video {
	v.removeBlockId()
	return v
}

// Render the block
func (v Video) Render() string {
	return Render(v.abstraction())
}

// Template for the block
func (v abstractionVideo) Template() string {
	return `{
	"type": "{{.Type}}",
	"title": {{.Title.Render}},
	
	"thumbnail_url": "{{.ThumbnailUrl}}",
	"video_url": "{{.VideoUrl}}",
	"alt_text": "{{.AltText}}"

{{if .Optional.AuthorName}},
		"author_name": "{{.AuthorName}}"
{{end}}

{{if .Optional.ProviderName}},
		"provider_name": "{{.ProviderName}}"
{{end}}

{{if .Optional.Description}},
		"description": {{.Description.Render}}
{{end}}

{{if .Optional.ProviderIconUrl}},
		"provider_icon_url": "{{.ProviderIconUrl}}"
{{end}}

{{if .Optional.TitleUrl}},
		"title_url": "{{.TitleUrl}}"
{{end}}

{{if .Optional.BlockId}},
		"block_id": "{{.BlockId}}"
{{end}}
	}`
}
