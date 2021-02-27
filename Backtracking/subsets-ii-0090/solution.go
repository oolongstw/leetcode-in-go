package main

import (
	"fmt"
	"sort"
)

// 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）
// 解题思路: 和 78 题相同，需要处理重复元素的情况
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		res = append(res, []int{})
		return res
	}
	sort.Ints(nums)
	path := []int{}
	used := make([]bool, len(nums))
	for i := 0; i <= len(nums); i++ {
		// 由组合问题求子集问题
		combine(nums, 0, i, path, &res, used) // subset = c(n,0) + ... + c(n,k) + ... + c(n,n)
	}
	return res
}

func combine(nums []int, start, k int, path []int, res *[][]int, used []bool) {
	if len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		// 去重逻辑: 本次循环不取重复数字，下次可能会取
		// if i >= 1 && !used[i-1] && nums[i] == nums[i-1] {
		// 	continue
		// }
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		combine(nums, i+1, k, path, res, used)
		// revert
		path = path[:len(path)-1]
		used[i] = false
	}
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
