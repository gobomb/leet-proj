package leetcode

func climbStairs(n int) int {
	saved := make([]int, n+1)
	return backtrackClimbStatirs(n, saved)
}

func backtrackClimbStatirs(n int, saved []int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n >= 2 {
		if saved[n-1] == 0 {
			saved[n-1] = backtrackClimbStatirs(n-1, saved) + saved[n-1]
		}
		if saved[n-2] == 0 {
			saved[n-2] = backtrackClimbStatirs(n-2, saved) + saved[n-2]
		}
	}
	return saved[n-1] + saved[n-2]
}

func dpClimbStairs(n int) int {
	a, b := 1, 1
	for i := 2; i < n+1; i++ {
		a, b = b, a+b
	}
	return b
}
