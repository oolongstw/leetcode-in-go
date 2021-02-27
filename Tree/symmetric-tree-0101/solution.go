package main

import (
	. "leetcode/Structures"
)

// 给定一个二叉树，检查它是否是镜像对称的。
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return search(root.Left, root.Right)
}

func search(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return search(left.Left, right.Right) && search(left.Right, right.Left)
}
