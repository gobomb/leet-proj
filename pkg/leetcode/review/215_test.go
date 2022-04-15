package review

import (
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := func() []struct {
		name string
		args args
		want int
	} {
		return []struct {
			name string
			args args
			want int
		}{
			{
				name: "",
				args: args{
					nums: []int{3, 2, 1, 5, 6, 4},
					k:    2,
				},
				want: 5,
			},
			{
				name: "",
				args: args{
					nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
					k:    4,
				},
				want: 4,
			},
			{
				name: "",
				args: args{
					nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
					k:    14,
				},
				want: -1,
			},
		}
	}

	toTest := []func([]int, int) int{findKthLargestQuickSelect, findKthLargestHeap}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums, tt.args.k); got != tt.want {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
