package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	for _, v := range []*TreeNode{&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{6, &TreeNode{7, nil, nil}, &TreeNode{8, nil, nil}}}}} {
		result := postorderTraversal(v)
		fmt.Println(result)
		result = postorderTraversal2(v)
		fmt.Println(result)
	}
}

// 难度：困难  栈/树/哈希表
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *TreeNode) []int {
	var result []int
	recursion(root, &result)
	return result
}

func recursion(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	recursion(node.Left, vals)
	recursion(node.Right, vals)
	*vals = append(*vals, node.Val)
}

// 迭代：栈
func postorderTraversal2(root *TreeNode) []int {
	stack := list.New()
	node := root
	var result []int
	if node == nil {
		return result
	}
	stack.PushFront(node)
	for stack.Len() > 0 {
		elem := stack.Front()
		stack.Remove(elem)
		node := elem.Value.(*TreeNode)
		if node.Left != nil {
			stack.PushFront(node.Left)
		}
		if node.Right != nil {
			stack.PushFront(node.Right)
		}
		result = append([]int{node.Val}, result...)
	}
	return result
}
