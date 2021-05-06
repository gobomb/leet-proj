package leetcode

/*
	ref: https://jaycechant.info/2020/leetcode-10-regular-expression-matching/

	𝐷𝑃𝑖𝑠,𝑖𝑝=𝑚𝑎𝑡𝑐ℎ(𝑠[𝑖𝑠],𝑝[𝑖𝑝])∧𝐷𝑃𝑖𝑠+1,𝑖𝑝+1 				 (𝑝[𝑖𝑝+1] 𝑖𝑠 𝑛𝑜𝑡 ∗)
	𝐷𝑃𝑖𝑠,𝑖𝑝=𝐷𝑃𝑖𝑠,𝑖𝑝+2∨(𝑚𝑎𝑡𝑐ℎ(𝑠[𝑖𝑠],𝑝[𝑖𝑝])∧𝐷𝑃𝑖𝑠+1,𝑖𝑝)		(𝑝[𝑖𝑝+1] 𝑖𝑠 ∗)
	b a*b
	aaab a*b
*/

type state int

const (
	notcheck state = iota
	match
	notmatch
)

var matchmap [][]state

func isMatch(s string, p string) bool {
	matchmap = make([][]state, len(s)+1)
	for i := range matchmap {
		matchmap[i] = make([]state, len(p)+1)
		matchmap[i][len(p)] = notmatch
	}
	matchmap[len(s)][len(p)] = match
	return checkMatch(s, p, 0, 0)
}

func checkMatch(s, p string, is, ip int) bool {
	// when ip == len(p), matchmap[*][ip]==notcheck(*!=len(s))
	if matchmap[is][ip] == notcheck {
		// when is == len(s), preCheck == false
		preCheck := is < len(s) && (s[is] == p[ip] || p[ip] == '.')

		// 注意这里的判断
		if ip+1 < len(p) && p[ip+1] == '*' {
			if checkMatch(s, p, is, ip+2) || (preCheck && checkMatch(s, p, is+1, ip)) {
				matchmap[is][ip] = match
			} else {
				matchmap[is][ip] = notmatch
			}
		} else { // ip+1>=len(p) || p[ip+1] !='*'
			if preCheck && checkMatch(s, p, is+1, ip+1) {
				matchmap[is][ip] = match
			} else {
				matchmap[is][ip] = notmatch
			}
		}
	}
	return matchmap[is][ip] == match
}
