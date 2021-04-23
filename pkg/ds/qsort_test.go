package ds

import (
	"sort"
	"testing"
)

func genRandSlice(length, max int) ([]int, []int) {
	r := randSliceInt(max, length)
	sr := make([]int, length)
	copy(sr, r)
	sort.Ints(sr)
	return r, sr
}

func Test_quick(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}

	for i := 0; i < 5; i++ {
		r, sr := genRandSlice(15, 50)
		tests = append(tests, struct {
			name string
			args args
			want []int
		}{"", args{r}, sr})
		t.Logf("%v\n", tests[i])
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quick(tt.args.arr)
		})
	}
}
