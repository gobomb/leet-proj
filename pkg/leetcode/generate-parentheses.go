package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	rss := []string{}
	var genParenth func(l, r int, rs string)
	genParenth = func(l, r int, rs string) {
		if l == 0 && r == 0 {
			// fmt.Printf("%v\n", rs)
			rss = append(rss, rs)
			return
		}
		if l != 0 {
			genParenth(l-1, r, rs+"(")
		}
		if r > l {
			genParenth(l, r-1, rs+")")
		}
	}

	genParenth(n, n, "")
	return rss
}
