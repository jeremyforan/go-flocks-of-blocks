package confirmationdialog

import (
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/compositiontext"
)

// todo make abstraction for this

type ConfirmationDialog struct {
	title compositiontext.CompositionText

	// text is a CompositionText object which can be either a PlainText or a MarkdownText
	text    compositiontext.CompositionText
	confirm compositiontext.CompositionText
	deny    compositiontext.CompositionText

	style common.ColorSchema

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
		title:   compositiontext.NewPlainText(title),
		text:    compositiontext.NewPlainText(text),
		confirm: compositiontext.NewPlainText(confirm),
		deny:    compositiontext.NewPlainText(deny),
		optionals: ConfirmationDialogOptions{
			Style: false,
		},
	}
}

// set the style
func (c *ConfirmationDialog) setStyle(style common.ColorSchema) {
	c.style = style
	c.optionals.Style = true
}

// set the style public
func (c ConfirmationDialog) SetStyle(style common.ColorSchema) ConfirmationDialog {
	c.setStyle(style)
	return c
}

// confirmationDialogAbstraction is used to render the confirmation dialog
type confirmationDialogAbstraction struct {
	Title   compositiontext.CompositionText
	Text    compositiontext.CompositionText
	Confirm compositiontext.CompositionText
	Deny    compositiontext.CompositionText
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
	raw := common.Render(c.abstraction())
	return common.Pretty(raw)
}
