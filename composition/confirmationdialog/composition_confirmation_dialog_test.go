package confirmationdialog

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"testing"
)

func TestNewConfirmationDialog(t *testing.T) {
	t.Run("NewConfirmationDialog", func(t *testing.T) {
		confirm := NewConfirmationDialog("title", "text", "confirm", "deny")
		if confirm.Render() == "" {
			t.Error("failed to render")
		}

		confirm = confirm.SetStyle(common.StylePrimary)
		if confirm.Render() == "" {
			t.Error("failed to render")
		}
		fmt.Println(confirm.Render())
	})
}
