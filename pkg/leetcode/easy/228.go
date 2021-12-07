package easy

import (
	"strconv"
)

/*
	228. Summary Ranges
*/

func summaryRanges(nums []int) []string {
	if len(nums) < 1 {
		return []string{}
	}
	res := []string{}
	last := nums[0]
	temp := []int{last, last - 2}

	for i := 1; i < len(nums); i++ {
		if nums[i]-last == 1 {
			temp[1] = nums[i]
			last = nums[i]
		} else if temp[0] < temp[1] {
			s := strconv.FormatInt(int64(temp[0]), 10) + "->" + strconv.FormatInt(int64(temp[1]), 10)
			res = append(res, s)
			temp[0] = nums[i]
			last = nums[i]
		} else {
			s := strconv.FormatInt(int64(temp[0]), 10)
			res = append(res, s)
			temp[0] = nums[i]
			last = nums[i]
		}
	}

	if temp[0] < temp[1] {
		s := strconv.FormatInt(int64(temp[0]), 10) + "->" + strconv.FormatInt(int64(temp[1]), 10)
		res = append(res, s)
	} else {
		s := strconv.FormatInt(int64(temp[0]), 10)
		res = append(res, s)
	}

	return res
}

// https://leetcode.com/problems/summary-ranges/discuss/402364/Golang-%3A-faster-than-100.00-of-Go-online-submissions
func summaryRanges2(nums []int) []string {
	if len(nums) < 1 {
		return []string{}
	}

	res := []string{}
	head := 0

	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}

		if head == i {
			res = append(res, strconv.FormatInt(int64(nums[i]), 10))
		} else {
			res = append(res, strconv.FormatInt(int64(nums[head]), 10)+"->"+strconv.FormatInt(int64(nums[i]), 10))
		}

		head = i + 1
	}

	return res
}
