package flocksofblocks

import (
	"fmt"
	flocksofblocks "github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestFileRender(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		file := flocksofblocks.NewFile("externalId", "source")
		file = file.AddBlockId("file1")
		output := file.Render()
		fmt.Println("File output: \n\n", output)
	})

}
