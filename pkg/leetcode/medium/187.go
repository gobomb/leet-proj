package medium

/*
	187. Repeated DNA Sequences
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
