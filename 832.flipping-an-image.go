package main

/*
 * @lc app=leetcode.cn id=832 lang=golang
 *
 * [832] 翻转图像
 *
 * https://leetcode-cn.com/problems/flipping-an-image/description/
 *
 * algorithms
 * Easy (75.55%)
 * Total Accepted:    54.7K
 * Total Submissions: 70K
 * Testcase Example:  '[[1,1,0],[1,0,1],[0,0,0]]'
 *
 * 给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。
 *
 * 水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。
 *
 * 反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[[1,1,0],[1,0,1],[0,0,0]]
 * 输出：[[1,0,0],[0,1,0],[1,1,1]]
 * 解释：首先翻转每一行: [[0,1,1],[1,0,1],[0,0,0]]；
 * ⁠    然后反转图片: [[1,0,0],[0,1,0],[1,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：[[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
 * 输出：[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 * 解释：首先翻转每一行: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]]；
 * ⁠    然后反转图片: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */
func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		A[i] = reverseAndFlip(A[i])
	}
	return A
}

func reverseAndFlip(data []int) []int {
	i, j := 0, len(data)-1
	for i <= j {
		if i == j {
			data[i] = flip(data[i])
		} else {
			data[i], data[j] = flip(data[j]), flip(data[i])
		}
		i++
		j--
	}
	return data
}

func flip(data int) int {
	if data == 1 {
		return 0
	} else {
		return 1
	}
}
