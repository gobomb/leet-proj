package easy

/*
	303. Range Sum Query - Immutable
*/

type NumArray interface {
	SumRange(left int, right int) int
}

/*
	136 ms	10.4 MB
*/
type NumArray0 struct {
	nums []int
}

func Constructor3(nums []int) NumArray {
	return &NumArray0{nums: nums}
}

func (na *NumArray0) SumRange(left int, right int) int {
	var sum int
	for i := left; i <= right; i++ {
		sum += na.nums[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

/*
	Runtime: 28 ms, faster than 90.00% of Go online submissions for Range Sum Query - Immutable.
	Memory Usage: 9.6 MB, less than 65.00% of Go online submissions for Range Sum Query - Immutable.
*/
type NumArray1 struct {
	dp []int
}

func Constructor31(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return &NumArray1{dp: append([]int{0}, nums...)}
}

func (na *NumArray1) SumRange(left int, right int) int {
	return na.dp[right+1] - na.dp[left]
}
