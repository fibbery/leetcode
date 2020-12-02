/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	window, need := make(map[rune]int), make(map[rune]int)
	for _, r := range t {
		need[r]++
	}
	left, right := 0, 0
	start, length := 0, math.MaxInt32
	valid := 0

	for right < len(s) {
		r := s[right]
		if _, ok := need[rune(r)]; ok {
			window[rune(r)]++
			if window[rune(r)] == need[rune(r)] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			l := s[left]
			if _, ok := need[rune(l)]; ok {
				if window[rune(l)] == need[rune(l)] {
					valid--
				}
				window[rune(l)]--
			}
			left++
		}
		right++
	}
	if length == math.MaxInt32 {
		return ""
	} else {
		return s[start:(start + length + 1)]
	}
}

// @lc code=end

