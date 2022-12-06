package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
	"testing"
)

func TestNewRadioButton(t *testing.T) {
	t.Run("NewRadioButton", func(t *testing.T) {
		rb := NewRadioButton("actionId").AddOption(composition.NewOption("option-1", "value-1")).AddOption(composition.NewOption("option-2", "value-a")).AddOption(composition.NewOption("option-3", "value-x"))
		output := rb.Render()
		fmt.Println(output)
	})
}
