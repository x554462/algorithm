package main

import "fmt"

func main() {
	for _, v := range [][][]int{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {1, 2, 3}}, {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {9, 10, 11, 12}}} {
		result := spiralOrder(v)
		fmt.Println(result)
	}
}

// 难度：中等  数组
// https://leetcode-cn.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}
	var result []int
	sum := m * n
	i, j := 0, 0
	a, b, c, d := 0, n-1, m-1, 0
	ai, aj, ri, rj := false, true, false, false
	for {
		//fmt.Printf("%d %d %d %d\n", a, b, c, d)
		//fmt.Printf("%d %d\n", i, j)
		result = append(result, matrix[i][j])
		sum--
		if sum <= 0 {
			break
		}
		if aj {
			//fmt.Printf("111111111111 %d %d\n", i, j)
			if j == b {
				i++
				aj = false
				ai = true
			} else if j++; j == b {
				b--
				aj = false
				ai = true
			}
		} else if ai {
			//fmt.Printf("222222222222 %d %d\n", i, j)
			if i == c {
				j++
				ai = false
				rj = true
			} else if i++; i == c {
				c--
				ai = false
				rj = true
			}
		} else if rj {
			//fmt.Printf("333333333333 %d %d\n", i, j)
			if j == d {
				i--
				rj = false
				ri = true
			} else if j--; j == d {
				d++
				rj = false
				ri = true
			}
		} else if ri {
			//fmt.Printf("444444444444 %d %d\n", i, j)
			if i == a {
				j--
				ri = false
				aj = true
			} else if i--; i <= a {
				a++
				ri = false
				aj = true
				i++
				j++
			}
		}
	}
	return result
}
