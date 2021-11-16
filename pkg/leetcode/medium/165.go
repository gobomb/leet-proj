package medium

import (
	"strconv"
	"strings"
)

/*
	165. Compare Version Numbers
*/

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	n1 := []int{}
	n2 := []int{}

	for _, v := range v1 {
		n, _ := strconv.Atoi(v)
		n1 = append(n1, n)
	}

	for _, v := range v2 {
		n, _ := strconv.Atoi(v)
		n2 = append(n2, n)
	}

	if len(n1) < len(n2) {
		d := len(n2) - len(n1)

		for i := 0; i < d; i++ {
			n1 = append(n1, 0)
		}
	} else {
		d := len(n1) - len(n2)

		for i := 0; i < d; i++ {
			n2 = append(n2, 0)
		}
	}

	for i := 0; i < len(n1); i++ {
		if n1[i] < n2[i] {
			return -1
		}

		if n1[i] > n2[i] {
			return 1
		}
	}

	return 0
}
