package main

import (
	"fmt"
)

// 解题思路
// 双指针 缩减搜索空间策略
func maxArea(height []int) int {
	n := len(height)      // n >= 2
	left, right := 0, n-1 // 先从宽度最长开始
	area := 0
	for left < right {
		tmpHeight := min(height[left], height[right])
		area = max(area, tmpHeight*(right-left))
		if height[left] < height[right] {
			// 如果左边高度小于右边高度，则移动右边=>宽度变小=>不可能获得最大值
			left++
		} else {
			right--
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
