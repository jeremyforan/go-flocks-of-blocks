package checkboxes

import (
	"go-flocks-of-blocks/composition/confirmationdialog"
	"go-flocks-of-blocks/composition/option"
)

// AddOption add option to checkboxes
func (c Checkboxes) AddOption(option option.Option) Checkboxes {
	c.addOption(option)
	return c
}

// AddInitialOption add initial option to checkboxes
func (c Checkboxes) AddInitialOption(option option.Option) Checkboxes {
	c.addInitialOption(option)
	c.addOption(option)
	return c
}

// AddConfirmationDialog add confirmation dialog to checkboxes
func (c Checkboxes) AddConfirmationDialog(confirmationDialog confirmationdialog.ConfirmationDialog) Checkboxes {
	c.addConfirmationDialog(confirmationDialog)
	return c
}

// RemoveConfirmationDialog remove confirmation dialog from checkboxes
func (c Checkboxes) RemoveConfirmationDialog() Checkboxes {
	c.removeConfirmationDialog()
	return c
}

// FocusOnLoad set focus on load to checkboxes
func (c Checkboxes) FocusOnLoad() Checkboxes {
	c.setFocusOnLoad(true)
	return c
}

// DisableFocusOnLoad remove focus on load from checkboxes
func (c Checkboxes) DisableFocusOnLoad() Checkboxes {
	c.setFocusOnLoad(false)
	return c
}
