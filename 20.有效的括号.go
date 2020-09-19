/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if s == "" {
		return false
	}
	var stack []uint8
	m := map[uint8]uint8{
		'}': '{',
		')': '(',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}

		if m[s[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack) -1]
		}
	}
	return len(stack) == 0
}
// @lc code=end

