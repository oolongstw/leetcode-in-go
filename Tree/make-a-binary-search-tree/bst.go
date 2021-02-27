package main

import (
	"fmt"
)

type TreeNode struct {
	value       int
	left, right *TreeNode
}

type BSTree struct {
	root *TreeNode
}

// Insert node to BST
func (bst *BSTree) Insert(val int) {
	n := &TreeNode{value: val, left: nil, right: nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func insertNode(root, newNode *TreeNode) {
	if newNode.value < root.value {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
}

// Search node in BST
func (bst *BSTree) Search(value int) (int, bool) {
	return search(bst.root, value)
}

func search(root *TreeNode, value int) (int, bool) {
	if root == nil {
		return 0, false
	}
	if value < root.value {
		return search(root.left, value)
	}
	if value > root.value {
		return search(root.right, value)
	}
	return root.value, true
}

// Remove node in BST
func (bst *BSTree) Remove(value int) {
	remove(bst.root, value)
}

func remove(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return nil
	}
	if value < root.value {
		root.left = remove(root.left, value)
	} else if value > root.value {
		root.right = remove(root.right, value)
	} else {
		// value == root.value condition
		if root.right == nil && root.left == nil {
			// 叶子节点直接删除
			root = nil
		} else if root.right != nil {
			// 被删除的节点有右子树
			root.value = findMin(root.right) // 找到右子树最小的节点 -> 替换
			root.right = remove(root.right, root.value)
		} else {
			// 被删除的节点有左子树
			root.value = findMax(root.left) // 找到左子树最大的节点 -> 替换
			root.left = remove(root.left, root.value)
		}
	}
	return root
}

func findMin(root *TreeNode) int {
	if root == nil {
		return 0
	}
	for root.left != nil {
		return findMin(root.left)
	}
	return root.value
}

func findMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	for root.right != nil {
		return findMax(root.right)
	}
	return root.value
}

func (bst *BSTree) InOrderTraverse(fn func(i interface{})) {
	inOrderTraverse(bst.root, fn)
}

func inOrderTraverse(root *TreeNode, fn func(i interface{})) {
	if root == nil {
		return
	}
	inOrderTraverse(root.left, fn)
	fn(root.value)
	inOrderTraverse(root.right, fn)
}

func main() {
	bst := &BSTree{}
	bst.Insert(1)
	bst.Insert(8)
	bst.Insert(6)
	bst.Insert(3)
	bst.Insert(10)
	bst.Insert(2)

	bst.Remove(1)
	res := []int{}
	bst.InOrderTraverse(func(i interface{}) {
		res = append(res, i.(int))
	})
	fmt.Println(res)

	fmt.Println(bst.Search(2)) // 2 true

	bst.Remove(2)
	fmt.Println(findMax(bst.root))
	fmt.Println(findMin(bst.root))
}
