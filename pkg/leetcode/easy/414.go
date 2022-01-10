package easy

import (
	"math"
)

/*
	414. Third Maximum Number
*/

func thirdMax(nums []int) int {
	v1, v2, v3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if v > v1 {
			v3 = v2
			v2 = v1
			v1 = v
		} else if v < v1 && v > v2 {
			v3 = v2
			v2 = v
		} else if v < v2 && v >= v3 {
			v3 = v
		}
	}

	if v3 == math.MinInt64 {
		return v1
	}

	return v3
}
