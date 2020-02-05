package main

import "fmt"

func main() {
	for _, v := range [][]string{{"xx", "yy"}, {"xy", "yx"}, {"xxyyxyxyxx", "xyyxyxxxyx"}, {"xxx", "xyy"}} {
		result := minimumSwap(v[0], v[1])
		fmt.Println(result)
	}
}

// 难度：中等  贪心算法/字符串
// https://leetcode-cn.com/problems/minimum-swaps-to-make-strings-equal/submissions/

func minimumSwap(s1 string, s2 string) int {
	var xy, yx = 0, 0
	for i, v := range s1 {
		if v != int32(s2[i]) {
			if v == 'x' {
				xy++
			} else {
				yx++
			}
		}
	}
	// （xy+yx)&1 == 1 // 判断奇偶
	if (xy+yx)%2 == 1 {
		return -1
	}
	return (xy+1)/2 + (yx+1)/2
}
