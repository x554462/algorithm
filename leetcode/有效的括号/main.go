package main

import "fmt"

func main() {
	result := isValid("")
	fmt.Println(result)
}

type BracketsNode struct {
	Val  int32
	Next *BracketsNode
}

// 难度：简单  栈/字符串
// https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	var bn *BracketsNode
	var l int32
	for _, v := range s {
		if v == ')' {
			l = '('
		} else if v == ']' {
			l = '['
		} else if v == '}' {
			l = '{'
		} else {
			l = 0
		}
		if bn != nil && l == bn.Val {
			bn = bn.Next
		} else {
			temp := &BracketsNode{v, bn}
			bn = temp
		}
	}
	if bn == nil {
		return true
	}
	return false
}
