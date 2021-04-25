package list

import (
	"fmt"
	"testing"
)


func TestNewDoubleLinkedList(t *testing.T) {
	l := NewDoubleLinkedList()
	l.LPush([]byte("1"), []byte("2"))
	l.RPush([]byte("3"), []byte("4"))
	res, err := l.LPop()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res) == "2")

	rangeRes, err := l.LRange(0, -1)
	if err != nil {
		panic(err)
	}
	fmt.Println(rangeRes)
}
