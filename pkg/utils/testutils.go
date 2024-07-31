package utils

import "github.com/google/go-cmp/cmp"

func ShouldEqualDiff(actual interface{}, expected ...interface{}) string {
	if len(expected) == 0 {
		return "expected should not be empty"
	}

	return cmp.Diff(expected[0], actual)
}
