package leetcode

func fullJustify(words []string, maxWidth int) []string {
	strs := []string{}
	lenc := 0
	rs := []string{}
	for i := range words {
		lenc += len(words[i])
		strs = append(strs, words[i])
		if i+1 == len(words) {
			cachestr := ""
			for j := 0; j < len(strs); j++ {
				cachestr += strs[j]
				if len(cachestr) < maxWidth {
					cachestr += " "
				}
			}
			for len(cachestr) < maxWidth {
				cachestr += " "
			}

			rs = append(rs, cachestr)
			strs = []string{}
			break
		}
		if len(strs)+lenc+len(words[i+1]) > maxWidth {
			newc := ""
			if len(strs) == 1 {
				newc += strs[0]
				for len(newc) < maxWidth {
					newc += " "
				}
			} else {
				for j := range strs {
					newc += strs[j]
					if j == len(strs)-1 {
						break
					}
					for k := 0; k < (maxWidth-lenc)/(len(strs)-1); k++ {
						newc += " "
					}
					if j < (maxWidth-lenc)%(len(strs)-1) {
						newc += " "
					}
				}
			}
			rs = append(rs, newc)
			strs = []string{}
			lenc = 0
		}
	}

	// for i := 0; i < len(rs); i++ {
	// 	log.Printf("%v\n", len(rs[i]))
	// }
	return rs
}
