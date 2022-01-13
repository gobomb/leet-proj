package medium

import (
	"math"
)

/*
	319. Bulb Switcher
*/

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
