package main

import (
	"fmt"
	"sort"
)

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// intervals 按照 start 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := 0; i < len(intervals); i++ {
		// 没有交集的情况
		if prev[1] < intervals[i][0] {
			res = append(res, prev)
			prev = intervals[i]
			continue
		}
		if intervals[i][1] > prev[1] {
			prev[1] = intervals[i][1]
			continue
		}
	}
	res = append(res, prev)

	return res
}

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{12, 13},
	}
	fmt.Println(merge(intervals))
}
