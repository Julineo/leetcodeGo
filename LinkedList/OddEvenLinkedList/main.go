/*
https://leetcode.com/explore/learn/card/linked-list/219/classic-problems/1208/

Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:

Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL
Example 2:

Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL
Note:

The relative order inside both the even and odd groups should remain as it was in the input.
The first node is considered odd, the second node even and so on ...
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
	oddeven(l)
	l.showList()
}

// changes odd/even sequence of linked list
func oddeven(head *ListNode) *ListNode {
	if head == nil {return nil}
	var evenHead, odd, even *ListNode
	odd = head
	even = head.Next
	evenHead = even
	for {
		if even == nil { break }
		odd.Next = odd.Next.Next
		if even.Next == nil { break }
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// prints linked list
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

// reverses linked list recursively
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

