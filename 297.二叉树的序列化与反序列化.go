/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	codec := Constructor()
	fmt.Println(codec.serialize(root))
	node := codec.deserialize(codec.serialize(root))
	fmt.Println(node)
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	Null string
	Sep  string
}

func Constructor() Codec {
	return Codec{
		Null: "null",
		Sep:  ",",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := ""
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			result = strings.Join([]string{result, this.Null}, this.Sep)
			return
		}
		result = strings.Join([]string{result, strconv.Itoa(node.Val)}, this.Sep)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result[1:]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, this.Sep)
	var buildTree func(items *[]string) *TreeNode
	buildTree = func(items *[]string) *TreeNode {
		if len(*items) < 1 {
			return nil
		}
		v := (*items)[0]
		*items = (*items)[1:]
		if v == this.Null {
			return nil
		}
		value, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		node := &TreeNode{Val: value}
		node.Left = buildTree(items)
		node.Right = buildTree(items)
		return node
	}
	root := buildTree(&values)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end
