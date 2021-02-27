package main

import "fmt"

// 解题思路
// dpMax[i] 表示以第 i 个元素的结尾的子数组，乘积最大的值，也就是这个数组必须包含第 i 个元素。
// 由于有负数，需要 dpMin 来记录以第 i 个元素的结尾的子数组，乘积最小的值
func maxProduct(nums []int) int {
	res := nums[0]
	size := len(nums)
	// dpMax := make([]int, size) // 存最大值
	// dpMin := make([]int, size) // 存最小值
	// dpMax[0] = nums[0]
	// dpMin[0] = nums[0]
	dpMax, dpMin := nums[0], nums[0]
	var preMax int

	for i := 1; i < size; i++ {
		// dpMax[i] = max(max(dpMax[i-1]*nums[i], nums[i]), dpMin[i-1]*nums[i]) // dpMax[i] 的取值无非就是三种，dpMax[i-1] * nums[i]、dpMin[i-1] * nums[i] 以及 nums[i]。
		// dpMin[i] = min(min(dpMax[i-1]*nums[i], nums[i]), dpMin[i-1]*nums[i]) // 同理 dpMin[i] 的取值也只有三种
		// res = max(dpMax[i], res)
		preMax = dpMax
		dpMax = max(max(preMax*nums[i], nums[i]), dpMin*nums[i]) // 继续优化 dp[i] 只和 dp[i-1] 相关，减少空间复杂度
		dpMin = min(min(preMax*nums[i], nums[i]), dpMin*nums[i])
		res = max(dpMax, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{2, -5, -2, -4, 3} // res = 24
	fmt.Println(maxProduct(nums))
}
