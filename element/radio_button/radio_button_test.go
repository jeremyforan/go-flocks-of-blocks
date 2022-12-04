package radio_button

import (
	"fmt"
	"go-flocks-of-blocks/composition/option"
	"testing"
)

func TestNewRadioButton(t *testing.T) {
	t.Run("NewRadioButton", func(t *testing.T) {
		rb := NewRadioButton("actionId").AddOption(option.NewOption("option-1", "value-1")).AddOption(option.NewOption("option-2", "value-a")).AddOption(option.NewOption("option-3", "value-x"))
		output := rb.Render()
		fmt.Println(output)
	})
}
