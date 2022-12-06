package block

import (
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	t.Run("valid header", func(t *testing.T) {
		header := NewHeader("header text")
		output := header.Render()
		fmt.Println("Header output: \n\n", output)
	})
}
