package medium

import (
	"bytes"
	"sort"
	"strconv"
)

/*
	179. Largest Number
*/

// 剑指offer 33
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		ii := strconv.FormatInt(int64(nums[i]), 10)
		jj := strconv.FormatInt(int64(nums[j]), 10)

		return ii+jj > jj+ii
	})

	var res bytes.Buffer

	if nums[0] == 0 {
		return "0"
	}

	for _, n := range nums {
		res.WriteString(strconv.FormatInt(int64(n), 10))
	}

	return res.String()
}
