package main

import "fmt"

type MyLinkedList struct {
	val int
	next *MyLinkedList
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var list MyLinkedList
	return list
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	cur := this
	for i := 0; i <= index; i++ {
		if i == index {
			return cur.val
		}
		if cur.next == nil {
			break
		}
		cur = cur.next
	}
	return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	newHead := &MyLinkedList{val: val}
	fmt.Printf("1 newHead: %v, %v\n", &newHead, newHead)
	fmt.Printf("2 this: %v, %v\n", &this, this)
	tmp := this
	fmt.Printf("3 tmp: %v, %v\n", &tmp, tmp)
	this = newHead
	fmt.Printf("4 this: %v, %v\n", &this, this)
	newHead.next = tmp
	fmt.Printf("5 newHead: %v, %v\n", &newHead, newHead)
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	var newTail, cur *MyLinkedList
	newTail.val = val
	for {
		if cur.next == nil {
			cur.next = newTail
		}
		cur = cur.next
	}
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	var newL, cur, prev *MyLinkedList
	cur = this
	newL.val = val
	for i := 0; i <= index; i++ {
		if i == index {
			newL.next = cur
			prev.next = newL
		}
		prev = cur
		cur = cur.next
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    var prev *MyLinkedList
	cur := this
	for i := 0; i <= index; i++ {
		if i == index {
			prev.next = cur.next
		}
		prev = cur
		cur = cur.next
	}
}
