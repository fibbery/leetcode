/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			tmp := head
			for tmp != slow {
				slow = slow.Next
				tmp = tmp.Next
			}
			return tmp
		}
	}
	return nil
}

// @lc code=end

