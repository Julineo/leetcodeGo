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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil {
        return t2 
    }
    if t2 == nil {
        return t1
    }
    t1.Val =+ t2.Val
    t1.Left = mergeTrees(t1.Left, t2.Left)
    t1.Right = mergeTrees(t1.Right, t2.Right)
    return t1
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	prev = nil
	var isValidBSTHelper func(root *TreeNode) bool
	isValidBSTHelper = func(root *TreeNode) bool {
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

	return isValidBSTHelper(root)
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
