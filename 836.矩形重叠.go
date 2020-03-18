/*
 * @lc app=leetcode.cn id=836 lang=golang
 *
 * [836] 矩形重叠
 * 思路：将矩形重叠转换成一维线段不交叠，即两个矩形投影到X，Y轴的线段不交叠，就可以知道两矩形是否重叠
 * 矩形1 的 X轴投影范围是 rec1[0]，rec2[2]
 * 矩形2 的 X轴投影范围是 rec2[0]，rec2[2]
 * 那么在X轴投影重叠的条件为  max(rec1[0]，rec[0]) < min(rec1[2]，rec2[2])
 * 同理Y轴也是如此
 */

// @lc code=start
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	xOverlap := max(rec1[0],rec2[0]) < min(rec1[2],rec2[2])
	yOverlap := max(rec1[1],rec2[1]) < min(rec1[3],rec2[3])
	return xOverlap && yOverlap
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
// @lc code=end

