package medium

/*
	134. Gas Station
*/

func canCompleteCircuit(gas []int, cost []int) int {
	tank, sum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		sum += gas[i] - cost[i]

		if sum < 0 {
			start = i + 1
			sum = 0
		}
	}

	if tank < 0 {
		return -1
	}

	return start
}
