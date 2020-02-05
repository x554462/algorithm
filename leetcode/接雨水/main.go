package main

import "fmt"

func main() {
	result := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println(result)
}

// 难度：困难  栈/数组/双指针
// https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	l := len(height)
	if l <= 2 {
		return 0
	}
	var arr = make([]int, l)
	arr[0] = height[0]
	for i := 1; i < l; i++ {
		if height[i] < arr[i-1] {
			arr[i] = arr[i-1]
		} else {
			arr[i] = height[i]
		}
	}
	var result int
	arr[l-1] = height[l-1]
	for l -= 2; l > 0; l-- {
		if height[l] > arr[l+1] {
			arr[l] = height[l]
		} else if arr[l] > arr[l+1] {
			arr[l] = arr[l+1]
		}
		result += arr[l] - height[l]
	}
	return result
}
