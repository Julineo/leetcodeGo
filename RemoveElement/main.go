package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeElement([]int{3,2,2,3},3))
}

func removeElement(nums []int, val int) (k int) {
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return
}
