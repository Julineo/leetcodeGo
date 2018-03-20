package main

import (
	"fmt"
	"sort"
)

func main () {
	fmt.Println(subsets([]int{3,2,1}))
}

var solution [][]int

func subsets(ums []int) [][]int {
	solution = [][]int{}
	sort.Ints(ums)
	subsetsRec(ums, 0, []int{})
	return solution
}

func subsetsRec(ums []int, index int, tempSolution []int) {
	if index == len(ums) {

	} else {

	}

}
