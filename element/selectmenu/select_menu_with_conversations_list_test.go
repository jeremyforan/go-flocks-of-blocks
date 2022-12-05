package selectmenu

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/confirmationdialog"
	"testing"
)

func TestNewSelectMenuWithConversationsList(t *testing.T) {
	t.Run("NewSelectMenuWithConversationsList", func(t *testing.T) {
		menu := NewSelectMenuWithConversationsList("actionID")
		confirm := confirmationdialog.NewConfirmationDialog("title", "text", "confirm", "deny")
		menu = menu.AddPlaceholder("placeholder").AddConfirmDialog(confirm)
		fmt.Println(menu.Section().Render())
	})
}
