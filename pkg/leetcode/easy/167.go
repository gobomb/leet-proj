package easy

/*
	167. Two Sum II - Input Array Is Sorted
*/

// 132 ms	3.2 MB
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{}
}

// 	4 ms	3 MB
func twoSumBiSe(numbers []int, target int) []int {
	for i := range numbers {
		target -= numbers[i]
		lo, hi, mid := i+1, len(numbers)-1, 0

		for lo <= hi {
			mid = lo + (hi-lo)>>1

			switch {
			case target > numbers[mid]:
				lo = mid + 1
			case target < numbers[mid]:
				hi = mid - 1
			case target == numbers[mid]:
				return []int{i + 1, mid + 1}
			}
		}

		target += numbers[i]
	}

	return []int{}
}
