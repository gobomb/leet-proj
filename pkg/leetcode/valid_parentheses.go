package leetcode

import (
	"fmt"
	// "justest/pkg/ds"
	// "sync"
)

func isValid(s string) bool {
	st := NewStack()
	pmap := make(map[byte]byte)
	pmap['['] = ']'
	pmap['('] = ')'
	pmap['{'] = '}'

	for i := range s {
		peek := st.Peek()
		if peek == nil {
			_, ok := pmap[s[i]]
			if !ok {
				return false
			}
			st.Push(s[i])
			continue
		}
		p, ok := peek.(byte)
		if !ok {
			return false
		}
		// fmt.Printf("peek %s\n", []byte{p})

		q, ok := pmap[p]
		if !ok {
			return false
		}
		// fmt.Printf("peek %s\n", []byte{p})

		if q == s[i] {
			st.Pop()
			continue
		}
		st.Push(s[i])
	}
	if st.Length() == 0 {
		return true
	} else {
		return false
	}
}

type stack struct {
	length int
	top    *node
}

type node struct {
	prev  *node
	value interface{}
}

func NewStack() *stack {
	return &stack{
		length: 0,
		top:    nil,
	}
}

func (s *stack) Peek() (v interface{}) {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *stack) Pop() (v interface{}) {
	if s.length == 0 {
		return nil
	}
	top := s.top
	s.top = top.prev
	s.length--
	return top.value
}
func (s *stack) Push(v interface{}) {
	top := &node{
		prev:  s.top,
		value: v,
	}
	s.top = top
	s.length++
}
func (s *stack) Length() int {
	return s.length
}

func testIsValid() {
	strs := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for i := range strs {
		rs := isValid(strs[i])
		fmt.Println(rs)
		fmt.Println()
	}
}
