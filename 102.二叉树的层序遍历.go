/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
type LevelTreeNode struct {
	*TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]LevelTreeNode, 1)
	queue[0] = LevelTreeNode{
		TreeNode: root,
		Level:    0,
	}
	result := make([][]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Level >= len(result) {
			result = append(result, make([]int, 0))
		}
		result[node.Level] = append(result[node.Level], node.Val)
		if node.Left != nil {
			queue = append(queue, LevelTreeNode{
				TreeNode: node.Left,
				Level:    node.Level + 1,
			})
		}
		if node.Right != nil {
			queue = append(queue, LevelTreeNode{
				TreeNode: node.Right,
				Level:    node.Level + 1,
			})
		}
	}
	return result
}

// @lc code=end

