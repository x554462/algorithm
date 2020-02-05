package main

import "fmt"

func main() {
	result := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)
}

// 难度：中等  哈希表/双指针/字符串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 || l == 1 {
		return l
	}
	bitmap := NewBitMap()
	var a, b uint8
	var i, j, max = 0, 1, 1
	a = s[i]
	bitmap.Add(a)
	for l > j {
		b = s[j]
		if bitmap.Has(b) {
			if max < j-i {
				max = j - i
			}
			for bitmap.Has(b) {
				bitmap.Del(s[i])
				i++
			}
		} else if l == j+1 {
			if j+1-i > max {
				max = j + 1 - i
			}
		}
		bitmap.Add(b)
		j++
	}
	return max
}

type Bitmap struct {
	words []uint64
}

func NewBitMap() *Bitmap {
	return &Bitmap{}
}
func (bitmap *Bitmap) Has(i uint8) bool {
	num := int(i)
	word, bit := num/64, uint(num%64)
	return (word < len(bitmap.words)) && ((bitmap.words[word] & (1 << bit)) != 0)
}

func (bitmap *Bitmap) Add(i uint8) {
	num := int(i)
	word, bit := num/64, uint(num%64)
	for word >= len(bitmap.words) {
		bitmap.words = append(bitmap.words, 0)
	}
	// 不存在num则将num加入bitmap
	if bitmap.words[word]&(1<<bit) == 0 {
		bitmap.words[word] |= 1 << bit
	}
}

func (bitmap *Bitmap) Del(i uint8) {
	num := int(i)
	word, bit := num/64, uint(num%64)

	if bitmap.words[word]&(1<<bit) != 0 {
		bitmap.words[word] &= bitmap.words[word] ^ (1 << bit)
	}
}
