package main

import "fmt"

// 解题思路: 分治法
// 数组中某数除自身以外各元素的乘积 = 该数左侧所有数的乘积 * 该数右侧所有数的乘积
// 时间复杂度: O(n) 空间复杂度: O(1)
func productExceptSelf(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	tmp := 1
	// 求每个数左侧所有数的乘积
	for i := 0; i < size; i++ {
		res[i] = tmp
		tmp *= nums[i]
	}
	tmp = 1
	// 求每个数右侧所有数的乘积，并与其左侧所有数乘积相乘
	for i := size - 1; i >= 0; i-- {
		res[i] *= tmp
		tmp *= nums[i]
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
