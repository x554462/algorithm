package main

import "fmt"

func main() {
	for _, v := range []int{1, 2, 3, 4, 9, 10} {
		result := climbStairs(v)
		fmt.Println(result)

	}
}

// 难度：简单  动态规划
// https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	l1, l2 := 2, 1
	var result int
	for i := 3; i <= n; i++ {
		result = l1 + l2
		l2 = l1
		l1 = result
	}
	return result
}
