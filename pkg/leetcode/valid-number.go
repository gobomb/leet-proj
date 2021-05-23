package leetcode

const (
	INVALID int = iota
	SPACE
	SIGN
	DIGIT
	DOT
	EXPONENT
)

// ref: https://zhuanlan.zhihu.com/p/20042325`
var transitionTable = [][]int{{-1, 0, 3, 1, 2, -1}, //0 no input or just spaces
	{-1, 8, -1, 1, 4, 5},    //1 input is digits
	{-1, -1, -1, 4, -1, -1}, //2 no digits in front just Dot
	{-1, -1, -1, 1, 2, -1},  //3 sign
	{-1, 8, -1, 4, -1, 5},   //4 digits and dot in front
	{-1, -1, 6, 7, -1, -1},  //5 input 'e' or 'E'
	{-1, -1, -1, 7, -1, -1}, //6 after 'e' input sign
	{-1, 8, -1, 7, -1, -1},  //7 after 'e' input digits
	{-1, 8, -1, -1, -1, -1}} //8 after valid input input space

func isNumber(s string) bool {
	state := INVALID

	for i := range s {
		c := s[i]
		tp := INVALID
		switch c {
		case ' ':
			tp = SPACE
		case '+', '-':
			tp = SIGN
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			tp = DIGIT
		case '.':
			tp = DOT
		case 'e', 'E':
			tp = EXPONENT
		}
		state = transitionTable[state][tp]
		if state == -1 {
			return false
		}
	}
	switch state {
	case 1, 4, 7, 8:
		return true
	}
	return false
}
