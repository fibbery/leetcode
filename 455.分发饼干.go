/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; i < len(s) && ans < len(g); i++ {
		if s[i] >= g[ans] {
			ans++
		}
	}
	return
}

// @lc code=end

