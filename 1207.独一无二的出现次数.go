/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int)
	for _, value := range arr {
		if _, ok := counts[value]; ok {
			counts[value]++
		} else {
			counts[value] = 1
		}
	}
	countsMap := make(map[int]int)
	for _, cnt := range counts {
		if _, ok := countsMap[cnt]; ok {
			return false
		} else {
			countsMap[cnt] = 1
		}
	}
	return true
}

// @lc code=end

