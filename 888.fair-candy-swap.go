/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 *
 * https://leetcode-cn.com/problems/fair-candy-swap/description/
 *
 * algorithms
 * Easy (55.47%)
 * Total Accepted:    24.3K
 * Total Submissions: 40.2K
 * Testcase Example:  '[1,1]\n[2,2]'
 *
 * 爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 根糖果棒的大小，B[j] 是鲍勃拥有的第 j 根糖果棒的大小。
 *
 * 因为他们是朋友，所以他们想交换一根糖果棒，这样交换后，他们都有相同的糖果总量。（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）
 *
 * 返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1] 是 Bob 必须交换的糖果棒的大小。
 *
 * 如果有多个答案，你可以返回其中任何一个。保证答案存在。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：A = [1,1], B = [2,2]
 * 输出：[1,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：A = [1,2], B = [2,3]
 * 输出：[1,2]
 *
 *
 * 示例 3：
 *
 *
 * 输入：A = [2], B = [1,3]
 * 输出：[2,3]
 *
 *
 * 示例 4：
 *
 *
 * 输入：A = [1,2,5], B = [2,4]
 * 输出：[5,4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * 1
 * 1
 * 保证爱丽丝与鲍勃的糖果总量不同。
 * 答案肯定存在。
 *
 *
 */
func fairCandySwap(A []int, B []int) []int {
	m := make(map[int]int)
	sumA, sumB := 0, 0
	for i, v := range A {
		sumA += v
		sumA[v] = i
	}
	for _, v := range B {
		sumB += v
	}
	delta := (sumA - sumB) / 2
	for _, v := range B {
		y := v
		x := y + delta
		if _, ok := m[x]; ok {
			return []int{x, y}
		}
	}
	return nil
}
