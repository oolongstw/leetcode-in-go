package main

import "fmt"

// 最长公共子串 二维动态规划
func LCS(s1, s2 string) (res int) {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	res = 0
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // dp formula
				res = max(res, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}
	fmt.Println(dp)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s1 := "abcdefg"
	s2 := "xyzabcd"
	fmt.Println(LCS(s1, s2))
}
