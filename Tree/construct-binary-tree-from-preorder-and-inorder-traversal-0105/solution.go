package main

import (
	. "leetcode/Structures"
)

// 根据一棵树的前序遍历与中序遍历构造二叉树
// 解题思路: 递归 找到根节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// preorder: root -> left -> right, inorder: left -> root -> right
	root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}
	var pos int
	for i, val := range inorder {
		if val == preorder[0] {
			pos = i // 中序遍历找到根节点 根节点左边为左子树 右边为右子树
			break
		}
	}

	// 区分左右子树的前序遍历 中序遍历序列
	leftInOrder := inorder[0:pos]
	leftPreOrder := preorder[1 : 1+len(leftInOrder)]
	rightInOrder := inorder[pos+1:]
	rightPreOrder := preorder[1+len(leftInOrder):]
	root.Left = buildTree(leftPreOrder, leftInOrder)
	root.Right = buildTree(rightPreOrder, rightInOrder)

	return root
}

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	buildTree(pre, in)
}
