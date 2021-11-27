package medium

import (
	"fmt"
	"testing"
)

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
				t:    0,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
				t:    2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 5, 9, 1, 5, 9},
				k:    2,
				t:    3,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				nums: []int{1999, 5333, 922, 119, 25, 39, 8, 100, 300, -20, -8},
				k:    5,
				t:    9,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 21, 31, 41, 51, 61},
				k:    2,
				t:    3,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2},
				k:    0,
				t:    1,
			},
			want: false,
		},
	}

	toTest := []func([]int, int, int) bool{containsNearbyAlmostDuplicate, containsNearbyAlmostDuplicate1}

	for _, f1 := range toTest {
		f := func(a args) bool {
			fmt.Printf("args %+v\n", a)
			return f1(a.nums, a.k, a.t)
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args); got != tt.want {
					t.Errorf("%v() = %+v, want %+v", funcName(f), got, tt.want)
				}
			})
		}
	}
}
