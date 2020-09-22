/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
// package main

// type TreeNode struct {
// 	Val         int
// 	Left, Right *TreeNode
// }

func flatten(root *TreeNode) {
	var preorder func(node *TreeNode) []*TreeNode
	preorder = func(node *TreeNode) (nodes []*TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		nodes = append(nodes, preorder(node.Left)...)
		nodes = append(nodes, preorder(node.Right)...)
		return
	}
	data := preorder(root)
	for i := 0; i < len(data)-1; i++ {
		data[i].Left, data[i].Right = nil, data[i+1]
	}
}

// @lc code=end
