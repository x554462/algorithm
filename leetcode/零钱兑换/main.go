package main

import (
	"fmt"
)

func main() {
	result := coinChange([]int{1, 2, 5}, 11)
	fmt.Println(result)
	result = coinChange2([]int{1, 2, 5}, 11)
	fmt.Println(result)
	result = coinChange3([]int{1, 2, 5}, 11)
	fmt.Println(result)
}

// 难度：中等  动态规划
// https://leetcode-cn.com/problems/coin-change/
// 详细题解 https://leetcode-cn.com/problems/coin-change/solution/dong-tai-gui-hua-tao-lu-xiang-jie-by-wei-lai-bu-ke/

// 递归暴力
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var result = -1
	for _, coin := range coins {
		if amount-coin < 0 {
			continue
		}
		subPro := coinChange(coins, amount-coin)
		if subPro == -1 {
			continue
		}
		if result == -1 || result > subPro+1 {
			result = subPro + 1
		}
	}
	return result
}

// 带备忘录的递归
var coinAmount = map[int]int{}

func coinChange2(coins []int, amount int) int {
	coinAmount = map[int]int{}
	return recursion(coins, amount)
}
func recursion(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if v, ok := coinAmount[amount]; ok {
		return v
	}
	var result = -1
	for _, coin := range coins {
		if amount-coin < 0 {
			continue
		}
		subPro := recursion(coins, amount-coin)
		if subPro == -1 {
			continue
		}
		if result == -1 || result > subPro+1 {
			result = subPro + 1
		}
	}
	coinAmount[amount] = result
	return result
}

// 动态规划
func coinChange3(coins []int, amount int) int {
	var coinAmount = map[int]int{}
	coinAmount[0] = 0
	for _, coin := range coins {
		coinAmount[coin] = 1
	}
	for i := 0; i <= amount; i++ {
		res := -1
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if v, ok := coinAmount[i-coin]; ok && v != -1 {
				if res == -1 || res > v+1 {
					res = v + 1
				}
			}
		}
		if _, ok := coinAmount[i]; !ok {
			coinAmount[i] = res
		}
	}
	return coinAmount[amount]
}
