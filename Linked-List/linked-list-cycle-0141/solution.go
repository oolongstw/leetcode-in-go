package main

import (
	. "leetcode/Structures"
)

// 判断链表中是否存在环
// 方法一: 快慢指针 如果有环 快指针将会追上慢指针 快指针一次前进2步，慢指针一次前进1步
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 方法二: 哈希表
func hasCycle2(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}
