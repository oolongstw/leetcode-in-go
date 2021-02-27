package main

import "fmt"

// 解题思路
// 利用map空间换时间，遍历一次数组即可
// `find + val = target` val 是nums当前元素值，find是需要寻找的元素
// 如果map[find]找到则返回{当前索引, map[find]}，否则map[find] = 当前元素索引
func twoSum(nums []int, target int) []int {
	if len(nums) < 1 {
		return nil
	}
	mp := map[int]int{}
	for i, val := range nums {
		find := target - val // find + val = target
		if v, ok := mp[val]; !ok {
			mp[find] = i
		} else {
			return []int{v, i}
		}
	}
	return nil
}

func main() {
	nums := []int{1, 2, 3, 4}
	target := 7
	fmt.Println(twoSum(nums, target))
}
