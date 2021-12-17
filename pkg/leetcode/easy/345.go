package easy

/*
	345. Reverse Vowels of a String
*/

func reverseVowels(s string) string {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	// var b strings.Builder
	var bt []byte
	var st []int

	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := vowels[s[i]]; ok {
			st = append(st, i)
		}
	}

	ind := 0

	for i := range s {
		if _, ok := vowels[s[i]]; ok {
			// b.WriteByte(s[st[ind]])
			bt = append(bt, s[st[ind]])

			ind++
		} else {
			// b.WriteByte(s[i])
			bt = append(bt, s[i])
		}
	}

	return string(bt)
}
