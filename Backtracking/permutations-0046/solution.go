package main

import "fmt"

// 给定一个 「没有重复」 数字的序列，返回其所有可能的全排列。
// 解题思路: 回溯算法 -> 递归之前做出选择，在递归之后撤销刚才的选择
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	used := make([]bool, len(nums))
	trace := []int{}
	backtrack(trace, nums, &res, used)
	return res
}

func backtrack(trace, nums []int, res *[][]int, used []bool) {
	// 结束条件
	if len(trace) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, trace) // 需要 copy slice to a new slice, trace 共享同一个底层数组 如果不 copy 会导致最终 res 中出现相同的结果
		*res = append(*res, tmp)
		return
	}
	// decision tree
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		// 在递归之前做选择，在递归之后撤销刚才的选择
		trace = append(trace, nums[i]) // make decision
		used[i] = true
		backtrack(trace, nums, res, used)
		trace = trace[:len(trace)-1] // revert decision
		used[i] = false
	}
}

// 方法2: swap 原地交换顺序 优化空间复杂度
func permute2(nums []int) [][]int {
	result := [][]int{}
	if len(nums) <= 0 {
		return result
	}
	start := 0
	process(nums, start, &result)
	return result
}

func process(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		t := make([]int, len(nums))
		copy(t, nums)
		*result = append(*result, t)
		// *result = append(*result, nums)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		process(nums, start+1, result)
		nums[i], nums[start] = nums[start], nums[i]
	}
	return
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
	fmt.Println(permute2(nums))
}
