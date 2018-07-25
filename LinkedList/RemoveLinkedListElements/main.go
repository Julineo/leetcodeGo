/*
https://leetcode.com/explore/learn/card/linked-list/219/classic-problems/1207/

Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
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
	l.AddAtTail(3)
	l.showList()
	l = removeElements(l, 3)
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

// removes all elements by given value
func removeElements(head *ListNode, val int) *ListNode {
	var cur, prev *ListNode
	cur = head
	for cur != nil {
		if cur.Val == val && cur == head {
			head = cur.Next
			cur = cur.Next
			continue
		}
		if cur.Val == val {
			cur = cur.Next
			prev.Next = cur
			continue
		}
		prev = cur
		cur = cur.Next
	}
	return head
}

