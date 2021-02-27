package main

import "fmt"

// 矩阵置零
func setZeroes(matrix [][]int) {
	row, col := len(matrix), len(matrix[0])
	// 找到为0的值，记录行信息和列信息
	rowRecord := map[int]bool{}
	colRecord := map[int]bool{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				rowRecord[i] = true
				colRecord[j] = true
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rowRecord[i] || colRecord[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}
