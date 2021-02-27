package main

import "fmt"

// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates(nums []int) int {
	// 思路是快慢指针+快排思想
	// 从左到右，左为后，右为前，初始时，前为索引1，后为索引0
	// 以后者为基准，判断前者是否比后者大，如果是，前者和后者加一之后的索引进将fast复制到low的位置上，把不重复的值插到对应的位置上
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}

	low, fast := 0, 1
	for fast < length {
		if nums[fast] != nums[low] {
			low++
			nums[low] = nums[fast]
			fast++
		} else {
			fast++
		}
	}
	return low + 1
}

func main() {
	nums := []int{1, 1, 1, 2, 3, 4, 5, 5}
	fmt.Println(removeDuplicates(nums))
}
