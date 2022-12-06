package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestNewOverflowMenu(t *testing.T) {
	menu := NewOverflowMenu("actionId")

	opt1 := flocksofblocks.NewOption("option-1", "value-1")
	opt2 := flocksofblocks.NewOption("option-2", "value-a")
	opt3 := flocksofblocks.NewOption("option-3", "value-x")

	menu = menu.AddOption(opt1).AddOption(opt2).AddOption(opt3)

	output := menu.Render()

	fmt.Println(output)

}
