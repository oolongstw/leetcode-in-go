package main

import (
	. "leetcode/Structures"
)

// 检验 s 中是否包含和 t 具有相同结构和节点值的子树
// DFS 枚举 s 中的每一个节点，判断这个点的子树是否和 t 相等。
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true // t == nil 一定是true
	}
	if s == nil {
		return false
	}

	return isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// 判断两棵树是否相同
func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
