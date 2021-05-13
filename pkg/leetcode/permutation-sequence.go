package leetcode

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func getPermutation(n int, k int) string {
	rs := ""
	if n == 0 {
		return rs
	}
	for i := 1; i < n+1; i++ {
		rs += string('0' + byte(i))
	}
	// log.Printf("%v\n", rs)
	rs = dfsFindKPermut(rs, k)

	return rs
}

func factorial(n int) int {
	f := 1
	for i := 1; i < n+1; i++ {
		f *= i
	}
	return f
}

func dfsFindKPermut(s string, k int) string {
	// log.Printf("args %v %v", s, k)
	l := len(s)
	if l == 1 {
		return s
	}
	// f := factorial(l)
	nextf := factorial(l - 1)

	if k <= nextf {
		return s[0:1] + dfsFindKPermut(s[1:], k)
	}
	id := (k - 1) / nextf
	remainder := (k-1)%nextf + 1
	log.Printf("id remainder %v %v\n", id, remainder)
	return s[id:id+1] + dfsFindKPermut(s[:id]+s[id+1:], remainder)
}
