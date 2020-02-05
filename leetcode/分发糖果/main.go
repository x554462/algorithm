package main

import "fmt"

func main() {
	result := candy([]int{1, 0, 2})
	fmt.Println(result)
}

// 难度：困难  贪心算法
// https://leetcode-cn.com/problems/candy/
func candy(ratings []int) int {
	result := 0
	l := len(ratings)
	var nums = make([]int, l)
	for k, v := range ratings {
		if k == 0 {
			nums[0] = 1
		} else {
			if v > ratings[k-1] {
				nums[k] = nums[k-1] + 1
			} else {
				nums[k] = 1
			}
		}
	}
	result += nums[l-1]
	l--
	for l > 0 {
		if ratings[l] < ratings[l-1] && nums[l] >= nums[l-1] {
			nums[l-1] = nums[l] + 1
		}
		result += nums[l-1]
		l--
	}
	return result
}
