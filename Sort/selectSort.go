package sort

// 工作原理是每次找出第 i 小的元素, 然后将这个元素与数组第 i 个位置上的元素交换。
func selectSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		pos := i
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				pos = j
			}
		}
		nums[i], nums[pos] = nums[pos], nums[i]
	}
}
