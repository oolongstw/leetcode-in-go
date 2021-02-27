package main

import "fmt"

// 解题思路
// 动态规划
// 爬第n阶楼梯的方法数量，等于 2 部分之和
// 爬上 n-1 阶楼梯的方法数量。因为再爬1阶就能到第n阶
// 爬上 n-2 阶楼梯的方法数量。因为再爬2阶就能到第n阶
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	fmt.Println(climbStairs(10))
}
