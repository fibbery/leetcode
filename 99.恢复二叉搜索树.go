/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
 */

// @lc code=start
// package main

// import "fmt"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 中序遍历之后得到一个递增数组，找出其中不符合递增的数即置换的数
func recoverTree(root *TreeNode) {
	nums := make([]int, 0)
	var inorder func(n *TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		nums = append(nums, n.Val)
		inorder(n.Right)
	}
	inorder(root)
	x, y := findSwapValue(nums)
	recover(root, x, y, 2)
}

func findSwapValue(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recover(root *TreeNode, x, y, count int) {
	if root == nil {
		return
	}
	if root.Val == x {
		root.Val = y
		count--
	} else if root.Val == y {
		root.Val = x
		count--
	}

	if count == 0 {
		return
	}
	recover(root.Left, x, y, count)
	recover(root.Right, x, y, count)
}

// func main() {
// 	root := &TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{
// 			Val: 3,
// 			Right: &TreeNode{
// 				Val: 2,
// 			},
// 		},
// 	}
// 	recoverTree(root)
// 	fmt.Println(root)
// }

// @lc code=end
