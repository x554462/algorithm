package main

import (
	"fmt"
	"math"
)

func main() {
	for _, v := range []int{123, -123, 120, 0, 9223372036854775719, 922337203685477, 2147483312, 1} {
		result := reverse(v)
		fmt.Println(result)
		result = reverse2(v)
		fmt.Println(result)
	}
}

// 难度：简单  数学
// https://leetcode-cn.com/problems/reverse-integer/

func reverse(x int) int {
	var result int
	var mod, times = 1, 0
	var numbers []int
	number := x % mod
	for number != x {
		number = x % mod
		numbers = append(numbers, (number*10)/mod)
		mod *= 10
		times++
	}
	for i, v := range numbers {
		result += v * int(math.Pow10(times-i-1))
	}
	if int(int32(result)) != result {
		return 0
	}
	return result
}

// 优化
func reverse2(x int) int {
	var min, max = math.MinInt32, math.MaxInt32
	var result int
	var mod, number = 1, 0
	for number != x {
		number = x % mod
		result *= 10
		result += (number * 10) / mod
		if min > result {
			return 0
		}
		if max < result {
			return 0
		}
		mod *= 10
	}
	return result
}
