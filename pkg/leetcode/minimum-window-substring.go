package leetcode

// func minWindow(s string, t string) string {
// 	str := s + " "
// 	save := &[]int{}
// 	for i := 0; i < len(s); i = getNextI(i, save) {
// 		s2 := t
// 		for j := i + 1; j <= len(s); j++ {
// 			if s2, del := checkAndDel(s[i], s2); del {
// 				*save = append(*save, i)
// 				for ii := i + 1; ii < j-1; ii++ {
// 					var d bool
// 					s2, d = checkAndDel(s[ii], s2)
// 					if d {
// 					}
// 				}
// 				s2, del = checkAndDel(s[j-1], s2)
// 				if del && len(s2) == 0 {
// 					if len(s[i:j]) < len(str) {
// 						str = s[i:j]
// 					}
// 				}
// 				continue
// 			}
// 		}
// 	}
// 	return ""
// }

// func checkAndDel(b byte, s string) (string, bool) {
// 	for i := range s {
// 		if b == s[i] {
// 			s = s[:i] + s[i+1:]
// 			return s, true
// 		}
// 	}
// 	return s, false
// }

// Time Limit Exceeded
func minWindow(s string, t string) string {
	// save := make(map[byte]int)

	tt := &T{
		l:   len(t),
		n:   0,
		str: t,
	}
	tt.reset()
	rs := s + " "
	str := ""
	queue := []int{}
	head := 0
	// var sFreq [256]int

	for i := -1; i < len(s); {
		i++
		if len(queue) != 0 {
			if head >= len(queue) {
				break
			}
			i = queue[head]
			// log.Println(queue, i)
			head++
		}
		if i >= len(s) {
			break
		}
		tt.reset()
		str = ""
		if !tt.checkAndUpdate(s[i]) {
			continue
		}

		str = str + string(s[i])
		if tt.isEnd() {
			if len(str) < len(rs) {
				rs = str

			}

			continue
		}

		for j := i + 1; j < len(s); j++ {
			if tt.checkAndUpdate(s[j]) {
				// log.Println(j)

				queue = append(queue, j)
			}
			str = str + string(s[j])
			// log.Println(str, i, j)
			if tt.isEnd() {
				if len(str) < len(rs) {
					rs = str
				}
				break
			}
		}

	}
	if rs == s+" " {
		return ""
	}
	return rs
}

type T struct {
	save map[byte]int
	n, l int
	str  string
}

func (t *T) reset() {
	t.save = make(map[byte]int)

	for i := 0; i < t.l; i++ {
		if v, ok := t.save[t.str[i]]; !ok {
			t.save[t.str[i]] = 1
		} else {
			t.save[t.str[i]] = v + 1
		}

	}
	t.n = 0
}

func (t *T) isEnd() bool {
	// log.Println(t.save, t.l, t.n)

	return t.n == t.l
}

func (t *T) checkAndUpdate(b byte) bool {
	if v, ok := t.save[b]; !ok {
		return false
	} else {
		if v == 0 {
			return true
		}
		t.save[b] = t.save[b] - 1
		t.n++
		return true
	}
}
