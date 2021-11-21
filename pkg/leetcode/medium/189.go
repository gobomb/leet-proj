package medium

import "log"

/*
	189. Rotate Array
*/

func rotate(nums []int, k int) {
	if k == 0 || len(nums) == 0 {
		return
	}

	if len(nums) < k {
		k %= len(nums)
	}

	for i := 0; i < k; i++ {
		nums[i], nums[len(nums)-k+i] = nums[len(nums)-k+i], nums[i]
	}

	rotate(nums[k:], k)
}

func rotate1(nums []int, k int) {
	if k == 0 || len(nums) == 0 {
		return
	}

	if len(nums) < k {
		k %= len(nums)
	}

	reverseNums(nums, 0, len(nums)-1)
	log.Println(nums)
	reverseNums(nums, 0, k-1)
	log.Println(nums)

	reverseNums(nums, k, len(nums)-1)
	log.Println(nums)

}

func reverseNums(nums []int, st, ed int) {
	for st < ed {
		nums[st], nums[ed] = nums[ed], nums[st]
		st++
		ed--
	}
}
