package main

import "fmt"

// 解题思路: 滑动窗口 窗口滑动的时机 窗口大小改变的时机
func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	n := len(s)
	record := make(map[byte]int)
	currentMax := 0 // 记录滑动窗口内相同字母出现次数的历史最大值
	for right = 0; right < n; right++ {
		record[s[right]]++
		currentMax = max(currentMax, record[s[right]])
		if right-left+1 > currentMax+k { // 超过长度的最大值
			// 窗口滑动 left 指针向右移
			record[s[left]]--
			left++
		}
	}
	return right - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s, k := "ABABB", 2
	fmt.Println(characterReplacement(s, k))
}
