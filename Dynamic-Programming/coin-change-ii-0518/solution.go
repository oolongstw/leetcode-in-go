package main

import "fmt"

// 0-1 背包问题 https://leetcode-cn.com/problems/coin-change-2/solution/dong-tai-gui-hua-wan-quan-bei-bao-wen-ti-by-liweiw/
// 若只使用 coins 中的前 i 个硬币的面值，若想凑出金额 j，有 dp[i][j] 种凑法 并且硬币数量不限
// dp[i][j] = dp[i][j-coins[i-1]] + dp[i-1][j] 选择 第i个硬币 或者 不选第i个硬币 两种情况
func change(amount int, coins []int) int {
	size := len(coins)
	dp := make([][]int, size+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	// base case: dp[i][0] = 1 dp[0][j] = 0
	for i := 0; i <= size; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= size; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j] = dp[i][j-coins[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[size][amount]
}

func main() {
	coins := []int{1, 2, 5}
	fmt.Println(change(5, coins))
}
