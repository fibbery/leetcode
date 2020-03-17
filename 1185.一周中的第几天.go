/*
 * @lc app=leetcode.cn id=1185 lang=golang
 *
 * [1185] 一周中的第几天
 */

// @lc code=start
func dayOfTheWeek(day int, month int, year int) string {
	date := []string{ "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday","Sunday"}
	if month == 1 || month == 2 {
		month += 12
		year -= 1
	}
	week := (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400) % 7;
	return date[week]
}
// @lc code=end

