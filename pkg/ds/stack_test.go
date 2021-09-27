package ds

import (
	"testing"
)

func Test_Stack(t *testing.T) {
	s := NewStack()
	if s.Peek() == nil {
		t.Logf("peek is nil \n")
	}

	s.Push("1")
	t.Logf("len %d\n", s.Length())

	s.Push("2")
	t.Logf("pop %s\n", s.Pop())
	t.Logf("pop %s\n", s.Pop())
}

func Test_SliceStack(t *testing.T) {
	s := NewSliceStack()
	if s.Peek() == nil {
		t.Logf("peek is nil \n")
	}

	s.Push("1")
	t.Logf("len %d\n", s.Length())

	s.Push("2")
	t.Logf("pop %s\n", s.Pop())
	t.Logf("pop %s\n", s.Pop())
}
