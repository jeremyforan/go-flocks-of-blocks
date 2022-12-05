package datepicker

import (
	"github.com/jeremyforan/go-flocks-of-blocks/composition/confirmationdialog"
	"time"
)

// AddInitialDate chain function to add initial date to an existing date picker
func (d DatePicker) AddInitialDate(initialDate time.Time) DatePicker {
	d.setInitialDate(initialDate)
	return d
}

// RemoveInitialDate remove add initial date from date picker
func (d DatePicker) RemoveInitialDate() DatePicker {
	d.removeInitialDate()
	return d
}

// AddConfirm chain function to add confirm to an existing date picker
func (d DatePicker) AddConfirm(confirm confirmationdialog.ConfirmationDialog) DatePicker {
	d.setConfirm(confirm)
	return d
}

// RemoveConfirm remove add confirm from date picker
func (d DatePicker) RemoveConfirm() DatePicker {
	d.removeConfirm()
	return d
}

// AddPlaceholder chain function to add placeholder to an existing date picker
func (d DatePicker) AddPlaceholder(placeholder string) DatePicker {
	d.setPlaceholder(placeholder)
	return d
}

// RemovePlaceholder remove add placeholder from date picker
func (d DatePicker) RemovePlaceholder() DatePicker {
	d.removePlaceholder()
	return d
}

// MakeFocused chain function to add focus to an existing date picker
func (d DatePicker) MakeFocused() DatePicker {
	d.setFocus(true)
	return d
}

// RemoveInitialFocus remove add focus from date picker
func (d DatePicker) RemoveInitialFocus() DatePicker {
	d.setFocus(false)
	return d
}
