/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归算法
// func inorderTraversal(root *TreeNode) []int {
// 	nums := make([]int, 0)
// 	var inorder func(node *TreeNode)
// 	inorder = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		inorder(node.Left)
// 		nums = append(nums, node.Val)
// 		inorder(node.Right)
// 	}
// 	inorder(root)
// 	return nums
// }

//使用morris算法
// func inorderTraversal(root *TreeNode) []int {
// 	ret := make([]int, 0)
// 	for root != nil {
// 		if root.Left != nil {
// 			predecessor := root.Left
// 			for predecessor.Right != nil && predecessor.Right != root {
// 				predecessor = predecessor.Right
// 			}

// 			if predecessor.Right == nil {
// 				predecessor.Right = root
// 				root = root.Left
// 			} else {
// 				ret = append(ret, root.Val)
// 				predecessor.Right = nil
// 				root = root.Right
// 			}

// 		} else {
// 			ret = append(ret, root.Val)
// 			root = root.Right
// 		}
// 	}
// 	return ret
// }

//堆栈算法
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	ret := make([]int, 0)
	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			pop := stack[len(stack)-1]
			ret = append(ret, pop.Val)
			stack = stack[:len(stack)-1]
			node = pop.Right
		}
	}
	return ret
}

// @lc code=end

