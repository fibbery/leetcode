/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	var current, max = 0, 0
	for _, value := range nums{
		if value == 1 {
			current ++
		}else{
			max = Max(current, max)
			current = 0
		}
	}
	return Max(current, max)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

