package medium

import (
	"math/rand"
)

/*
	398. Random Pick Index
*/

type Solution struct {
	m map[int][]int
}

func ConstructorSolution(nums []int) Solution {
	s := Solution{
		make(map[int][]int),
	}

	for i, n := range nums {
		s.m[n] = append(s.m[n], i)
	}

	return s
}

func (s *Solution) Pick(target int) int {
	sl := s.m[target]
	r := rand.Intn(len(sl))
	return sl[r]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
