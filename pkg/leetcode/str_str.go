package leetcode

import "fmt"

// runtime 400 ms
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	var r int = -1
	for i := range haystack {
		k := i
		for j := 0; j < len(needle); j++ {
			if k >= len(haystack) {
				r = -1
				break
			}
			if haystack[k] == needle[j] {
				if j == 0 {
					r = k
				}
				// fmt.Printf("\nk:%d\n", k)
				if j == len(needle)-1 {
					return r
				}
				k++
				continue
			}
			if haystack[k] != needle[j] {
				r = -1
				break
			}
		}
	}
	return r
}

// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0028.Implement-strStr/
// runtime 0 ms
func strStr1(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

func testStrStr() {
	tests := []struct {
		a string
		b string
	}{
		{"asdcsds", "csd"},
		{"", ""},
		{"aaaaa", "bba"},
		{"mississippi",
			"issipi"},
		{"aaaaaaaaaaaaaaaaaaa",
			"aaaaaaaaaaa",
		},
	}

	for _, test := range tests {
		r := strStr1(test.a, test.b)
		fmt.Println(r)
	}
}
