package main

import (
	"fmt"
)

// 解题思路
// 位运算 计算两数相加
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		tmp := a ^ b     // 异或的作用: 相加各个二进制位的值
		b = (a & b) << 1 // &的作用: 计算进位
		a = tmp
	}

	return a
}

func main() {
	fmt.Println(getSum(10, 2))
	// fmt.Println(getSum(-2, 3))
}
