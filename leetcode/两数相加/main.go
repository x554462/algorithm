package main

import "fmt"

func main() {
	var l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var l2 = &ListNode{5, &ListNode{6, &ListNode{6, nil}}}
	result := addTwoNumbers(l1, l2)
	for {
		fmt.Printf("%d -> ", result.Val)
		if result.Next == nil {
			break
		}
		result = result.Next
	}
	fmt.Println("nil")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 难度：中等  链表/数学
// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var arr []int
	var ln = &ListNode{}
	carry := 0
	for {
		value := 0
		if l1 != nil && l2 != nil {
			value = l1.Val + l2.Val + carry
			if value >= 10 {
				value -= 10
				carry = 1
			} else {
				carry = 0
			}
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			value = l1.Val + carry
			if value >= 10 {
				value -= 10
				carry = 1
			} else {
				carry = 0
			}
			l1 = l1.Next
		} else if l2 != nil {
			value = l2.Val + carry
			if value >= 10 {
				value -= 10
				carry = 1
			} else {
				carry = 0
			}
			l2 = l2.Next
		} else if carry == 1 {
			value = 1
			carry = 0
		} else {
			break
		}
		arr = append(arr, value)
	}
	head := ln
	l := len(arr) - 1
	for k, value := range arr {
		ln.Val = value
		if l == k {
			ln.Next = nil
		} else {
			ln.Next = &ListNode{}
			ln = ln.Next
		}

	}
	return head
}
