package main

import (
	"fmt"
	"sort"
)

// 解题思路
// 数组遍历 时间复杂度 O(n^2)
// 关键考虑去重
func threeSum(nums []int) [][]int {
	res := [][]int{}
	size := len(nums)
	if size <= 2 {
		return res
	}
	sort.Ints(nums) // 排序
	var left, right int

	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			// 如果 nums[i]大于 0，则三数之和必然无法等于 0，结束循环
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			// 去重
			continue
		}
		left, right = i+1, size-1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 可能不止一个满足条件，接着找
				for left < right && nums[left] == nums[left+1] {
					// 连续相等的排除
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
