/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) == 0 {
		return false
	}
	if len(s1) == 0 {
		return true
	}

	//构建需要的个数map
	need := make(map[rune]int)
	for _, v := range s1 {
		need[v]++
	}

	//统计window中的个数
	window := make(map[rune]int)
	left, right := 0, 0
	//有效值
	valid := 0
	for right < len(s2) {
		// 移入字符到窗口
		x := s2[right]
		//扩大right窗口
		right++
		if need[rune(x)] > 0 {
			window[rune(x)]++
			if window[rune(x)] == need[rune(x)] {
				valid++
			}
		}

		//滑动窗口条件，就是窗口串大小正好比目标子串大的时候
		for right-left >= len(s1) {
			// 如果需求条件满足，证明有次子串
			if valid == len(need) {
				return true
			}
			//将要移出的字符
			y := s2[left]
			//移动左边界
			left++
			//窗口内数据更新
			if need[rune(y)] > 0 {
				if need[rune(y)] == window[rune(y)] {
					valid--
				}
				window[rune(y)]--
			}
		}
	}

	return false
}

// @lc code=end

