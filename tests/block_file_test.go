package block

import (
	"fmt"
	"testing"
)

func TestFileRender(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		file := NewFile("externalId", "source")
		file = file.AddBlockId("file1")
		output := file.Render()
		fmt.Println("File output: \n\n", output)
	})

}
