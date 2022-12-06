package flocksofblocks

import "testing"

func TestNewInputTest(t *testing.T) {
	t.Run("NewInputTest", func(t *testing.T) {
		input := NewInputTest("Block")

		output := input.Render()
		t.Log(output)
	}
}
