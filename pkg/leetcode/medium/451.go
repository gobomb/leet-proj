package medium

/*
	451. Sort Characters By Frequency
*/

func frequencySort(s string) string {
	rs := ""

	freq := make(map[byte]int)

	bucket := make([]string, len(s)+1)

	for i := range s {
		freq[s[i]]++
	}

	for k, v := range freq {
		vv := v
		for vv != 0 {
			bucket[v] = bucket[v] + string(k)
			vv--
		}
	}

	for i := len(s); i > 0; i-- {
		if bucket[i] != "" {
			rs = rs + bucket[i]
		}
	}

	return rs
}
