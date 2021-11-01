package leetcode

import (
	"log"
	"math"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var canReach, needChoose, step int
	for i, n := range nums {
		if i > canReach {
			return -1
		}
		if i+n > canReach {
			canReach = i + n

			if i+n >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}

var dept *int

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var canReach, needChoose, step int
	for i, n := range nums {
		if i > canReach {
			return false
		}
		if i+n > canReach {
			canReach = i + n

			if i+n >= len(nums)-1 {
				return true
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return false
}

var memo []int

func jump2(nums []int) int {
	memo = make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = math.MaxInt64
	}
	dept = new(int)
	*dept = -1

	tryJump(&nums, 0)
	log.Printf("%v\n", memo)

	return memo[0]
}

// []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
// [0 2 -1 -1 -1 -1 -2 2 1 -1 -1 0]

func tryJump(nums *[]int, i int) (b bool, maxd int) {
	*dept++
	defer func() {

		*dept--
	}()
	if i >= len(*nums) {
		maxd = *dept

		return false, maxd
	}
	if i == len(*nums)-1 {
		memo[i] = 0
		maxd = *dept

		return true, maxd
	}

	var found bool

	if memo[i] == math.MaxInt64 {
		for j := 1; j <= (*nums)[i]; j++ {
			ok, maxd1 := tryJump(nums, i+j)
			if ok {
				if memo[i] > (maxd1 - *dept) {
					memo[i] = maxd1 - *dept
					maxd = maxd1
				}
				found = true
			}
		}

		if !found {
			memo[i] = -1
		}
	} else {
		if memo[i] < 0 {
			found = false
		} else {
			maxd = memo[i] + *dept
			found = true
		}
	}

	return found, maxd
}
