/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	num := x
	if  num < 0 {
		return false
	}
	result := 0
	for {
		if num == 0 {
			break
		}
		result = result * 10 + num % 10
		num /= 10
	}
	return result == x
}

// @lc code=end

