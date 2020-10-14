/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonChars([]string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}))
}

// @lc code=start
func commonChars(A []string) []string {
	minFreq := [26]int{}
	for index, _ := range minFreq {
		minFreq[index] = 101
	}
	for _, word := range A {
		freq := [26]int{}
		for _, r := range word {
			freq[r-'a']++
		}
		for i, _ := range freq {
			if freq[i] < minFreq[i] {
				minFreq[i] = freq[i]
			}
		}
	}
	result := make([]string, 0)
	for i, v := range minFreq {
		for c := 0; c < v; c++ {
			result = append(result, string('a'+i))
		}
	}
	return result
}

// @lc code=end
