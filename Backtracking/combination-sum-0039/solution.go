package main

import (
	"fmt"
	"sort"
)

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合
// 解题思路: 排序 + 剪枝
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	dfs(candidates, 0, target, []int{}, &res)
	return res
}

func dfs(candidates []int, start, target int, current []int, res *[][]int) {
	if 0 == target {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			break // 剪枝
		}
		current = append(current, candidates[i])
		dfs(candidates, i, target-candidates[i], current, res) // 下层递归的 target
		current = current[:len(current)-1]                     // revert
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
