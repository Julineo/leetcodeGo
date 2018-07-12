/*
https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1296/

Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
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
	l = removeNthFromEnd(l, 5)
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

// removes nodes
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var slow, fast, prev *ListNode
	slow, fast = head, head
	// move fast n-times foreward
	for i := 1; i < n; i++ {
		fast = fast.Next
	}
	// move fast and slow until the end of list
	for {
		if fast.Next == nil {
			break
		}
		prev = slow
		fast = fast.Next
		slow = slow.Next
		fmt.Println(fast, slow, prev)
	}
	if prev == nil {
		return slow.Next
	}
	prev.Next = slow.Next
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

