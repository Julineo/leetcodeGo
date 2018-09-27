package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	var recWalk func(t *tree.Tree, ch chan int)
	recWalk = func(t *tree.Tree, ch chan int) {
		if t != nil {
			recWalk(t.Left, ch)

			ch <- t.Value

			recWalk(t.Right, ch)
		}
	}

	recWalk(t, ch)

	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	return false
}

func main() {
//	fmt.Println(tree.New(1))
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
}
