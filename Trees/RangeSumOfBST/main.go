
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

type stack []*tree.Tree

func (s stack) Push(v *tree.Tree) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, *tree.Tree) {
    l := len(s)
    return  s[:l-1], s[l-1]
}

func main() {
	t := tree.New(1)
	s := make(stack, 0)
	s = s.Push(t)



	for i := 0; i < 10; i++ {

		fmt.Println(s)
		s, p := s.Pop()
		if p != nil {
			fmt.Println("p.Value", p.Value)
			fmt.Println("l:", p.Left, "r:", p.Right)
			s = s.Push(p.Left)
			s = s.Push(p.Right)
		}
	}
}

func rangeSumBST(root *tree.Tree, l int, r int) int {
	return 0
}
