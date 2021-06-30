package leetcode

import "log"

func candy(ratings []int) int {
	rs := 0
	candis := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candis[i] += candis[i-1] + 1
		}
	}
	log.Printf("%+v\n", candis)

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candis[i] <= candis[i+1] {
			candis[i] = candis[i+1] + 1
		}
	}
	log.Printf("%+v\n", candis)
	for _, c := range candis {
		rs += c + 1
	}
	return rs
}
