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
	input = []int{10,9,8,9,10}

	//creating linked list
	ListNodes := make([]ListNode,len(input))
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
		ListNodes[i].Val = input[i]
		if i < len(input) - 1 {
			ListNodes[i].Next = &ListNodes[i+1]
		}
	}


	fmt.Println(isPalindrome(&ListNodes[0]))

}

func isPalindrome(head *ListNode) bool {
    slow, fast := head, head
    var prev *ListNode

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        next := slow.Next
        slow.Next = prev
        prev = slow
        slow = next
    }

    first, second := prev, slow
    if fast != nil { second = second.Next }

    for second != nil {
        if first.Val != second.Val {
            return false
        }
        first, second = first.Next, second.Next
    }
    return true
}
