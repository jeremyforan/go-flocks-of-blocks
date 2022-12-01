package confirmationdialog

import (
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/compositiontext"
)

const (
	confirmationDialogTitleLimit   = 100
	confirmationDialogTextLimit    = 30
	confirmationDialogConfirmLimit = 30
	confirmationDialogDenyLimit    = 30
)

// todo make abstraction for this
type ConfirmationDialog struct {
	title compositiontext.CompositionText

	// text is a CompositionText object which can be either a PlainText or a MarkdownText
	text    compositiontext.CompositionText
	confirm compositiontext.CompositionText
	deny    compositiontext.CompositionText

	// todo: make a generic style type
	style common.ColorSchema
}

// NewConfirmationDialog creates a new confirmation dialog
// todo: might consider making better input names
func NewConfirmationDialog(title string, text string, confirm string, deny string) ConfirmationDialog {

	//// todo: should I truncate the compositiontext if it is too long?
	//if len(title) > confirmationDialogTitleLimit {
	//	title = title[:confirmationDialogTitleLimit]
	//}
	//
	//if len(confirm) > confirmationDialogConfirmLimit {
	//	confirm = confirm[:confirmationDialogConfirmLimit]
	//}
	//
	//if len(deny) > confirmationDialogDenyLimit {
	//	deny = deny[:confirmationDialogDenyLimit]
	//}
	//
	//if len(compositiontext.compositiontext) > confirmationDialogTextLimit {
	//	compositiontext.compositiontext = compositiontext.compositiontext[:confirmationDialogTextLimit]
	//}

	return ConfirmationDialog{
		title:   compositiontext.NewPlainText(title),
		text:    compositiontext.NewPlainText(text),
		confirm: compositiontext.NewPlainText(confirm),
		deny:    compositiontext.NewPlainText(deny),
	}
}

// confirmationDialogAbstraction is used to render the confirmation dialog
type confirmationDialogAbstraction struct {
	Title   compositiontext.CompositionText
	Text    compositiontext.CompositionText
	Confirm compositiontext.CompositionText
	Deny    compositiontext.CompositionText
	Style   common.ColorSchema
}

// create an abstraction for the template
func (c ConfirmationDialog) abstraction() confirmationDialogAbstraction {
	return confirmationDialogAbstraction{
		Title:   c.title,
		Text:    c.text,
		Confirm: c.confirm,
		Deny:    c.deny,
		Style:   c.style,
	}
}

// create the template
func (c confirmationDialogAbstraction) Template() string {
	return `{
	"title": {{.title.Render}},
	"compositiontext": {{.compositiontext.Render}},
	"confirm": {{.confirm.Render}},
	"deny": {{.deny.Render}},
	"style": {{.style}}
}`
}

// Render the template
func (c ConfirmationDialog) Render() string {
	return common.Render(c.abstraction())
}
