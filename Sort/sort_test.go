package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{5, 7, 8, 4, 12, 1, 6}
	quickSort(nums)
	fmt.Println(nums)
}

func TestSelectSort(t *testing.T) {
	nums := []int{5, 7, 8, 4, 12, 1}
	selectSort(nums)
	fmt.Println(nums)
}

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 7, 8, 4, 12, 1}
	bubbleSort(nums)
	fmt.Println(nums)
}

func TestInsertSort(t *testing.T) {
	nums := []int{5, 7, 8, 4, 12, 1}
	insertSort(nums)
	fmt.Println(nums)
}

func TestHeapSort(t *testing.T) {
	nums := []int{5, 7, 8, 4, 12, 1}
	heapSort(nums)
	fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
	nums := []int{5, 7, 8, 4, 12, 1}
	ret := mergeSort(nums)
	fmt.Println(ret)
}
