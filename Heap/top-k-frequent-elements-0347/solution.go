package main

import (
	"container/heap"
	"fmt"
)

// 前 K 个高频元素
// 解题思路: 堆排序 -> 大顶堆
func topKFrequent(nums []int, k int) []int {
	// hash 记录元素出现的频率
	freq := map[int]int{}
	for _, val := range nums {
		freq[val]++
	}
	pq := &PriorityQueue{}
	heap.Init(pq) // 构建堆
	for key, count := range freq {
		heap.Push(pq, &Item{key: key, count: count})
	}
	var result []int
	for len(result) < k {
		item := heap.Pop(pq).(*Item)
		result = append(result, item.key)
	}
	return result
}

type Item struct {
	key   int
	count int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 注意: 因为golang中的heap是按最小堆组织的, 所以count越大, Less()越小，越靠近堆顶.
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	top := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return top
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}
