package sort

// 快速排序 又称 partition-exchange sort 划分交换排序，快速排序使用分治法（Divide and conquer）策略来把一个串（list）分为两个子串（sub-lists）
// 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面。在这个分区退出之后，该基准就处于最后正确的位置上
// in place change 不用占据额外的空间
func quickSort(nums []int) {
	_qsort(nums, 0, len(nums))
}

func _qsort(nums []int, l, r int) {
	if l < r {
		p := _partition(nums, l, r) // 根据 position 切分为两部分 继续递归
		_qsort(nums, l, p-1)
		_qsort(nums, p+1, r)
	}
}

func _partition(arr []int, l, r int) int {
	pos := l
	pivot := arr[pos] // 取最左边的为 pivot
	for i := l + 1; i < r; i++ {
		if arr[i] < pivot {
			pos++
			arr[pos], arr[i] = arr[i], arr[pos]
		} else {
			continue
		}
	}
	arr[pos], arr[l] = arr[l], arr[pos]
	return pos
}
