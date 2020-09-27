/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key == root.Val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 找到左子树的最大值，或者右子树的最小值，进行替换
		min := getMinNode(root.Right)
		root.Val = min.Val
		root.Right = deleteNode(root.Right, min.Val)

	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func getMinNode(n *TreeNode) *TreeNode {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

// @lc code=end
