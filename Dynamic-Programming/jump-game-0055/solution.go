package main

import (
	"fmt"
)

// 解题思路 数组中只要没有0就一定能跳到最后
// 遍历数组，更新能跳到的最远距离
func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	longest := 0
	for i := 0; i < n; i++ {
		if i > longest {
			return false
		}
		longest = max(nums[i]+i, longest)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 0, 0}
	fmt.Println(canJump(nums))
}
