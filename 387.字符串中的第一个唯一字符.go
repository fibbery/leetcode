/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	m := make(map[int32]int)
	for _, v := range s {
		m[v-'a']++
	}
	for i, v := range s {
		if m[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

