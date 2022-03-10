package dynamicprogram

/*
	376. Wiggle Subsequence
*/

func wiggleMaxLength2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	p, n := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			p = n + 1
		} else if nums[i] < nums[i-1] {
			n = p + 1
		}
	}

	return max(p, n)
}

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	dp := make([]int, len(nums))
	last := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		last[i] = nums[i] - nums[i-1]
	}

	dp[0] = 1
	if nums[1] != nums[0] {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	rs := 0
	for i := 2; i < len(dp); i++ {
		if nums[i]-nums[0] == 0 {
			dp[i] = 1
		} else {
			dp[i] = 2
		}

		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]

			if diff != 0 && last[j] != 0 && (diff^last[j]) < 0 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if dp[i] > rs {
			rs = dp[i]
		}
	}

	// log.Println(last, dp)

	return rs
}
