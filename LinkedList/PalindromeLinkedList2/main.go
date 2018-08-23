package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := []int{9, 4, 3, 4, 5, 7, 4, 2}
	input = []int{10, 9, 8, 8, 9, 10}
	input = []int{10,9,7,5,7,9,10}
	input = []int{1,2}

	//creating linked list
	ListNodes := make([]ListNode, len(input))
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
		ListNodes[i].Val = input[i]
		if i < len(input)-1 {
			ListNodes[i].Next = &ListNodes[i+1]
		}
		fmt.Printf("node: %p %v\n", &ListNodes[i], ListNodes[i])
	}

	fmt.Println(isPalindrome(&ListNodes[0]))
}

// checks if linked list is palindrome
func isPalindrome(head *ListNode) bool {
	// empty or single LL
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head
	var prev, next *ListNode

	// reverse first half
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next = slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	fmt.Println(&prev, prev)
	fmt.Println(&next, next)

	if fast != nil {
		next = next.Next
	}



	// check if palindrome
	for prev != nil {
		fmt.Println(&prev, prev)
		fmt.Println(&next, next)
		if prev.Val != next.Val { return false }
		prev = prev.Next
		next = next.Next
	}
	return true
}
