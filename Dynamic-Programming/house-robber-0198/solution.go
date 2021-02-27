package main

import "fmt"

// 动态规划
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	// 子问题
	// f(k) = 偷 [0..k) 房间中的最大金额
	// f(0) = 0
	// f(1) = nums[0]
	// f(k) = max{ rob(k-1), nums[k-1] + rob(k-2) }
	for k := 2; k <= n; k++ {
		dp[k] = max(dp[k-1], dp[k-2]+nums[k-1])
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
