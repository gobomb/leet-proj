package easy

import "sort"

/*
	455. Assign Cookies
*/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i := 0

	for j := 0; i < len(g) && j < len(s); j = j + 1 {
		if g[i] <= s[j] {
			i++
		}
	}

	return i
}
