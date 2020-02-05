package main

import "fmt"

func main() {
	for _, v := range [][]int{{9, 2, 7, 11, 15}} {
		result := twoSum(v[1:], v[0])
		fmt.Println(result)
		result = twoSum2(v[1:], v[0])
		fmt.Println(result)
	}
}

// 难度：简单  数组/哈希表
// https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	var result = make([]int, 2)
	l := len(nums)
	for k, v := range nums {
		if k < l-1 {
			for kk, vv := range nums[k+1:] {
				if v+vv == target {
					result[0] = k
					result[1] = kk + k + 1
				}
			}
		}
	}
	return result
}

func twoSum2(nums []int, target int) []int {
	var result = make([]int, 2)
	var m = map[int]int{}
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		} else {
			m[num] = i
		}
	}
	return result
}
