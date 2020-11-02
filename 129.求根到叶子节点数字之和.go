/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, presum int) int {
	if node == nil {
		return 0
	}

	presum = presum*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return presum
	}

	return dfs(node.Left, presum) + dfs(node.Right, presum)
}

// @lc code=end

