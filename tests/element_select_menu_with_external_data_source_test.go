package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
	"testing"
)

// todo: add massive amount of validation

func TestNewSelectMenuWithExternalDataSource(t *testing.T) {
	t.Run("NewSelectMenuWithExternalDataSource", func(t *testing.T) {
		menu := NewSelectMenuWithExternalDataSource("actionId")

		opt := composition.NewOption("initial", "value-1")
		menu = menu.AddInitialOption(opt).AddPlaceholder("placeholder").SetMinQueryLength(6)

		fmt.Println(menu.Section())
	})
}
