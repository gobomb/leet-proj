package leetcode

func isMatchwildcard(s string, p string) bool {
	matchmap = make([][]state, len(s))
	for i := range matchmap {
		matchmap[i] = make([]state, len(p))
	}

	return checkMatchWC(s, p, 0, 0)
}

func checkMatchWC(s, p string, si, pi int) (b bool) {
	// fmt.Printf("2  %v %v %v %v \n", si, pi, len(s), len(p))
	// defer func() {
	// 	fmt.Println(b)
	// }()
	if si == len(s) && pi == len(p) {
		return true
	}
	if pi == len(p) && si < len(s) {
		// p最后一位是*
		// 需考虑越界
		if pi-1 >= 0 && p[pi-1] == '*' {
			return true
		}
		return false
	}
	// 考虑多个 *
	if si == len(s) && pi < len(p) {
		if p[pi] != '*' {
			return false
		} else {
			return checkMatchWC(s, p, si, pi+1)
		}
	}

	if matchmap[si][pi] == notcheck {
		if p[pi] != '*' {
			preCheck := (s[si] == p[pi] || p[pi] == '?')
			if preCheck && checkMatchWC(s, p, si+1, pi+1) {
				matchmap[si][pi] = match
			} else {
				matchmap[si][pi] = notmatch
			}
		} else {
			if pi+1 == len(p) {
				// p最后一位是*
				matchmap[si][pi] = match
			} else if p[pi+1] == '*' {
				// 需考虑到连续的 *
				if checkMatchWC(s, p, si, pi+1) {
					matchmap[si][pi] = match
				} else {
					matchmap[si][pi] = notmatch
				}
			} else {
				// 注意匹配 ？
				if (s[si] == p[pi+1] || p[pi+1] == '?') && checkMatchWC(s, p, si+1, pi+2) || checkMatchWC(s, p, si+1, pi) {
					matchmap[si][pi] = match
				} else {
					matchmap[si][pi] = notmatch
				}
			}
		}
	}
	return matchmap[si][pi] == match
}
