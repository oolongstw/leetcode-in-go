package main

import "fmt"

// 假如计划在第 i 天卖出股票，那么最大利润的差值一定是在[0, i-1] 之间选最低点买入
// 所以遍历数组，依次求每个卖出时机的的最大差值，再从中取最大值。
// 滑动窗口法解决
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	size := len(prices)
	min, max := 0, 1
	profit := prices[max] - prices[min]
	for max < size {
		if prices[min] > prices[max] {
			min = max // 遇到价格更低时直接重置买入价格
			max++
		} else {
			profit = bigger(profit, prices[max]-prices[min])
			max++
		}
	}
	if profit < 0 {
		return 0
	}
	return profit
}

func bigger(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	prices := []int{7, 3, 6, 1, 2}
	fmt.Println(maxProfit(prices))
}
