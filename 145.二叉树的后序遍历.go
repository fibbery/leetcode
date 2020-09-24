/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(node *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	ret := make([]int, 0)
	var previous *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if (node.Left == nil && node.Right == nil) || (node.Right == nil && previous == node.Left) || (previous == node.Right) {
			ret = append(ret, node.Val)
			previous = node
			stack = stack[:len(stack)-1]
			node = nil
		} else {
			node = node.Right
		}
	}
	return ret
}

// @lc code=end

