package main

import (
	"fmt"
	. "leetcode/Structures"
)

// 利用栈进行迭代的进行二叉树的中序遍历
// 递归本质上使用了系统的栈
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	res := []int{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root) // 模拟递归调用 不断往左子树方向走，每走一次就将当前节点保存到栈中
			root = root.Left
		}
		top := stack[len(stack)-1]   // 栈顶元素
		res = append(res, top.Val)   // 打印node value
		stack = stack[:len(stack)-1] // pop stack
		root = top.Right             // 左边遍历完 遍历右结点
	}
	return res
}

// 先序遍历 非递归版 root -> left -> right
func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]   // top node
		stack = stack[:len(stack)-1] // pop
		root = top.Right             // 转向右子树
	}
	return res
}

func main() {
	ret := []int{}
	fmt.Println(ret == nil)
}
