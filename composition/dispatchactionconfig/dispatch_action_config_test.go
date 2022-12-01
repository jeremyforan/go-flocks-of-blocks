package dispatchactionconfig

import (
	"fmt"
	"testing"
)

func TestNewDispatchActionConfig(t *testing.T) {
	t.Run("NewDispatchActionConfig", func(t *testing.T) {
		dispatchActionConfig := NewDispatchActionConfig().OnEnterPressed().OnCharacterEntered()

		output := dispatchActionConfig.Render()
		fmt.Println(output)
	})
}
