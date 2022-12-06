package flocksofblocks

// todo make abstraction for this

type ConfirmationDialog struct {
	title CompositionText

	// text is a CompositionText object which can be either a PlainText or a MarkdownText
	text    CompositionText
	confirm CompositionText
	deny    CompositionText

	style ColorSchema

	optionals ConfirmationDialogOptions
}

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
{{if .Optional.Style}},	
	"style": "{{.Style}}"
{{end}}
}`
}

// Render the template
func (c ConfirmationDialog) Render() string {
	raw := Render(c.abstraction())
	return Pretty(raw)
}
