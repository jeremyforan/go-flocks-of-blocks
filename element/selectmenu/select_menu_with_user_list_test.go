package selectmenu

import (
	"fmt"
	"testing"
)

func TestNewMultiSelectMenuWithUserList(t *testing.T) {
	menu := NewSelectMenuWithUserList("actionId")

	menu = menu.SetInitialUser("Sarah P")

	menu = menu.AddPlaceholder("placeholder")

	output := menu.Section().Render()

	fmt.Println(output)

}
