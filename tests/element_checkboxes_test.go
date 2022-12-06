package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestNewCheckboxes(t *testing.T) {
	t.Run("NewCheckboxes", func(t *testing.T) {
		checkboxes := NewCheckboxes("that check check boom!!").AddInitialOption(flocksofblocks.NewOption("option1", "option1"))
		output := checkboxes.Render()
		fmt.Println(flocksofblocks.Pretty(output))
	})
}
