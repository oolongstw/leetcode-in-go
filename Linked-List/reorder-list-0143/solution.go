package main

import (
	. "leetcode/Structures"
)

// 重排链表
// 方法一: 遍历链表存储到 map 中，然后按照给定顺讯重新排列
// 方法二: 寻找中间节点 -> 翻转后半部分链表
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	// 快慢指针寻找中间节点 如果节点个数是偶数的话，slow 走到的是左端点，利用这一点，我们可以把奇数和偶数的情况合并，不需要分开考虑
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	reverse := reverseList(slow.Next) // 后半部分翻转
	slow.Next = nil
	// 连接链表
	for reverse != nil {
		next := head.Next
		tmp := reverse.Next
		head.Next = reverse
		reverse.Next = next
		reverse = tmp
		head = next
	}
}

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode // 所需三个节点 pre cur next

	length := 0
	for head != nil {
		next = head.Next
		head.Next = pre
		// pre 和 head 节点平移
		pre = head
		head = next
		length++
	}

	// head == nil
	return pre
}
