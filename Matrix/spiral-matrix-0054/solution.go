package main

import "fmt"

// 给定一个包含 m x n 个元素的矩阵（m 行, n 列），按照顺时针螺旋顺序，返回矩阵中的所有元素。
// 解题思路: 设置left, right, top, bottom 四个边界，由外层想内层遍历
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	res := []int{}
	target, num := m*n, 1 // target 总共要打印的元素个数
	left, right := 0, n-1
	top, bottom := 0, m-1
	// 外圈到内圈遍历
	for num <= target {
		// left to right
		for i := left; i <= right && num <= target; i++ {
			if matrix[top][i] == 6 {
				fmt.Println(100)
			}
			res = append(res, matrix[top][i])
			num++
		}
		top++
		// top to bottom
		for i := top; i <= bottom && num <= target; i++ {
			res = append(res, matrix[i][right])
			num++
		}
		right--
		// right to left
		for i := right; i >= left && num <= target; i-- {
			res = append(res, matrix[bottom][i])
			num++
		}
		bottom--
		// bottom to top
		for i := bottom; i >= top && num <= target; i-- {
			res = append(res, matrix[i][left])
			num++
		}
		left++
	}

	return res
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}
