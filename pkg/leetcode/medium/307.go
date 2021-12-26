package medium

import (
	"justest/pkg/leetcode/easy"
)

/*
	307. Range Sum Query - Mutable
*/

type NumArrayMutable interface {
	easy.NumArray
	Update(index int, val int)
}

type NumArray struct {
	nums []int
}

func ConstructorNumArray(nums []int) NumArrayMutable {
	return &NumArray{nums}
}

func (na *NumArray) Update(index int, val int) {
	na.nums[index] = val
}

func (na *NumArray) SumRange(left int, right int) int {
	var sum int

	for i := left; i <= right; i++ {
		sum += na.nums[i]
	}

	return sum
}

// Time Limit Exceeded
type NumArray1 struct {
	dp []int
}

func ConstructorNumArray1(nums []int) NumArrayMutable {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return &NumArray1{dp: append([]int{0}, nums...)}
}

// O(n)
func (na *NumArray1) Update(index int, val int) {
	old := na.dp[index+1] - na.dp[index]

	for i := index + 1; i < len(na.dp); i++ {
		na.dp[i] -= old
		na.dp[i] += val
	}
}

func (na *NumArray1) SumRange(left int, right int) int {
	return na.dp[right+1] - na.dp[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
