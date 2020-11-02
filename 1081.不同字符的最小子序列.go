/*
 * @lc app=leetcode.cn id=1081 lang=golang
 *
 * [1081] 不同字符的最小子序列
 */

// @lc code=start
func smallestSubsequence(s string) string {
	counts := make([]int, 26)
	for _, v := range s {
		counts[v-'a']++
	}

	stack := NewStack()
	for _, r := range s {
		counts[r-'a']--
		if stack.Contains(r) {
			continue
		}
		for stack.Last() >= r {
			//如果后续继续出现，则可以pop
			last := stack.Last()
			if counts[last-'a'] != 0 {
				stack.PopLast()
				continue
			}
			break
		}
		stack.Push(r)
	}
	return string(stack.GetData())
}

type Stack struct {
	data []rune
	m    map[rune]int
}

func NewStack() *Stack {
	return &Stack{
		m: make(map[rune]int),
	}
}

func (s *Stack) Push(val rune) {
	s.data = append(s.data, val)
	if _, ok := s.m[val]; ok {
		s.m[val]++
	} else {
		s.m[val] = 1
	}
}

func (s *Stack) PopLast() rune {
	if len(s.data) > 0 {
		v := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		s.m[v]--
		if s.m[v] == 0 {
			delete(s.m, v)
		}
		return v
	}
	return -1
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Last() int32 {
	if s.Size() > 0 {
		return s.data[len(s.data)-1]
	}
	return -1
}

func (s *Stack) GetData() []rune {
	return s.data
}

func (s *Stack) Contains(val rune) bool {
	if _, ok := s.m[val]; ok {
		return true
	}
	return false
}

// @lc code=end

