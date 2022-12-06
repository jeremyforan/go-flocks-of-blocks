package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestNewMultiSelectMenuWithConversationsList(t *testing.T) {
	t.Run("NewMultiSelectMenuWithConversationsList", func(t *testing.T) {
		menu := NewMultiSelectMenuWithConversationsList("actionID")
		confirm := flocksofblocks.NewConfirmationDialog("title", "text", "confirm", "deny")
		menu = menu.AddPlaceholder("placeholder").AddConfirmDialog(confirm)
		fmt.Println(menu.Section().Render())
	})
}
