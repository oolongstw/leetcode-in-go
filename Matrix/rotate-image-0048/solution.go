package main

import "fmt"

// 旋转图像 图像顺时针旋转90度 n*n 的矩阵
// 解题思路: 找到旋转的元素的映射关系
func rotate(matrix [][]int) {
	n := len(matrix)
	// 旋转的映射关系 (i, j) -> (j, n - 1 - i)
	// 按照圈遍历，由外到内
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ { // 先从外圈横向遍历
			tmp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
