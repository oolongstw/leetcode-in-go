package main

import (
	"fmt"
	. "leetcode/Structures"
)

// 二叉树的层序遍历
// 解题思路 利用队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	nodeNum := 1
	// 问题: 使用队列 无法区分队列中的结点来自哪一层
	for len(queue) != 0 {
		cnt := 0
		tmp := []int{}
		// 解决方案: 在每一层遍历开始前，先记录该层中的结点数量 nodeNum
		for i := 0; i < nodeNum; i++ {
			cur := queue[0]
			if cur != nil {
				fmt.Println(cur.Val)
				tmp = append(tmp, cur.Val)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				cnt++
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				cnt++
			}
			queue = queue[1:] // pop queue

		}
		nodeNum = cnt
		res = append(res, tmp)
	}

	return res
}
