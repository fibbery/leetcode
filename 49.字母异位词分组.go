/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) (ret [][]string) {
	cntsMap := make(map[[26]int][]string)
	for _, s := range strs {
		cnts := [26]int{}
		for _, v := range s {
			cnts[v-'a']++
		}
		cntsMap[cnts] = append(cntsMap[cnts], s)
	}

	for _, values := range cntsMap {
		ret = append(ret, values)
	}
	return
}

// @lc code=end

