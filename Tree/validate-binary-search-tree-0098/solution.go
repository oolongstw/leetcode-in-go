package main

import (
	. "leetcode/Structures"
	"math"
)

// 解题思路: 中序遍历
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := []int{}
	walk(root, &res)
	for i := 0; i < len(res)-1; i++ {
		if res[i+1] <= res[i] {
			return false
		}
	}
	return true
}

// left -> root -> right
func walk(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		walk(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root != nil {
		walk(root.Right, res)
	}
}

// 方法2: 在遍历的过程中比较左右子树节点和根节点的大小
func valid(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
