package sort

// 堆排序: 实质上是一个二叉树，分为大顶堆和小顶堆
// 堆排序的基本思想是：将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点。将其与末尾元素进行交换，此时末尾就为最大值。
// 然后将剩余n-1个元素重新构造成一个堆，这样会得到n个元素的次小值。如此反复执行，便能得到一个有序序列了
// 1. 创建一个堆 H[0... n-1].
// 2. 把堆顶（最大值）和堆尾互换.
// 3. 把堆的尺寸缩小 1，重新建堆
// 重复步骤2，直到堆的尺寸为 1
// 图解: https://www.cnblogs.com/chengxiao/p/6129630.html
func heapSort(nums []int) {
	n := len(nums)
	if n < 1 {
		return
	}
	// heapify 自底向上 从半长处开始
	for i := n / 2; i >= 0; i-- {
		heap(nums, i)
	}
	for j := n - 1; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		// 重新建堆
		heap(nums[:j], 0)
	}
}

// heap 建堆操作
func heap(nums []int, i int) {
	n := len(nums)
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < n && nums[largest] < nums[left] {
		largest = left
	}
	if right < n && nums[largest] < nums[right] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		// reheap at largest position
		heap(nums, largest)
	}
}
