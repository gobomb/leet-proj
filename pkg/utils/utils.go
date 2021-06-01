package utils

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	return max(a, b)
}
func Min(a, b int) int {
	return min(a, b)
}
