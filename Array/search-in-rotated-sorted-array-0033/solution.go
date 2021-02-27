package main

import "fmt"

// 解题思路
// 二分查找 判断 target 是否在区间内
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			// [l, mid] 递增
			if nums[l] <= target && nums[mid] >= target {
				// target 在该区间内
				r = mid
			} else {
				// 否则从后半部分开始找
				l = mid + 1
			}
		} else {
			if nums[mid] <= target && nums[r] >= target {
				// target 在该区间内
				l = mid
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 4))
}
