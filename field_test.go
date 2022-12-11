package flocksofblocks

import (
	"fmt"
	"testing"
)

//func TestBlockComposition(t *testing.T) {
//	t.Run("should render a block", func(t *testing.T) {
//		block := blockId{}
//
//		block.setBlockId("hello")
//
//		fmt.Println("Helllo worldddd")
//
//		block.setBlockId("gobye")
//
//		fmt.Println("Helllo worldddd")
//	})
//}

func TestNewFoo(t *testing.T) {
	t.Run("should create a new foo", func(t *testing.T) {
		foo := NewFoo()

		foo = foo.SetBlockId("unpside down")

		fmt.Println(foo)
	})
}
