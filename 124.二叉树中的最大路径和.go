/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
import "math"

//后序遍历
func maxPathSum(root *TreeNode) int {
	sum := math.MinInt32
	var maxSum func(*TreeNode) int
	maxSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//避免选入负数节点
		left := max(maxSum(node.Left), 0)
		right := max(maxSum(node.Right), 0)
		sum = max(node.Val+left+right, sum)
		return node.Val + max(left, right)
	}
	maxSum(root)
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
