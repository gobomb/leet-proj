package easy

import (
	"bytes"
)

var morseCode = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	diff := map[string]struct{}{}
	for i := range words {
		buf := bytes.Buffer{}
		for j := 0; j < len(words[i]); j++ {
			index := words[i][j] - 'a'

			_, err := buf.WriteString(morseCode[index])
			if err != nil {
				panic(err)
			}
		}

		s := buf.String()
		_, ok := diff[s]
		if !ok {
			diff[s] = struct{}{}
		}
	}

	return len(diff)
}
