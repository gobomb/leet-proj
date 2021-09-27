package ds

type slice_stack struct {
	s []interface{}
}

func NewSliceStack() Stack {
	return &slice_stack{
		s: []interface{}{},
	}
}

func (s *slice_stack) Peek() (v interface{}) {
	if len(s.s) == 0 {
		return nil
	}
	return s.s[len(s.s)-1]
}

func (s *slice_stack) Pop() (v interface{}) {
	if len(s.s) == 0 {
		return nil
	}
	top := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return top
}
func (s *slice_stack) Push(v interface{}) {
	s.s = append(s.s, v)
}

func (s *slice_stack) Length() int {
	return len(s.s)
}
