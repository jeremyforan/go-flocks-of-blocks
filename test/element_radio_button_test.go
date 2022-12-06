package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestNewRadioButton(t *testing.T) {
	t.Run("NewRadioButton", func(t *testing.T) {
		rb := NewRadioButton("actionId").AddOption(flocksofblocks.NewOption("option-1", "value-1")).AddOption(flocksofblocks.NewOption("option-2", "value-a")).AddOption(flocksofblocks.NewOption("option-3", "value-x"))
		output := rb.Render()
		fmt.Println(output)
	})
}
