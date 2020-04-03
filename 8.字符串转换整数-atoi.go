/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 * 使用有限状态机解决，参考官方解题思路，把所有可能出现的状态列举出来
 *           |   ' '    |  +/-    |  number   | other
 *   Start   |   Start  |  Signed |  Number   | End
 *   Signed  |   End    |  End    |  Number   | End
 *   Number  |   End    |  End    |  Number   | End
 *   End     |   End    |  End    |  End      | End
 */

// @lc code=start
var (
	StatusMap = map[string][4]string{
		"Start":  {"Start", "Signed", "Number", "End"},
		"Signed": {"End", "End", "Number", "End"},
		"Number": {"End", "End", "Number", "End"},
		"End":    {"End", "End", "End", "End"},
	}
	Digits = [10]string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	Max = 2 << 30 - 1
	Min = - (2 << 30)
)

type Automaton struct {
	sign   int
	answer int
	state  string
}

func newAutomaton() *Automaton {
	return &Automaton{
		sign:   1,
		answer: 0,
		state:  "Start",
	}
}

func (a *Automaton) getCol(c string) int {
	if c == " " {
		return 0
	} else if c == "+" || c == "-" {
		return 1
	} else if isDigit(c) {
		return 2
	} else {
		return 3
	}
}

func (a *Automaton) get(c string) {
	a.state = StatusMap[a.state][a.getCol(c)]
	if a.state == "Number" {
		a.answer = a.answer*10 + parseInt(c)
		if a.sign == 1 {
			a.answer = minInt(a.answer, Max)
		} else {
			a.answer = minInt(a.answer, -Min)
		}
	}

	if a.state == "Signed" {
		if c == "+" {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}

func isDigit(c string) bool {
	for _, value := range Digits {
		if value == c {
			return true
		}
	}
	return false
}

func parseInt(c string) int {
	for index, value := range Digits {
		if value == c {
			return index
		}
	}
	return 0
}


func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func myAtoi(str string) int {
	automaton := newAutomaton()
	for _, value := range str {
		automaton.get(string(value))
	}
	return automaton.sign * automaton.answer
}
// @lc code=end

