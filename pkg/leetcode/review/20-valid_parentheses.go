package review

import (
	"log"
	// "justest/pkg/ds"
	// "sync"
)

// 20 valid_parentheses

func isValid(s string) bool {
	m := make(map[byte]byte)
	m['{'] = '}'
	m['('] = ')'
	m['['] = ']'

	st := []byte{}

	for i := 0; i < len(s); i++ {
		log.Printf("%s", st)
		if t, ok := m[s[i]]; ok {
			st = append(st, t)
		} else if len(st) != 0 && s[i] == st[len(st)-1] {
			st = st[:len(st)-1]
		} else {
			return false
		}
	}

	return len(st) == 0
}
