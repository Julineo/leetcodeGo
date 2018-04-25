package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	input := []int{9,4,3,4,5,7,4,2}
	input = []int{10,9,8,8,9,10}
	//input = []int{10,9,8,9,10}

	//creating linked list
	ListNodes := make([]ListNode,len(input))
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
		ListNodes[i].Val = input[i]
		if i < len(input) - 1 {
			ListNodes[i].Next = &ListNodes[i+1]
		}
		fmt.Printf("node: %p %v\n", &ListNodes[i], ListNodes[i])
	}

	fmt.Println(ListNodes)
	fmt.Println(isPalindrome(&ListNodes[0]))
	fmt.Println(ListNodes)

}

func isPalindrome(head *ListNode) bool {
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
	return false
}

