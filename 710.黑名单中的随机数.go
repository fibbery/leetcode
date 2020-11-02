/*
 * @lc app=leetcode.cn id=710 lang=golang
 *
 * [710] 黑名单中的随机数
 */

// @lc code=start
type Solution struct {
	length int
	bMap   map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	length := N - len(blacklist)
	bMap := make(map[int]int)
	for _, v := range blacklist {
		bMap[v] = -1
	}
	last := N - 1
	for _, v := range blacklist {
		if v > length-1 {
			continue
		}
		for _, ok := bMap[last]; ok; {
			last--
			_, ok = bMap[last]
		}
		bMap[v] = last
		last--
	}
	return Solution{length: length, bMap: bMap}
}

func (this *Solution) Pick() int {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(this.length)
	if v, ok := this.bMap[index]; ok {
		return v
	}
	return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
// @lc code=end
