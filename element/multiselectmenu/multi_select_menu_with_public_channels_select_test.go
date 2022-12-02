package multiselectmenu

import (
	"fmt"
	"testing"
)

func TestNewMultiSelectMenuWithPublicChannelsSelect(t *testing.T) {
	t.Run("NewMultiSelectMenuWithPublicChannelsSelect", func(t *testing.T) {
		menu := NewMultiSelectMenuWithPublicChannelsSelect("actionID")
		fmt.Println(menu.Section().Render())
	})
}
