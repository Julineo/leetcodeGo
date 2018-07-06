package main

import (
	"testing"
	"fmt"
)

func TestMyLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtTail(6)
	fmt.Println(list.Get(0))//0
	fmt.Println(list.Get(1))//-1
//	list.AddAtHead(8)
	fmt.Println(list.Get(0))
	fmt.Println(list)
}
