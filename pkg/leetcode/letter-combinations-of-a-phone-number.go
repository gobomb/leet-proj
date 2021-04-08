package leetcode

var m = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return find(0, digits, "")
}

func find(index int, digits string, s string) (rs []string) {
	if index >= len(digits) {
		return []string{s}
	}
	t := digits[index : index+1]
	for i := range m[t] {
		sappend := s + m[t][i]
		rs = append(rs, find(index+1, digits, sappend)...)
	}
	return
}
