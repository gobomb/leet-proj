package review

import (
	"sort"
)

// 15. 3Sum
// 27min

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	rs := [][]int{}

	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1

		// 需要和上一次枚举的数不相同
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				rs = append(rs, []int{nums[i], nums[j], nums[k]})
				
				j++

				if k <= j {
					break
				}

				for nums[j] == nums[j-1] && k > j {
					j++
				}
			}

			if nums[i]+nums[j]+nums[k] > 0 {
				k--

				if k <= j {
					break
				}

				for nums[k] == nums[k+1] && k > j {
					k--
				}
			}

			if nums[i]+nums[j]+nums[k] < 0 {
				j++

				if k <= j {
					break
				}

				for nums[j] == nums[j-1] && k > j {
					j++
				}
			}
		}
	}

	return rs
}
