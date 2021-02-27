package main

import (
	"fmt"
	"sort"
)

// 给定一个数组 candidates 和一个目标数 target，找出 candidates 中所有可以使数字和为 target 的组合
// candidates 中的每个数字在每个组合中只能使用一次
// 解题思路: 排序 + 剪枝，和 39 题相同，另外处理重复数字的问题
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	fmt.Println(candidates)
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
		if i > start && candidates[i] == candidates[i-1] {
			continue // 去重的关键逻辑: 下次循环 过滤掉重复的数字，本次循环不过滤
		}
		current = append(current, candidates[i])
		dfs(candidates, i+1, target-candidates[i], current, res) // 下层递归的 target
		current = current[:len(current)-1]                       // revert
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
