/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */

// @lc code=start
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	m := make(map[int][][]int)
	keys := make([]int, 0)
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			d := distance(r, c, r0, c0)
			if _, ok := m[d]; !ok {
				keys = append(keys, d)
				m[d] = [][]int{
					{r, c},
				}
			} else {
				m[d] = append(m[d], []int{r, c})
			}
		}
	}
	sort.Ints(keys)
	rs := make([][]int, 0)
	for _, v := range keys {
		rs = append(rs, m[v]...)
	}
	return rs
}

func distance(r1, c1, r0, c0 int) int {
	return abs(r1-r0) + abs(c1-c0)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end

