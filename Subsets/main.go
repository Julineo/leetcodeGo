package main

import (
	"fmt"
	"sort"
)

func main () {
	fmt.Println(subsets([]int{3,2,1}))
}

var solution [][]int
var tempSolution []int

func subsets(ums []int) [][]int {
	solution = [][]int{}
	sort.Ints(ums)
	subsetsRec(ums, 0)
	return solution
}

func subsetsRec(ums []int, index int) {
	if index == len(ums) {
		tmp := make([]int, len(tempSolution))//to avoid references to ums
		copy(tmp,tempSolution)
		solution = append(solution, tmp)
	} else {
		subsetsRec(ums, index + 1)
		tempSolution = append(tempSolution,ums[index])
		subsetsRec(ums, index + 1)
		if len(tempSolution) > 0 {
			tempSolution = tempSolution[:len(tempSolution)-1]
		}
	}

}
