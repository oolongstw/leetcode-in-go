package main

import (
	. "leetcode/Structures"
)

// 翻转二叉树
// 解题思路 递归交换 左右子节点
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}
