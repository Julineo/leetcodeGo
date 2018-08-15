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
	input = []int{10,9,7,8,7,9,10}
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

	fmt.Println(ListNodes)
	fmt.Println(isPalindrome(&ListNodes[0]))
	fmt.Println(ListNodes)

}

//checks if linked list is palindrome
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//reverse first half of a list
	slow, fast := head, head
	var prev, next *ListNode

	for fast != nil && fast.Next != nil {
		fmt.Println("loop")
		fast = fast.Next.Next
		next = slow.Next
		slow.Next = prev
		prev = slow
		slow = next
		fmt.Printf("fast: %p %v\n", fast, fast)
		fmt.Printf("next: %p %v\n", next, next)
		fmt.Printf("slow: %p %v\n", slow, slow)
		fmt.Printf("prev: %p %v\n", prev, prev)
	}

	//getting ready for comperison, assuming odd and even cases
	var first, second *ListNode
	if fast == nil {
		first = prev
		second = next
	} else {
		first = prev
		second = next.Next
	}

	//compare two parts of a list
	for first != nil {
		fmt.Printf("first: %p %v\n", first, first)
		fmt.Printf("second: %p %v\n", second, second)
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}
