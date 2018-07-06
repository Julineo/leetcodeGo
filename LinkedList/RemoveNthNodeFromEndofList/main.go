package main

import "fmt"

 type ListNode struct {
	Val int
	Next *ListNode
 }

func main() {
	l := ListNode{}
	l.Val = 1
	fmt.Println(l, &l)
	for i := 2; i <= 5; i++{
		l.AddAtTail(i)
	}
	l.showList()
}

func (this *ListNode) showList() {
	var cur *ListNode
	cur = this
	for cur != nil {
		fmt.Println(cur, &cur)
		cur = cur.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return &ListNode{}
}

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

