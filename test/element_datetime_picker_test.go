package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestDateTimePicker(t *testing.T) {
	t.Run("NewDateTimePicker", func(t *testing.T) {
		dateTimePicker := NewDateTimePicker("datetimepicker-action")
		output := dateTimePicker.Render()
		fmt.Println(flocksofblocks.Pretty(output))
	})
}
