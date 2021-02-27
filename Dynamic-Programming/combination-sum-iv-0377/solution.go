package main

import "fmt"

// 给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
// f(k) = sum(f(k-i)) | i in nums
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	// base case
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if i >= n {
				dp[i] += dp[i-n]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 4
	res := combinationSum4(nums, target)
	fmt.Println(res)
}
