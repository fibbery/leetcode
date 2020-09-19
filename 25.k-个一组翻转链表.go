/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prev, end := dummy, dummy
	for end.Next != nil {
		for index := k; index > 0 && end != nil; index-- {
			end = end.Next
		}
		if end == nil {
			break
		}
		//临时存储 准备翻转区域的起点以及未翻转区域的起点
		start, next := prev.Next, end.Next
		//对准备翻转区域的数据进行翻转
		end.Next = nil
		prev.Next = reverse(start)
		start.Next = next
		prev, end = start, start
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}
// @lc code=end

