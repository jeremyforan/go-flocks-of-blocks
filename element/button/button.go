package button

import (
	common "github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/confirmationdialog"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
	"net/url"
)

const (
	textLengthLimit               = 75
	valueLengthLimit              = 2000
	urlLengthLimit                = 3000
	AccessibilityLabelLengthLimit = 75
)

// Button as defined in slack
type Button struct {
	slackType element.ElementType
	text      compositiontext.CompositionText
	actionId  string

	// optionals
	url                *url.URL
	value              string
	style              common.ColorSchema
	confirm            confirmationdialog.ConfirmationDialog
	accessibilityLabel string

	// optionals help with the template rendering
	optionals buttonOptionals
}

// NewButton create a new button element for an action.
func NewButton(text string, actionId string) Button {
	button := Button{
		slackType: element.Button,
		text:      compositiontext.NewPlainText(text),
		actionId:  actionId,
		optionals: buttonOptionals{
			Url:                false,
			Value:              false,
			Style:              false,
			Confirm:            false,
			AccessibilityLabel: false,
		},
	}

	return button
}

type buttonAbstraction struct {
	Type     string                          // required
	Text     compositiontext.CompositionText // required
	ActionId string                          // required

	// optionals
	Url                string
	Value              string
	Style              string
	Confirm            confirmationdialog.ConfirmationDialog
	AccessibilityLabel string

	// optionals help with the template rendering
	Optionals buttonOptionals
}

func (b Button) Render() string {
	return common.Render(b.abstraction())
}

// setUrl sets the url for the button.
func (b *Button) setUrl(url *url.URL) {
	b.url = url
	b.optionals.Url = true
}

// removeUrl removes the url from the button.
func (b *Button) removeUrl() {
	b.url = nil
	b.optionals.Url = false
}

// setValue sets the value for the button.
func (b *Button) setValue(value string) {
	b.value = value
	b.optionals.Value = true
}

func (b *Button) removeValue() {
	b.value = ""
	b.optionals.Value = true
}

func (b *Button) setStyle(style common.ColorSchema) {
	if style == common.StyleDefault {
		b.optionals.Style = false
	} else {
		b.style = style
		b.optionals.Style = true
	}
}

// setConfirmationDialog sets the confirmation dialog for the button.
func (b *Button) setConfirmationDialog(confirm confirmationdialog.ConfirmationDialog) {
	b.confirm = confirm
}

// setAccessibilityLabel sets the style for the button.
func (b *Button) setAccessibilityLabel(label string) {
	b.accessibilityLabel = label
	b.optionals.AccessibilityLabel = true
}

// todo: make primary / default / danger

// Template returns the template for the button.
func (b buttonAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"action_id": "{{.ActionId}}",
	"text": {{.Text.Render}}

{{if .Optionals.Url}},
	"url": "{{.Url}}"
{{end}}

{{if .Optionals.Value}},
	"value": "{{.Value}}"
{{end}}

{{if .Optionals.Style}},
	"style": "{{.Style}}"
{{end}}

{{if .Optionals.Confirm}},
	"confirm": {{.Confirm.Render}}
{{end}}

{{if .Optionals.AccessibilityLabel}},
	"accessibility_label": "{{.AccessibilityLabel}}"
{{end}}

}`
}

func (b Button) Section() {}
func (b Button) Action()  {}

func (b Button) ElementRender() {}
