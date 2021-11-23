package medium

/*
	187. Repeated DNA Sequences
*/

/*
	16 ms	9.7 MB
*/
func findRepeatedDnaSequences(s string) []string {
	ss := make(map[string]int)
	res := []string{}

	for i := 0; i+9 < len(s); i++ {
		temp := s[i : i+10]

		if _, ok := ss[temp]; ok {
			ss[temp]++
			if ss[temp] == 2 {
				res = append(res, temp)
			}
		} else {
			ss[temp]++
		}
	}

	return res
}

var numMap = map[byte]uint32{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

var upperMask uint32 = 0xFFF00000

/*
	Runtime: 12 ms, faster than 85.88% of Go online submissions for Repeated DNA Sequences.
	Memory Usage: 5.4 MB, less than 90.59% of Go online submissions for Repeated DNA Sequences.

	https://leetcode.com/problems/repeated-dna-sequences/discuss/318415/Golang-solution-rolling-hash
*/
func findRepeatedDnaSequences1(s string) []string {
	sMap := map[uint32]bool{}
	res := []string{}

	var v uint32
	for i := 0; i < len(s); i++ {
		num := numMap[s[i]]
		v <<= 2
		v |= num
		v = v &^ upperMask

		if i < 9 {
			continue
		}

		if count, ok := sMap[v]; !ok {
			sMap[v] = false
		} else if !count {
			sMap[v] = true
			res = append(res, s[i-9:i+1])
		}
	}
	return res
}
