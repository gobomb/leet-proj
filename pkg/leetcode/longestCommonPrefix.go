package leetcode

import "fmt"

func longestCommonPrefix(strs []string) string {
	// fmt.Println(strs)
	if len(strs) == 0 {
		return ""
	}
	var res string
	res = strs[0]
	var ind string
	for i := range strs {
		if len(strs[i]) == 0 {
			return ""
		}
		if len(res) > len(strs[i]) {
			var j int
			for j = range strs[i] {
				if strs[i][j] != res[j] {
					ind = strs[i][:j]
					break
				}
				ind = strs[i][:j+1]
			}
		} else {
			var j int
			for j = range res {
				if strs[i][j] != res[j] {
					ind = strs[i][:j]
					break
				}
				ind = strs[i][:j+1]
			}
		}
		res = ind
	}
	// fmt.Println(res)
	return res
}

func testlongestCommonPrefix() {
	strs1 := []string{"a"}
	str := longestCommonPrefix(strs1)
	fmt.Println(str == "a")
	strs1 = []string{}
	str = longestCommonPrefix(strs1)
	fmt.Println(str == "")

	strs1 = []string{"ab", "a"}
	str = longestCommonPrefix(strs1)
	fmt.Println(str == "a")
	// strs1 = []string{"", "flight"}
	// str = longestCommonPrefix(strs1)

	strs1 = []string{"acb", "b", "", ""}
	str = longestCommonPrefix(strs1)
	fmt.Println(str == "")
	strs1 = []string{"fli", "flight", ""}
	str = longestCommonPrefix(strs1)
	fmt.Println(str == "")

	fmt.Println(str == "")
	strs1 = []string{}
	str = longestCommonPrefix(strs1)
	fmt.Println(str == "")

	strs1 = []string{"fli", "flight"}
	str = longestCommonPrefix(strs1)
	fmt.Println(str == "fli")
	strs1 = []string{"flower", "flow", "flight"}
	str = longestCommonPrefix(strs1)
	fmt.Println(str == "fl")

	strs2 := []string{"dog", "racecar", "car"}
	str = longestCommonPrefix(strs2)
	fmt.Println(str == "")
}
