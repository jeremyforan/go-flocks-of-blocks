package selectmenu

import (
	"fmt"
	"testing"
)

func TestNewMultiSelectMenuWithPublicChannelsSelect(t *testing.T) {
	t.Run("NewSelectMenuWithPublicChannelsSelect", func(t *testing.T) {
		menu := NewSelectMenuWithPublicChannelsSelect("actionID").EnableResponseUrlEnabled().FocusOnLoad()
		fmt.Println(menu.Section().Render())
	})
}
