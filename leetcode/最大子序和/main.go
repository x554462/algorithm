package main

import "fmt"

func main() {
	for _, v := range [][]int{[]int{-2, 1, -3, 4, -1, 2, -1, 1, -5, 4}} {
		result := maxSubArray(v)
		fmt.Println(result)
	}
}

// 难度：简单  数组/分治算法/动态规划
// https://leetcode-cn.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	var max, currentMax = 0, 0
	for k, v := range nums {
		if k == 0 {
			max = v
			currentMax = max
		} else {
			if currentMax >= 0 {
				currentMax += v
			} else {
				currentMax = v
			}
			if max < currentMax {
				max = currentMax
			}
		}
	}
	return max
}
