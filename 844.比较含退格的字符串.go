/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */
package main

func main() {
	backspaceCompare("y#fo##f", "y#f#o##f")
}

// @lc code=start
func backspaceCompare(S string, T string) bool {
	array_a := getArray(S)
	array_b := getArray(T)
	if len(array_a) != len(array_b) {
		return false
	}
	for i, _ := range array_a {
		if array_a[i] != array_b[i] {
			return false
		}
	}
	return true
}

func getArray(S string) []rune {
	array := make([]rune, 0)
	for _, v := range S {
		if v != '#' {
			array = append(array, v)
			continue
		}

		if len(array) > 0 {
			array = array[:len(array)-1]
		}
	}
	return array
}

// @lc code=end
