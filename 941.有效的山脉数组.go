/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */
package main

// @lc code=start
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	left, right := 0, len(A)-1
	for left+1 < len(A)-1 && A[left] < A[left+1] {
		left++
	}
	for right-1 > 0 && A[right-1] > A[right] {
		right--
	}

	return left == right
}

// @lc code=end
