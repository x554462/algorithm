package main

import (
	"fmt"
	"math"
)

func main() {
	for _, v := range []string{"42", "   -42", "41-93 with words", "words and 987", "-91283472332", "+1", "+-2", "   +0 123"} {
		result := myAtoi(v)
		fmt.Println(result)
	}
}

// 难度：中等  数学/字符串
// https://leetcode-cn.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
	var min, max = math.MinInt32, math.MaxInt32
	var result, times, flag = 0, 1, false
	for _, v := range str {
		if v <= '9' && v >= '0' {
			flag = true
			result = (result * 10) + int(v-'0')
			if times == 1 && result > max {
				result = max
				break
			} else if times == -1 && -result < min {
				result = -min
				break
			}
		} else if v == '-' {
			if flag {
				break
			}
			flag = true
			times = -1
		} else if v == '+' {
			if flag {
				break
			}
			flag = true
		} else if v != ' ' {
			if flag {
				break
			}
			break
		} else if flag {
			break
		}
	}
	return result * times
}
