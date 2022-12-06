package composition

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestNewConfirmationDialog(t *testing.T) {
	t.Run("NewConfirmationDialog", func(t *testing.T) {
		confirm := NewConfirmationDialog("title", "text", "confirm", "deny")
		if confirm.Render() == "" {
			t.Error("failed to render")
		}

		confirm = confirm.SetStyle(flocksofblocks.StylePrimary)
		if confirm.Render() == "" {
			t.Error("failed to render")
		}
		fmt.Println(confirm.Render())
	})
}
