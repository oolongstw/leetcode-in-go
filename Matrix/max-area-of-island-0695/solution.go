package main

import "fmt"

var (
	row, col int
)

// 岛屿的最大面积
// 解题思路: dfs
func maxAreaOfIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}
	row, col = len(grid), len(grid[0])
	area, max := 0, 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			dfs(grid, i, j, &area)
			if max < area {
				max = area // 维护一个最大值
			}
			area = 0 // 下次调用重置 area 的值
		}
	}
	return max
}

func dfs(grid [][]int, i, j int, area *int) {
	if i >= row || j >= col || i < 0 || j < 0 || grid[i][j] != 1 {
		return
	}
	*area++ // grid[i][j] = 1 面积加一
	// 上下左右 继续找
	grid[i][j] = 2 // 防止重复扫描
	dfs(grid, i, j-1, area)
	dfs(grid, i, j+1, area)
	dfs(grid, i-1, j, area)
	dfs(grid, i+1, j, area)
}

func main() {
	grid := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
