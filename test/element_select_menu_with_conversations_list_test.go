package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestNewSelectMenuWithConversationsList(t *testing.T) {
	t.Run("NewSelectMenuWithConversationsList", func(t *testing.T) {
		menu := NewSelectMenuWithConversationsList("actionID")
		confirm := flocksofblocks.NewConfirmationDialog("title", "text", "confirm", "deny")
		menu = menu.AddPlaceholder("placeholder").AddConfirmDialog(confirm)
		fmt.Println(menu.Section().Render())
	})
}
