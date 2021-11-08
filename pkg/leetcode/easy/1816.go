package easy

/* 
	1816. Truncate Sentence
*/

import (
	"strings"
)

func truncateSentence(s string, k int) string {
	return strings.Join(strings.Fields(s)[:k], " ")
}
