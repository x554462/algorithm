package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := restoreIpAddresses("25525511135")
	fmt.Println(result)
	result = restoreIpAddresses("010010")
	fmt.Println(result)
}

// 难度：中等  字符串/回溯算法
// https://leetcode-cn.com/problems/restore-ip-addresses/
func restoreIpAddresses(s string) []string {
	var result []string
	l := len(s)
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				for d := 1; d < 4; d++ {
					if a+b+c+d == l {
						i1, _ := strconv.ParseInt(s[0:a], 10, 64)
						i2, _ := strconv.ParseInt(s[a:a+b], 10, 64)
						i3, _ := strconv.ParseInt(s[a+b:a+b+c], 10, 64)
						i4, _ := strconv.ParseInt(s[l-d:], 10, 64)
						if i1 <= 255 && i2 <= 255 && i3 <= 255 && i4 <= 255 {
							ip := fmt.Sprintf("%d.%d.%d.%d", i1, i2, i3, i4)
							if len(ip) == l+3 {
								result = append(result, ip)
							}
						}
					}
				}
			}
		}
	}
	return result
}
