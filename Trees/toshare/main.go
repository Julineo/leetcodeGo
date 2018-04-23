package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main () {
	nodes := read()
	fmt.Println(isValidBST(&nodes[len(nodes) - 1]))
	nodes = read2()
	fmt.Println(isValidBST(&nodes[len(nodes) - 1]))
}

var prev *TreeNode

func isValidBST(root *TreeNode) bool {
	prev = nil
	return isValidBSTHelper(root)
}
//need helper function to nill the prev value
func isValidBSTHelper(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBSTHelper(root.Left) {
		return false
	}
	if prev != nil && prev.Value >= root.Value {
		return false
	}
	prev = root
	return isValidBSTHelper(root.Right)
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
		{3, -1, -1},
		{6, -1, -1},
		{1, -1, -1},
		{4, 0, 1},
		{5, 2, 3},
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

func read2() []TreeNode {
	var input = []struct {
		val   int
		left  int
		right int
	}{
		{0, -1, -1},
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
