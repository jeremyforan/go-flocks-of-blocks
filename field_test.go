package flocksofblocks

import (
	"fmt"
	"testing"
)

func TestNewFoo(t *testing.T) {
	t.Run("should create a new foo", func(t *testing.T) {
		foo := NewFoo()

		foo = foo.SetBlockId("unpside down")

		fmt.Println(foo)
	})
}
