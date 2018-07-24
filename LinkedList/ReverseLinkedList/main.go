/*
https://leetcode.com/explore/learn/card/linked-list/219/classic-problems/1205/

Reverse Linked List
  Go to Discuss
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
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
	l.showList()
	l = reverseListRecursive(l, nil)
	l = nil
	l = reverseListRecursive(l, nil)
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

// reverses linked list iteratively
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur, head, next := head, head, head
	for next != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	head = prev
	return head
}

func reverseListRecursive(curr, prev *ListNode) *ListNode {
	var head, next *ListNode
	if curr == nil {
		return nil
	}
	if curr.Next == nil {
		head = curr
		curr.Next = prev
		return head
	}
	next = curr.Next
	curr.Next = prev
	return reverseListRecursive(next, curr)
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

