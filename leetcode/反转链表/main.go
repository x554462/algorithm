package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := reverseList2(l)
	if result == nil {
		return
	}
	for {
		fmt.Printf("%d -> ", result.Val)
		if result.Next == nil {
			break
		}
		result = result.Next
	}
	fmt.Println("nil")
}

// 难度：简单  链表
// https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var result *ListNode
	for {
		temp := &ListNode{head.Val, result}
		result = temp
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	return result
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}
