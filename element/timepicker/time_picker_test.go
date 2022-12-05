package timepicker

import "testing"

func TestNewTimePicker(t *testing.T) {
	t.Run("NewTimePicker", func(t *testing.T) {
		tp := NewTimePicker("actionId")
		output := tp.Render()
		t.Log(output)
	})
}
