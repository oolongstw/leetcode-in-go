package main

import (
	. "leetcode/Structures"
)

// 判断一个链表是否为回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	data := []int{}
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	left, right := 0, len(data)-1
	for left < right {
		if data[left] != data[right] {
			return false
		}
		left++
		right--
	}
	return true
}
