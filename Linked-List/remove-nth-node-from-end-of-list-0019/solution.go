package main

import (
	. "leetcode/Structures"
)

// 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := &ListNode{}
	node.Next = head
	first := head
	length := 0
	for first != nil {
		length++
		first = first.Next
	}
	// 确定链表长度
	pos := length - n
	two := node
	for pos != 0 {
		pos--
		two = two.Next
	}
	two.Next = two.Next.Next // 找到需要删除的节点
	return node.Next
}

// 方法二: 只遍历一次的方法 two pointers 快慢指针
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	// 哨兵节点
	node := &ListNode{}
	node.Next = head
	// 双指针
	slow, fast := node, node
	// 两个指针被 n 个节点分开，第一个指针移到底的时候，开始执行删除操作
	// 两个指针同时向后移动
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 第二个指针没到底之前，两个指针一起后移
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return node.Next
}
