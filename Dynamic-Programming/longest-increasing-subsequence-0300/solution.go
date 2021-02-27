package main

import "fmt"

// 解题思路 动态规划 ===> 数学归纳法
// dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度
// dp[i] = max(dp[i], dp[j] + 1 | j in (0, i)) 时间复杂度 O(n^2)
func lengthOfLIS(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	for k := range dp {
		dp[k] = 1 // init
	}
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			// (0, i) 的区间里找 nums[i] 小的数
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for _, val := range dp {
		if val > res {
			res = val
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
