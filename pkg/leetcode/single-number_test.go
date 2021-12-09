package leetcode

import (
	"reflect"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
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
				name: "1",
				args: args{
					nums: []int{2, 2, 1},
				},
				want: 1,
			},
			{
				name: "2",
				args: args{
					nums: []int{4, 1, 2, 1, 2},
				},
				want: 4,
			},
			{
				name: "3",
				args: args{
					nums: []int{1},
				},
				want: 1,
			},
		}
	}
	toTest := []func([]int) int{singleNumber, sigleNumberXor, singleNumberXor2}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
