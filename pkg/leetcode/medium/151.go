package medium

import "strings"

/*
	151. Reverse Words in a String
*/

/*
	Runtime: 4 ms, faster than 54.59% of Go online submissions for Reverse Words in a String.
	Memory Usage: 3.2 MB, less than 69.62% of Go online submissions for Reverse Words in a String.
*/
func reverseWords(s string) string {
	sf := strings.Fields(s)

	for i, j := 0, len(sf)-1; i < j; i, j = i+1, j-1 {
		sf[i], sf[j] = sf[j], sf[i]
	}

	sj := strings.Join(sf, " ")

	return sj
}
