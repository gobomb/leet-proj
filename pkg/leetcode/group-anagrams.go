package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	strmap := make(map[string][]string)
	for i := range strs {
		s := SortString(strs[i])
		sli := strmap[s]
		sli = append(sli, strs[i])
		strmap[s] = sli
	}
	rs := [][]string{}
	for _, v := range strmap {
		rs = append(rs, v)
	}
	return rs
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
