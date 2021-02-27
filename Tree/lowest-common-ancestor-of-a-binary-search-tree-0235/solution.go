package main

import (
	. "leetcode/Structures"
)

// 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 利用BST特性 左子树节点的值都小于根节点 右子树节点的值都大于根节点
	// 根据节点值大小遍历不同的子树
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
