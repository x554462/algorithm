package main

import (
	"container/list"
	"fmt"
)

func main() {
	for _, s := range []string{"())", "(((", "()", "()))(("} {
		result := minAddToMakeValid(s)
		fmt.Println(result)
		result = minAddToMakeValid2(s)
		fmt.Println(result)
	}
}

// 难度：中等  贪心算法/栈
// https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/
func minAddToMakeValid(S string) int {
	var left, right = 0, 0
	for _, v := range S {
		if v == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		} else {
			left++
		}
	}
	return left + right
}

func minAddToMakeValid2(S string) int {
	stack := list.New()
	for _, v := range S {
		elem := stack.Front()
		if v == ')' {
			if elem != nil && elem.Value.(rune) != v {
				stack.Remove(elem)
				continue
			}
		}
		stack.PushFront(v)
	}
	return stack.Len()
}
