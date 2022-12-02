package multiselectmenu

import (
	"fmt"
	"testing"
)

func TestNewMultiSelectMenuWithUserList(t *testing.T) {
	menu := NewMultiSelectMenuWithUserList("actionId")

	menu = menu.AddInitialUser("Sarah P").AddInitialUser("user2")

	menu = menu.AddPlaceholder("placeholder")

	fmt.Println(menu.Section())

}
