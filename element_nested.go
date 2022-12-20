package flocksofblocks

import "net/url"

// Button as defined in slack
type Button struct {
	slackType ElementType
	text      CompositionText
	actionId

	// optionals
	url     string
	value   string
	style   ColorSchema
	confirm ConfirmationDialog
	accessibilityLabel

	// optionals help with the template rendering
	optionals buttonOptionals
}

// NewButton create a new button element for an action.
func NewButton(text string, actionId string) Button {
	button := Button{
		slackType: ButtonElement,
		text:      NewPlainText(text),
		optionals: buttonOptionals{
			Url:                false,
			Value:              false,
			Style:              false,
			Confirm:            false,
			AccessibilityLabel: false,
		},
	}
	button.actionId.SetValue(actionId)

	return button
}

type buttonOptionals struct {
	Url                bool
	Value              bool
	Style              bool
	Confirm            bool
	AccessibilityLabel bool
}

type buttonAbstraction struct {
	Type     string          // required
	Text     CompositionText // required
	ActionId string          // required

	// optionals
	Url                string
	Value              string
	Style              string
	Confirm            ConfirmationDialog
	AccessibilityLabel string

	// optionals help with the template rendering
	Optionals buttonOptionals
}

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

func (b Button) Section() Section {
	return NewSection().AddAccessory(b)
}

func (b Button) ElementRender() {}

////////////////////////////////////////////////////////////////////////////////////
// Button Abstraction

func (b *Button) abstraction() buttonAbstraction {
	url := ""
	if b.optionals.Url {
		url = b.url
	}
	return buttonAbstraction{
		Type:               b.slackType.String(),
		Text:               b.text,
		ActionId:           b.actionId.String(),
		Url:                url,
		Value:              b.value,
		Style:              b.style.String(),
		Confirm:            b.confirm,
		AccessibilityLabel: b.accessibilityLabel.String(),
		Optionals:          b.optionals,
	}
}

// AddUrl chain function to add url to an existing button
func (b Button) AddUrl(url *url.URL) Button {
	b.setUrl(url)
	return b
}

// RemoveUrl chain function to remove url from an existing button
func (b Button) RemoveUrl() Button {
	b.removeUrl()
	return b
}

// SetValue sets the value for the button.
func (b Button) SetValue(value string) Button {
	b.setValue(value)
	return b
}

func (b Button) RemoveValue() Button {
	b.removeValue()
	return b
}

// MakeStylePrimary chain method that sets the style of the button to primary.
func (b Button) MakeStylePrimary() Button {
	b.setStyle(StylePrimary)
	return b
}

// MakeStyleDanger invoke option sets the style of the button to primary.
func (b Button) MakeStyleDanger() Button {
	b.setStyle(StyleDanger)
	return b
}

// MakeStyleDefault invoke option sets the style of the button to primary.
func (b Button) MakeStyleDefault() Button {
	b.setStyle(StyleDefault)
	return b
}

// AddConfirmationDialog sets the confirmation dialog for the button.
func (b Button) AddConfirmationDialog(confirm ConfirmationDialog) Button {
	b.setConfirmationDialog(confirm)
	return b
}

// RemoveConfirmationDialog removes the confirmation dialog from the button.
func (b Button) RemoveConfirmationDialog() Button {
	b.confirm = ConfirmationDialog{}
	b.optionals.Confirm = false
	return b
}

// SetAccessibilityLabel sets the style for the button.
func (b Button) SetAccessibilityLabel(label string) Button {
	b.accessibilityLabel.SetValue(label)
	return b
}

// RemoveAccessibilityLabel removes the style from the button.
func (b Button) RemoveAccessibilityLabel() Button {
	b.optionals.AccessibilityLabel = false
	return b
}

////////////
//func (b Button) Render() string {
//	return Render(b.abstraction())
//}
//
//// setUrl sets the url for the button.
//func (b *Button) setUrl(url *url.URL) {
//	b.url = url.String()
//	b.optionals.Url = true
//}
//
//// removeUrl removes the url from the button.
//func (b *Button) removeUrl() {
//	b.url = ""
//	b.optionals.Url = false
//}
//
//// setValue sets the value for the button.
//func (b *Button) setValue(value string) {
//	b.value = value
//	b.optionals.Value = true
//}
//
//func (b *Button) removeValue() {
//	b.value = ""
//	b.optionals.Value = true
//}
//
//func (b *Button) setStyle(style ColorSchema) {
//	if style == StyleDefault {
//		b.optionals.Style = false
//	} else {
//		b.style = style
//		b.optionals.Style = true
//	}
//}
//
//// setConfirmationDialog sets the confirmation dialog for the button.
//func (b *Button) setConfirmationDialog(confirm ConfirmationDialog) {
//	b.confirm = confirm
//}
//
//// setAccessibilityLabel sets the style for the button.
//func (b *Button) setAccessibilityLabel(label string) {
//	b.accessibilityLabel = label
//	b.optionals.AccessibilityLabel = true
//}
