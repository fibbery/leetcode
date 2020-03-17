/*
 * @lc app=leetcode.cn id=1160 lang=golang
 *
 * [1160] 拼写单词
 * 思路：通过建立map映射表来处理，先把chars里面各个字符作为key,出现频次作为value。
 * 		 循环words的时候，每个word也建立一个map映射，然后循环该map，通过看频次是不是不大于chars里面出现的频次来
 *		 来确定是不是包含这个word
 */

// @lc code=start
func countCharacters(words []string, chars string) int {
	charsCount := make(map[rune]int)
	for _, value := range chars {
		charsCount[value]++
	}
	length := 0
	for _, word := range words {
		wordCount := make(map[rune]int)
		for _, value := range word {
			wordCount[value] ++
		}
		ok := true
		for key, value := range wordCount {
			if value > charsCount[key] {
				ok = false
				break
			}
		}
		if ok {
			length += len(word)
		}
	}
	return length
}
// @lc code=end

