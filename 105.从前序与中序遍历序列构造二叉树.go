/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}
	if len(preorder) == 0 {
		return nil
	}

	//找到根节点在中序遍历中的index
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}

	//确定root节点
	root := &TreeNode{Val: preorder[0]}
	//确定左子树在前序便利中的位置  [1 :  (index + 1)] 在中序遍历中的位置 [0 : index]
	root.Left = buildTree(preorder[1:(index+1)], inorder[:index])
	//确定右子树在前序遍历中的位置[index+1 : ] 在中序遍历中的位置[index+1 :]
	root.Right = buildTree(preorder[(index+1):], inorder[index+1:])
	return root
}

// @lc code=end
