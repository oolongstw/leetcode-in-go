package main

import (
	"fmt"
)

// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）
func subsets(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		res = append(res, []int{})
		return res
	}
	path := []int{}
	for i := 0; i <= len(nums); i++ {
		// 由组合问题求子集问题
		combine(nums, 0, i, path, &res) // subset = c(n,0) + ... + c(n,k) + ... + c(n,n)
	}
	return res
}

// 候选集合是 (nums[start]...nums[n])
func combine(nums []int, start, k int, path []int, res *[][]int) {
	if len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		combine(nums, i+1, k, path, res)
		// revert
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
