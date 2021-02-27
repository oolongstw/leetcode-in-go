package main

import (
	"fmt"
	. "leetcode/Structures"
)

// 解题思路: 递归 DFS 节点不为空时分别求左右子树的高度的最大值
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	depth := max(l, r)
	return depth + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	node := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
	}
	fmt.Println(maxDepth(node))
}
