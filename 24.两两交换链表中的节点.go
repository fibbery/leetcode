/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prev, end := dummy, dummy
	for end.Next != nil {
		for i := 0 ; i < 2 && end != nil; i++{
			end  = end.Next
		}
		if end == nil{
			break
		}
		start, next := prev.Next, end.Next
		end.Next = nil
		prev.Next = reverse(start)
		start.Next = next
		prev, end = start, start
	}
	return dummy.Next
}
// @lc code=end

