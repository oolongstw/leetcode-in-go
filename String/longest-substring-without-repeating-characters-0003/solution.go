package main

import (
	"fmt"
)

// 解题思路: 滑动窗口
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	record := make(map[byte]int) // 记录 字符对应的位置
	maxLength := 1

	var (
		left, right int
	)
	left = 0
	for right = 0; right < n; right++ {
		pos, ok := record[s[right]]
		if ok && left <= pos {
			// 有重复的字符 并且在 left 指针右边，left 指针向右移
			left = pos + 1
		} else {
			maxLength = max(maxLength, right-left+1)
		}
		record[s[right]] = right
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sample := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(sample))
}
