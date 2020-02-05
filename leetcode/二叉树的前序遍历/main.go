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
	for _, v := range []*TreeNode{&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, nil}}} {
		result := preorderTraversal(v)
		fmt.Println(result)
		result = preorderTraversal2(v)
		fmt.Println(result)
	}
}

// 难度：中等  栈/树/哈希表
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	var result []int
	recursion(root, &result)
	return result
}

func recursion(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	*vals = append(*vals, node.Val)
	recursion(node.Left, vals)
	recursion(node.Right, vals)
}

// 迭代：栈
func preorderTraversal2(root *TreeNode) []int {
	stack := list.New()
	node := root
	var result []int
	for stack.Len() > 0 || node != nil {
		for node != nil {
			stack.PushFront(node)
			result = append(result, node.Val)
			node = node.Left
		}
		elem := stack.Front()
		stack.Remove(elem)
		node = elem.Value.(*TreeNode)
		node = node.Right
	}
	return result
}
