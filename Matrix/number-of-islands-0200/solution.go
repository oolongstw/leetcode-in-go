package main

import "fmt"

var (
	row, col int
)

// 岛屿数量
// 解题思路: dfs 先找到第一个不为 0 的格子 开始 dfs
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	res := 0
	row, col = len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			} else {
				grid[i][j] = '2'
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int) {
	if i >= row || j >= col || i < 0 || j < 0 || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2' // 防止重复搜索
	// 上下左右 四个方向搜索
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
