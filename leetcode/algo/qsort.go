package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 1, 6, 6, 9}
	QSort(&nums)
	fmt.Println(nums)
}

func QSort(nums *[]int) {
	qSort(nums, 0, len(*nums)-1)
}

func qSort(nums *[]int, left, right int) {
	if left < right {
		l, r := left, right
		v := (*nums)[l]
		for l < r {
			for l < r && (*nums)[r] > v {
				r--
			}
			if l < r {
				(*nums)[l] = (*nums)[r]
				l++
			}
			for l < r && (*nums)[l] < v {
				l++
			}
			if l < r {
				(*nums)[r] = (*nums)[l]
				r--
			}
		}
		(*nums)[l] = v
		qSort(nums, left, l-1)
		qSort(nums, l+1, right)
	}
}
