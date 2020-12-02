/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	cnt := make(map[int]int)
	ans := 0
	for _, a := range A {
		for _, b := range B {
			cnt[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			ans = ans + cnt[-c-d]
		}
	}
	return ans
}

// @lc code=end

