package main

import "fmt"

// 解题思路: 动态规划
// 用 f(i) 代表以第 i 个数结尾的「连续子数组的最大和」
// 于是可以写出这样的动态规划转移方程 f(i) = max{f(i-1) + nums[i], nums[i]}
func maxSubArray(nums []int) int {
	size := len(nums)
	res := nums[0]
	if size <= 1 {
		return res
	}
	// dp := make([]int, size)
	// dp[0] = nums[0]
	prevMax := nums[0]
	nowMax := 0
	for i := 1; i < size; i++ {
		// dp[i] = bigger(dp[i-1]+nums[i], nums[i]) // 继续优化: dp[i] 只和 dp[i-1] 相关可以只用一个变量 prevMax来维护 把空间复杂度降到 O(1)
		// res = bigger(dp[i], res)
		if prevMax < 0 {
			nowMax = nums[i]
		} else {
			nowMax = prevMax + nums[i]
		}
		// nowMax = bigger(prevMax+nums[i], nums[i])
		prevMax = nowMax
		res = bigger(nowMax, res)
	}
	return res
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
