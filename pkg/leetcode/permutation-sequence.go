package leetcode

import "log"

func init() {
	log.SetFlags(log.Lshortfile)
}

func getPermutation(n int, k int) string {
	rs := ""
	if n == 0 {
		return rs
	}
	s := []int{}
	for i := 1; i < n+1; i++ {
		s = append(s, i)
	}
	ss := permute(s)
	// log.Printf("%v %v\n", ss[k-1], s)
	for i := range ss[k-1] {
		rs = rs + string('0'+byte(ss[k-1][i]))
	}
	return rs
}
