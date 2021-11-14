package medium

import (
	"strings"
)

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

/*
	Runtime: 7 ms, faster than 28.67% of Go online submissions for Reverse Words in a String.
	Memory Usage: 2.9 MB, less than 89.88% of Go online submissions for Reverse Words in a String.
*/
func reverseWordsWithOutStl(s string) string {
	b := []byte(s)

	b = trimSpace(b)
	reverse(b)
	reverseWord(b)

	return string(b)
}

var SPACE byte = ' '

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseWord(b []byte) {
	start := 0

	for i := range b {
		if b[i] == SPACE && b[i-1] != SPACE {
			reverse(b[start:i])
			start = i + 1
		}
	}

	reverse(b[start:])
}

func trimSpace(b []byte) []byte {
	count := 0
	group := 0

	for i := range b {
		// log.Printf("%s %v %v %v", b, i, group, count)
		// b[i] may be changed
		temp := b[i]

		if b[i] != ' ' {
			b[group+count] = b[i]
			if i+1 != len(b) && b[i+1] == ' ' {
				b[group+count+1] = b[i+1]
			}
			count++
		}

		// log.Printf("%s %v %v %v %c", b, i, group, count, temp)

		if i+1 != len(b) && temp != ' ' && b[i+1] == ' ' {
			group++
		}
		// log.Printf("%s %v %v %v", b, i, group, count)
		// log.Printf("\n\n")
	}

	if b[len(b)-1] == ' ' {
		return b[:count+group-1]
	}

	return b[:count+group]
}
