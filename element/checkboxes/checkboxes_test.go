package checkboxes

import (
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/option"
	"fmt"
	"testing"
)

func TestNewCheckboxes(t *testing.T) {
	t.Run("NewCheckboxes", func(t *testing.T) {
		checkboxes := NewCheckboxes("that check check boom!!").AddInitialOption(option.NewOption("option1", "option1"))
		output := checkboxes.Render()
		fmt.Println(common.Pretty(output))
	})
}
