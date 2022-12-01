package button

import (
	"go-flocks-of-blocks/common"
	go_slack_blocks "go-flocks-of-blocks/composition/confirmationdialog"
	"net/url"
)

/////////////////////
// Chain Functions //
/////////////////////

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
	b.setStyle(common.StylePrimary)
	return b
}

// MakeStyleDanger invoke option sets the style of the button to primary.
func (b Button) MakeStyleDanger() Button {
	b.setStyle(common.StyleDanger)
	return b
}

// MakeStyleDefault invoke option sets the style of the button to primary.
func (b Button) MakeStyleDefault() Button {
	b.setStyle(common.StyleDefault)
	return b
}

// AddConfirmationDialog sets the confirmation dialog for the button.
func (b Button) AddConfirmationDialog(confirm go_slack_blocks.ConfirmationDialog) Button {
	b.setConfirmationDialog(confirm)
	return b
}

// RemoveConfirmationDialog removes the confirmation dialog from the button.
func (b Button) RemoveConfirmationDialog() Button {
	b.confirm = go_slack_blocks.ConfirmationDialog{}
	b.optionals.Confirm = false
	return b
}

// SetAccessibilityLabel sets the style for the button.
func (b Button) SetAccessibilityLabel(label string) Button {
	b.setAccessibilityLabel(label)
	return b
}

// RemoveAccessibilityLabel removes the style from the button.
func (b Button) RemoveAccessibilityLabel() Button {
	b.optionals.AccessibilityLabel = false
	return b
}
