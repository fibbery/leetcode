/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
// func preorderTraversal(root *TreeNode) []int {
// 	data := make([]int, 0)
// 	if root == nil {
// 		return data
// 	}
// 	data = append(data, root.Val)
// 	data = append(data, preorderTraversal(root.Left)...)
// 	data = append(data, preorderTraversal(root.Right)...)
// 	return data
// }

// func preorderTraversal(node *TreeNode) []int {
// 	stack := make([]*TreeNode, 0)
// 	data := make([]int, 0)
// 	for node != nil || len(stack) > 0 {
// 		if node != nil {
// 			data = append(data, node.Val)
// 			stack = append(stack, node)
// 			node = node.Left
// 		} else {
// 			back := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			node = back.Right
// 		}
// 	}
// 	return data
// }

func preorderTraversal(node *TreeNode) []int {
	data := make([]int, 0)
	for node != nil {
		if node.Left != nil {
			predecessor := node.Left

			//寻找前驱结点
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				data = append(data, node.Val)
				predecessor.Right = node
				node = node.Left
			} else {
				node = node.Right
				predecessor.Right = nil
			}

		} else {
			data = append(data, node.Val)
			node = node.Right
		}
	}
	return data
}

// @lc code=end

