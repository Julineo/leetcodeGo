package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	//input := []int{9,4,3,4,5,7,4,2}
	input := []int{10,9,8,8,9,10}
	input = []int{10,9,8,9,10}

	//creating linked list
	ListNodes := make([]ListNode,len(input))
	for i := 0; i < len(input) - 1; i++ {
		ListNodes[i].Val = input[i]
		ListNodes[i].Next = &ListNodes[i+1]
	}

	isPalindrome(&ListNodes[0])
	isPalindrome(&ListNodes[1])
	isPalindrome(&ListNodes[2])
}

func isPalindrome(head *ListNode) bool {
	fmt.Println(head)
	return false
}
