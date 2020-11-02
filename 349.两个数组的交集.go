/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */
package main

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]struct{})
	for _, v := range nums1 {
		m1[v] = struct{}{}
	}

	m2 := make(map[int]struct{})
	for _, v := range nums2 {
		m2[v] = struct{}{}
	}

	sz := make([]int, 0)
	for k, _ := range m1 {
		if _, ok := m2[k]; ok {
			sz = append(sz, k)
		}
	}
	return sz
}

// @lc code=end
