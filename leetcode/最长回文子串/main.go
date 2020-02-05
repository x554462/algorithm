package main

import "fmt"

func main() {
	for _, s := range []string{"cbbd", "cbbbbd", "aas", "ass", "abba", "avavav", "aaas", "aaaas", "a", ""} {
		result := longestPalindrome(s)
		fmt.Println(result)
	}
}

// 难度：中等  字符串/动态规划
// https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return s
	}
	var maxRange, left, right = 0, 0, 0

	var i, j = 1, 1
	for i < length {
		var min, max = 0, 0
		for i > 0 && s[i] == s[i-1] {
			if i-1 < 0 {
				break
			}
			i = i - 1
		}
		for j < length-1 && s[j] == s[j+1] {
			if j+1 > length {
				break
			}
			j = j + 1
		}
		min, max = i, j
		j++
		i = j
		for min-1 >= 0 && max+1 < length {
			if s[min-1] != s[max+1] {
				break
			}
			min, max = min-1, max+1
		}
		if maxRange < max-min {
			maxRange = max - min
			right = max
			left = min
		}
	}

	var result []uint8
	for ; left <= right; left++ {
		result = append(result, s[left])
	}
	return string(result)
}
