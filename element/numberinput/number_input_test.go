package numberinput

import (
	"fmt"
	"testing"
)

func TestNewNumberInput(t *testing.T) {
	t.Run("NewNumberInput", func(t *testing.T) {
		number := NewNumberInput("actionId").FocusOnLoad().DecimalAllowed().MinValue(1).MaxValue(10).InitialValue(5).Placeholder("placeholder")

		output := number.Input("Label").Render()
		fmt.Println(output)
	})
}
