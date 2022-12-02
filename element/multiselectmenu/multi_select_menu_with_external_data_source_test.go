package multiselectmenu

import (
	"fmt"
	"go-flocks-of-blocks/composition/option"
	"testing"
)

func TestNewMultiSelectMenuWithExternalDataSource(t *testing.T) {
	t.Run("NewMultiSelectMenuWithExternalDataSource", func(t *testing.T) {
		menu := NewMultiSelectMenuWithExternalDataSource("actionId")

		opt := option.NewOption("initial", "value-1")
		menu.AddInitialOption(opt)

		fmt.Println(menu.Section())
	})
}
