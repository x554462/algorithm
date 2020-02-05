package main

import "fmt"

func main() {
	for _, v := range []*TreeNode{&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{17, nil, nil}}}, nil} {
		result := maxDepth2(v)
		fmt.Println(result)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 难度：简单  树/深度优先遍历
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

// 递归解法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	}
	return rightMax + 1
}

// 遍历树解法
var max, depth = 0, 0

func maxDepth2(root *TreeNode) int {
	max, depth = 0, 0
	recursion(root)
	return max
}
func recursion(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node)
	depth++
	if depth > max {
		max = depth
	}
	if node.Left != nil {
		recursion(node.Left)
	}
	if node.Right != nil {
		recursion(node.Right)
	}
	depth--
}
