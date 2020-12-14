/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) (ret []int) {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, v := range p {
		need[v]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		x := s[right]
		right++
		if need[rune(x)] > 0 {
			window[rune(x)]++
			if window[rune(x)] == need[rune(x)] {
				valid++
			}
		}

		for (right - left) >= len(p) {
			if valid == len(need) {
				ret = append(ret, left)
			}
			y := s[left]
			left++
			if need[rune(y)] > 0 {
				if window[rune(y)] == need[rune(y)] {
					valid--
				}
				window[rune(y)]--
			}
		}
	}
	return
}

// @lc code=end

