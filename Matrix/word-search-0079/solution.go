package main

import "fmt"

// 给定一个二维网格和一个单词，找出该单词是否存在于网格中
// 解题思路: 回溯算法 backtracking
var (
	row, col int
	visited  [][]bool
)

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}
	row, col = len(board), len(board[0])
	visited = make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(i, j, 0, board, word) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j, idx int, board [][]byte, word string) bool {
	// 数组越界 或者 字符不匹配 或者 已经访问过 返回 false
	if i >= row || j >= col || i < 0 || j < 0 || board[i][j] != word[idx] || visited[i][j] {
		return false
	}

	if idx == len(word)-1 {
		return board[i][j] == word[idx]
	}

	// 上下左右 四个方向
	visited[i][j] = true
	if dfs(i, j-1, idx+1, board, word) || dfs(i, j+1, idx+1, board, word) || dfs(i-1, j, idx+1, board, word) || dfs(i+1, j, idx+1, board, word) {
		return true
	}
	visited[i][j] = false // 状态重置

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
