package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	result := reverseKGroup(l, 2)
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

// 难度：困难  链表
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	var cur = head
	var curPos = 0
	var temp, lastNode *ListNode
	for {
		curPos++
		if curPos%k == 0 {
			lastNode = cur
			curPos++
			cur = cur.Next
			if cur.Next == nil {
				break
			}
		}
		if lastNode == nil {
			temp = cur.Next
			cur.Next = cur.Next.Next
			temp.Next = head
			head = temp
		} else {
			temp = cur.Next
			cur.Next = cur.Next.Next
			temp.Next = lastNode.Next
			lastNode.Next = temp
		}
		if cur.Next == nil {
			break
		}
	}
	if (curPos+1)%k != 0 {
		if lastNode != nil {
			cur = lastNode.Next
			for cur.Next != nil {
				temp = cur.Next
				cur.Next = cur.Next.Next
				temp.Next = lastNode.Next
				lastNode.Next = temp
			}
		} else {
			cur = head
			for cur.Next != nil {
				temp = cur.Next
				cur.Next = cur.Next.Next
				temp.Next = head
				head = temp
			}
		}
	}
	return head
}
