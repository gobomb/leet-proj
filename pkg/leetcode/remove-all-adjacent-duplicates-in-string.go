package leetcode

import (
	"container/list"
)

func removeDuplicatesFromString(s string) string {
	st := list.New()
	for i := 0; i < len(s); i++ {
		if st.Back() != nil && st.Back().Value.(byte) == s[i] {
			st.Remove(st.Back())
			continue
		}
		st.PushBack(s[i])
	}
	var ns string
	for st.Front() != nil {
		ns += string(st.Front().Value.(byte))
		st.Remove(st.Front())
	}
	return ns
}

func removeDuplicatesFromString2(s string) string {
	st := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		l := len(st)
		if l != 0 && st[l-1] == s[i] {
			st = st[:l-1]
			continue
		}
		st = append(st, s[i])
	}

	return string(st)
}
