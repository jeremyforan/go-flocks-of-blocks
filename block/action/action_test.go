package action

import "testing"

func TestNewActionBlock(t *testing.T) {
	t.Run("NewActionBlock", func(t *testing.T) {
		action := NewAction("Block")

		output := action.Render()
		t.Log(output)
	})
}
