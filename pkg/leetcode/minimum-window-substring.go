package leetcode

func minWindow(s string, t string) string {
	var i, j int
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	isNeed := len(t)
	var rs string
	for i != len(s) {
		if _, ok := need[s[i]]; ok {
			break
		}
		i++
	}
	for j = i; j != len(s); j++ {
		if _, ok := need[s[j]]; ok {
			need[s[j]]--
			if need[s[j]] >= 0 {
				isNeed--
			}
			if isNeed == 0 {
				if len(rs) == 0 || len(rs) > j-i+1 {
					rs = s[i : j+1]
				}
				if _, ok := need[s[i]]; ok {
					need[s[i]]++
					if need[s[i]] > 0 {
						isNeed++
					}
					i++
				}
				for ; i < j; i++ {
					if _, ok := need[s[i]]; ok {
						if need[s[i]] >= 0 {
							if isNeed == 0 && (len(rs) == 0 || len(rs) > j-i+1) {
								rs = s[i : j+1]
							}
							break
						} else {
							need[s[i]]++
							if need[s[i]] > 0 {
								isNeed++
							}
						}
					}
				}
			}
		}
	}
	return rs
}
