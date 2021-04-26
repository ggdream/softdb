package bst

import (
	"fmt"
	"testing")


func TestNew(t *testing.T) {
	tree := New(0, 20)
	tree.Insert(-1, 50)
	tree.Insert(1, 30)
	tree.Insert(20, 40)
	tree.Insert(10, -20)
	if err := tree.Delete(1); err != nil {
		panic(err)
	}
	if err := tree.Update(-1, 30); err != nil {
		panic(err)
	}
	fmt.Println(tree.Search(-1))
}
