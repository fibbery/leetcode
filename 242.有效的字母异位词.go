/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	chars := [26]int{}
	for _, value := range t{
		chars[value - 'a'] ++
	}

	for _, value := range s {
		if chars[value - 'a'] == 0 {
			return false
		}
		chars[value - 'a'] --
	}
	return true
}
// @lc code=end

