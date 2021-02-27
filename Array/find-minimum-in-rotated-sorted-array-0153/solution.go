package main

import "fmt"

// 按照升序排序的数组在预先未知的某个点(这个点有可能不是数组的index?!)上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2]
// 解题思路
// 二分查找
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right { // 实际上是不会跳出循环，当 left==right 时直接返回
		if nums[left] <= nums[right] { // 如果 nums[left] <= nums[right] [left,right] 递增，直接返回
			return nums[left]
		}
		mid := left + (right-left)/2
		if nums[left] <= nums[mid] { // [left,mid] 连续递增，则在 [mid+1,right] 查找
			left = mid + 1
		} else {
			right = mid // [left,mid] 不连续，在 [left,mid] 查找
		}
	}
	return nums[left]
}

func main() {
	nums := []int{3, 1, 2}
	fmt.Println(findMin(nums))
}
