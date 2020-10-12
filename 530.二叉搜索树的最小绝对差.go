/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 */

package main

import "github.com/ethereum/go-ethereum/common/math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	result, pre := math.MaxInt64, -1
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if pre != -1 && node.Val-pre < result {
			result = node.Val - pre
		}
		pre = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return result
}

// @lc code=end
