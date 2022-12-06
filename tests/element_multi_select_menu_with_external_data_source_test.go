package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
	"testing"
)

// todo: add massive amount of validation

func TestNewMultiSelectMenuWithExternalDataSource(t *testing.T) {
	t.Run("NewMultiSelectMenuWithExternalDataSource", func(t *testing.T) {
		menu := NewMultiSelectMenuWithExternalDataSource("actionId")

		opt := composition.NewOption("initial", "value-1")
		menu = menu.AddInitialOption(opt).MaxSelectedItems(3).AddPlaceholder("Select items")

		fmt.Println(menu.Section())
	})
}
