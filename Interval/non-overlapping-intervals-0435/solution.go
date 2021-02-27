package main

import (
	"fmt"
	"sort"
)

// 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	res := 0
	for i := 1; i < len(intervals); i++ {
		// 有重叠区间
		if prev[1] > intervals[i][0] {
			res++
			// 比较 end 大小
			if prev[1] > intervals[i][1] {
				prev = intervals[i] // 选 end 小的 因为要找到需要移除区间的最小数量
			}
			continue
		}
		prev = intervals[i]
	}

	return res
}

func main() {
	intervals := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}
	fmt.Println(eraseOverlapIntervals(intervals))
}
