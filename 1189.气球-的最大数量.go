/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	bCount, aCount, lCount, oCount, nCount := 0, 0, 0, 0, 0
	for _, value := range text {
		switch value {
		case 'b':
			bCount++
		case 'a':
			aCount++
		case 'l':
			lCount++
		case 'o':
			oCount++
		case 'n':
			nCount++
		}
	}
	lCount /= 2
	oCount /= 2
	counts := []int{bCount, aCount, lCount, oCount, nCount}
	min := counts[0]
	for i := 1; i < len(counts); i++ {
		if counts[i] < min {
			min = counts[i]
		}
	}
	return min
}
// @lc code=end

