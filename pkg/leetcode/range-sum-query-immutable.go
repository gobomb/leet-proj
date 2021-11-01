package leetcode

type NumArray struct {
	nums []int
}

func Constructor3(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (na *NumArray) SumRange(left int, right int) int {
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
