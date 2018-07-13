/*

*/
package main

import "fmt"

 type ListNode struct {
	Val int
	Next *ListNode
 }

func main() {
	var l *ListNode
	l = &ListNode{}
	l.Val = 1
	for i := 2; i <= 5; i++{
		l.AddAtTail(i)
	}
	l.showList()
	l = reverseList(l)
	fmt.Println("end")
	l.showList()
}

// prints tes linked list
func (this *ListNode) showList() {
	var cur *ListNode
	cur = this
	for cur != nil {
		fmt.Println(cur, &cur)
		cur = cur.Next
	}
}

// reverses linked list
func reverseList(head *ListNode) *ListNode {
	cur := head
	for {
		tmp := cur.Next
		head = cur.Next
		head.Next = cur
		cur.Next = tmp
		cur = tmp.Next
		break
	}
	return head
}

// adds nodes at tail
func (this *ListNode) AddAtTail(val int)  {
	var newTail, cur *ListNode
	newTail = &ListNode{val, nil}
	cur = this
	for {
		if cur.Next == nil {
			cur.Next = newTail
			return
		}
		cur = cur.Next
	}
}

