package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"testing"
)

func TestNewDatePicker(t *testing.T) {
	t.Run("NewDatePicker", func(t *testing.T) {
		datePicker := NewDatePicker("datepicker-action")

		output := datePicker.Render()
		fmt.Println(common.Pretty(output))

	})
}
