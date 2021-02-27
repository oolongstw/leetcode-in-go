package main

import (
	. "leetcode/Structures"
	"strconv"
)

var res []string

// 二叉树的所有路径: 返回所有从根节点到叶子节点的路径。
func binaryTreePaths(root *TreeNode) []string {
	res = []string{}
	if root == nil {
		return res
	}
	dfs(root, "")
	return res
}

func dfs(root *TreeNode, str string) {
	str += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, str)
	} else {
		if root.Left != nil {
			dfs(root.Left, str+"->")
		}
		if root.Right != nil {
			dfs(root.Right, str+"->")
		}
	}
}
