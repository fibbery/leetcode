/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareNode(root.Left, root.Right)
}

func compareNode(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	isEqual := compareNode(left.Left, right.Right) && compareNode(left.Right, right.Left)
	return isEqual && (left.Val == right.Val)
}

// @lc code=end
