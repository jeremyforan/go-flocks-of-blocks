package plaintextinput

import (
	"fmt"
	"testing"
)

func TestNewPlainTextInput(t *testing.T) {
	t.Run("NewPlainTextInput", func(t *testing.T) {
		pti := NewPlainTextInput("actionId").SetMaxLength(12).SetMinLength(3).FocusOnLoad()
		output := pti.Render()
		fmt.Println(output)
	})
}
