package main

import "fmt"

func rob(nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	if N == 1 {
		return nums[0]
	}
	if N == 2 {
		return max(nums[0], nums[1])
	}
	// house[1] 和 house[N] 只能抢一个
	// 在 house[1]~house[N-1] 和 house[2]~house[N] 两者中找最大值
	dp := make([]int, N)
	// house[1]~house[N-1] 和 198 题相同
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= N-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	tmp := dp[N-1]
	// house[2]~house[N]
	dp[1] = nums[1]
	for i := 2; i <= N-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return max(tmp, dp[N-1])
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
