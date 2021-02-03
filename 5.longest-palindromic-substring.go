/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (32.51%)
 * Total Accepted:    456.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a"
 * 输出："a"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "ac"
 * 输出："a"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由数字和英文字母（大写和/或小写）组成
 *
 *
 */
package main

import "fmt"

func longestPalindrome(str string) string {
	size := len(str)
	if size <= 1 {
		return str
	}
	dp := make([][]bool, size)
	for i, _ := range dp {
		dp[i] = make([]bool, size)
	}

	ans := ""
	for x := 1; x <= size; x++ {
		for i := 0; i <= size-x; i++ {
			j := i + x - 1
			if j == i {
				dp[i][j] = true
			} else {
				if j-i == 1 {
					dp[i][j] = (str[i] == str[j])
				} else {
					dp[i][j] = dp[i+1][j-1] && (str[i] == str[j])
				}
			}

			if dp[i][j] && x > len(ans) {
				ans = str[i : j+1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome("cxbbx"))

}
