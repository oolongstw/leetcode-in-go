package main

import (
	. "leetcode/Structures"
)

var (
	node int
)

// 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	var cnt int
	inorder(root, k, &cnt)
	return node
}

// 二叉树中序遍历
func inorder(root *TreeNode, k int, cnt *int) {
	if root == nil {
		return
	}
	if *cnt == k {
		// 已经满足条件的直接返回
		return
	}
	inorder(root.Left, k, cnt)
	*cnt++
	if *cnt == k {
		node = root.Val
	}
	inorder(root.Right, k, cnt)
}
