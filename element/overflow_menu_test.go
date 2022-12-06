package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
	"testing"
)

func TestNewOverflowMenu(t *testing.T) {
	menu := NewOverflowMenu("actionId")

	opt1 := composition.NewOption("option-1", "value-1")
	opt2 := composition.NewOption("option-2", "value-a")
	opt3 := composition.NewOption("option-3", "value-x")

	menu = menu.AddOption(opt1).AddOption(opt2).AddOption(opt3)

	output := menu.Render()

	fmt.Println(output)

}
