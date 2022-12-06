package element

import (
	"fmt"
	"testing"
)

func TestNewSelectMenuWithUserList(t *testing.T) {
	menu := NewSelectMenuWithUserList("actionId")

	menu = menu.SetInitialUser("Sarah P")

	menu = menu.AddPlaceholder("placeholder")

	output := menu.Section().Render()

	fmt.Println(output)

}
