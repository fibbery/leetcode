/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, H int) int {
	left, right := 1, getMaxSpeed(piles)
	for left <= right {
		mid := left + (right-left)/2
		if canFinish(piles, mid, H) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func canFinish(piles []int, speed, h int) bool {
	for _, v := range piles {
		h = h - int(math.Ceil(float64(v)/float64(speed)))
		if h < 0 {
			return false
		}
	}
	return true
}

func getMaxSpeed(piles []int) int {
	max := piles[0]
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	return max
}}

// @lc code=end

