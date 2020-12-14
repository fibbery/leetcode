/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int)
	max := 0
	left, right := 0, 0
	for right < len(s) {
		x := s[right]
		right++
		window[x]++
		for window[x] > 1 {
			//数据重复的时候滑动窗口
			y := s[left]
			left++
			window[y]--
		}
		max = maxInt(max, right-left)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

