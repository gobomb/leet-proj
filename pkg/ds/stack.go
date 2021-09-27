package ds

import (
	"sync"
)

type Stack interface {
	Pop() interface{}
	Push(interface{})
	Length() int
	Peek() interface{}
}

type stack struct {
	length int
	top    *node
	lock   *sync.RWMutex
}

type node struct {
	prev  *node
	value interface{}
}

func NewStack() Stack {
	return &stack{
		length: 0,
		top:    nil,
		lock:   &sync.RWMutex{},
	}
}

func (s *stack) Peek() (v interface{}) {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *stack) Pop() (v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.length == 0 {
		return nil
	}
	top := s.top
	s.top = top.prev
	s.length--
	return top.value
}
func (s *stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
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
