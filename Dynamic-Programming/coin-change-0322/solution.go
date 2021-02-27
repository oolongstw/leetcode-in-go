package main

import (
	"fmt"
)

// 计算可以凑成总金额所需的最少的硬币个数 https://leetcode-cn.com/problems/coin-change/
// 解题思路 https://leetcode-cn.com/problems/coin-change/solution/dong-tai-gui-hua-shi-yong-wan-quan-bei-bao-wen-ti-/
// f(n) = min{f(n-coin)+1 | coin in coins}
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := range f {
		f[i] = amount + 1 // 初始化为 amount + 1 方便最终取最小值
	}
	f[0] = 0
	for i := 0; i < amount+1; i++ {
		for _, c := range coins {
			// remain coin
			if i-c < 0 {
				continue
			}
			f[i] = min(f[i], f[i-c]+1)
		}
	}
	// 最多 amount 个硬币(全是1的话)
	if f[amount] > amount {
		return -1
	}
	return f[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{5}
	amount := 6
	fmt.Println(coinChange(coins, amount))
}
