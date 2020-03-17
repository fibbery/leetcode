/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]bool)
	for _, value := range J {
		jewels[value] = true
	}
	count := 0
	for _, value := range S {
		_, ok := jewels[value]
		if ok {
			count ++
		}
	}
	return count
}
// @lc code=end

