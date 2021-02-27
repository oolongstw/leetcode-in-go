package main

import (
	. "leetcode/Structures"
)

// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	head := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// l1 还有剩余的节点
	if l1 != nil {
		cur.Next = l1
	}
	// l2 还有剩余的节点
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}
