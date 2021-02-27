package main

import "fmt"

// 解题思路 状态转移方程 f[i][j] = f[i-1][j] + f[i][j-1]
func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		f[i][0] = 1
	}
	for i := 0; i < n; i++ {
		f[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}

	return f[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(7, 3))
}
