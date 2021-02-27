package main

import (
	. "leetcode/Structures"
)

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode // 所需三个节点 pre cur next

	for head != nil {
		next = head.Next
		head.Next = pre
		// pre 和 head 节点平移
		pre = head
		head = next
	}

	// head == nil
	return pre
}
