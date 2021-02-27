package main

import (
	"fmt"
	"sort"
)

// 给定一个可包含重复数字的序列，返回所有不重复的全排列
// 解题思路: 关键在于如何去重 排序 + 去重
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	sort.Ints(nums) // 排序
	trace := []int{}
	process(nums, trace, used, &res)
	return res
}

func process(nums, trace []int, used []bool, res *[][]int) {
	if len(trace) == len(nums) {
		tmp := make([]int, len(trace))
		copy(tmp, trace)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 判断重复元素
		if used[i] {
			continue
		}
		if i >= 1 && nums[i] == nums[i-1] && !used[i-1] {
			// i-1 没用过 但是 nums[i] == nums[i-1] 则表示: 同一层需要剪枝
			continue
		}
		// make decision
		trace = append(trace, nums[i])
		used[i] = true
		process(nums, trace, used, res)
		// revert decision
		trace = trace[:len(trace)-1]
		used[i] = false
	}
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(permuteUnique(nums))
}
