package leetcode

/*

	771. Jewels and Stones

*/

func numJewelsInStones(jewels string, stones string) int {
	jm := make(map[byte]struct{})
	rs := 0
	for i := range jewels {
		jm[jewels[i]] = struct{}{}
	}
	for i := range stones {
		if _, ok := jm[stones[i]]; ok {
			rs++
		}
	}
	return rs
}
