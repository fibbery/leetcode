/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
package main

// @lc code=start
func plusOne(digits []int) []int {
	preadd := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+preadd > 9 {
			digits[i] = (digits[i] + preadd) % 10
			addon = 1
		} else {
			digits[i] = digits[i] + preadd
			addon = 0
		}
	}
	if addon == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end
