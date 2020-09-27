/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	max := MaxIndex(nums)
	node := &TreeNode{Val: nums[max]}
	node.Left = constructMaximumBinaryTree(nums[:max])
	node.Right = constructMaximumBinaryTree(nums[max+1:])
	return node
}

func MaxIndex(nums []int) int {
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

// @lc code=end

