package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"testing"
)

func TestDateTimePicker(t *testing.T) {
	t.Run("NewDateTimePicker", func(t *testing.T) {
		dateTimePicker := NewDateTimePicker("datetimepicker-action")
		output := dateTimePicker.Render()
		fmt.Println(common.Pretty(output))
	})
}
