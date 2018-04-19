package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nodes := read()

	for i, node := range(nodes) {
		fmt.Printf("%p\n", &nodes[i])
		printNode(&node)
	}

	//passing root node
	fmt.Println(nodes)
	fmt.Println(maxDepth(&nodes[len(nodes) - 1]))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dl := 1 + maxDepth(root.Left)
	dr := 1 + maxDepth(root.Right)
	if dl > dr {
		return dl
	} else {
		return dr
	}
}

func read() []TreeNode {
	var input = []struct {
		val   int
		left  int
		right int
	}{
		{15, -1, -1},
		{7, -1, -1},
		{9, -1, -1},
		{20, 0, 1},
		{3, 2, 3},
	}

	var nodes []TreeNode = make([]TreeNode, len(input))

	var val, indexLeft, indexRight int
	for i := 0; i < len(input); i++ {
		val, indexLeft, indexRight = input[i].val, input[i].left, input[i].right
		nodes[i].Value = val
		if indexLeft >= 0 {
			nodes[i].Left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].Right = &nodes[indexRight]
		}
	}

	return nodes
}

func printNode(n *TreeNode) {
	fmt.Print("Value: ", n.Value)
	if n.Left != nil {
		fmt.Print(" Left: ", n.Left.Value)
	}
	if n.Right != nil {
		fmt.Print(" Right: ", n.Right.Value)
	}
	fmt.Println()
}
