/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	current := head
	for l1 != nil || l2 != nil {
		num := current.Val
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		current.Val = num % 10
		if l1 != nil || l2 != nil || num/10 > 0 {
			current.Next = &ListNode{Val: num/10}
			current = current.Next
		}
	}
	return head
}
// @lc code=end

