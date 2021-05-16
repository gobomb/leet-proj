package leetcode

func addBinary(a string, b string) string {
	changei := func(a string) []int {
		ai := []int{}
		for i := 0; i < len(a); i++ {
			// for j:=len()
			ii := len(a) - 1 - i
			if a[ii] == '0' {
				ai = append(ai, 0)
			}
			if a[ii] == '1' {
				ai = append(ai, 1)
			}
		}
		return ai
	}
	ai := changei(a)
	bi := changei(b)
	ci := make([]int, max(len(ai), len(bi)))
	lci := len(ci)
	for i := 0; i < lci; i++ {
		var check int
		if i >= len(ai) {
			check = bi[i] + ci[i]
		} else if i >= len(bi) {
			check = ai[i] + ci[i]
		} else {
			check = ai[i] + bi[i] + ci[i]
		}
		switch check {
		case 0:
			continue
		case 1:
			ci[i] = 1
		case 2:
			if i+1 == len(ci) {
				ci = append(ci, 0)
			}
			ci[i+1] = ci[i+1] + 1
			ci[i] = 0
		case 3:
			if i+1 == len(ci) {
				ci = append(ci, 0)
			}
			ci[i+1] = ci[i+1] + 1
			ci[i] = 1
		}
	}
	// log.Printf("%v\n", ci)
	var rs string
	for i := 0; i < len(ci); i++ {
		if ci[i] == 0 {
			rs = "0" + rs
		}
		if ci[i] == 1 {
			rs = "1" + rs
		}
	}
	return rs
}
