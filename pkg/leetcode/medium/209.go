package medium

/*
	209. Minimum Size Subarray Sum
*/

func minSubArrayLen(target int, nums []int) int {
	var res int = len(nums) + 1
	var sum int = nums[0]
	i, j := 0, 0

	for {
		if sum >= target {
			if res > j-i+1 {
				res = j - i + 1
			}

			sum -= nums[i]
			i++
			if i >= len(nums) {
				break
			}
		} else if sum < target {
			j++

			if j >= len(nums) {
				break
			}

			sum += nums[j]
		}

		if j < i {
			j = i
		}
	}

	if res > len(nums) {
		return 0
	}
	return res
}
