/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 * 思路：回文串就是指顺看逆看都是一样顺序的字符串，形如 abba 或 abcba
 *		 先把字符出现的次数统计起来(map)，那么去循环map的时候可以知道每次回文
 *		 字符串增加的长度为 (字母出现次数) / 2 * 2， 如果一个字母出现五次，那么他提供的长度就是4。
 *		 同时意味着回文字符串长度此时都为偶数。
 *		 如果第一次发现一个字母出现奇数次，那么可以让这个字母成为对称中心的字符，可以让回文字符串长度增加1
 *		 但是如果再次遇到出现奇数次的字符，我们就不需要改变他的长度了，否则就无法完成回文对称的操作了
 */

// @lc code=start
func longestPalindrome(s string) int {
	dict := make(map[rune]int)
	for _, value := range s {
		dict[value] ++
	}
	length := 0
	for _, count := range dict {
		length += count / 2 * 2
		if count % 2 == 1 && length % 2 == 0 {
			length ++
		}
	}
	return length
}
// @lc code=end

