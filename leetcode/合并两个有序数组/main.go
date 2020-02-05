package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := make([]int, 6, 6)
	nums2 := make([]int, 3, 3)
	nums1[0] = 1
	nums1[1] = 2
	nums1[2] = 3
	nums2[0] = 2
	nums2[1] = 5
	nums2[2] = 6
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

// 从尾部开始遍历
func merge2(nums1 []int, m int, nums2 []int, n int) {
	k := n + m - 1
	i := m - 1
	j := n - 1

	if i < 0 && j < 0 {
		return
	}

	for j >= 0 && i >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

// 错误解法
func merge3(nums1 []int, m int, nums2 []int, n int) {
	var nums []int
	if m == 0 {
		nums1 = nums2
	} else if n != 0 {
		var i, j = 0, 0
		for {
			if i < m && j < n {
				if nums1[i] <= nums2[j] {
					nums = append(nums, nums1[i])
					i++
				} else {
					nums = append(nums, nums2[j])
					j++
				}
			} else if i < m {
				nums = append(nums, nums1[i])
				i++
			} else if j < n {
				nums = append(nums, nums2[j])
				j++
			} else {
				break
			}
		}
	}
	nums1 = nums
}
