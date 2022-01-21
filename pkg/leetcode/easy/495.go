package easy

/*
	495. Teemo Attacking
*/

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	rs := 0

	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] < duration {
			rs += timeSeries[i] - timeSeries[i-1]
		} else {
			rs += duration
		}
	}

	rs += duration
	return rs
}
