package sort

// 分治法（Divide and Conquer）
func mergeSort(nums []int) []int {
	size := len(nums)
	if size < 2 {
		return nums
	}
	mid := size / 2
	lNums := nums[0:mid]
	rNums := nums[mid:]
	lNums = mergeSort(lNums)   // 排左边
	rNums = mergeSort(rNums)   // 排右边
	return merge(lNums, rNums) // 合并
}

// left & right 两个已经排好序的序列进行合并
func merge(left, right []int) []int {
	m, n := len(left), len(right)
	res := []int{}
	i, j := 0, 0
	for i < m && j < n {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	if i < m {
		res = append(res, left[i:]...)
	}
	if j < n {
		res = append(res, right[j:]...)
	}
	return res
}
