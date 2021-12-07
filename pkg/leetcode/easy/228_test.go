package easy

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := func() []struct {
		name string
		args args
		want []string
	} {
		return []struct {
			name string
			args args
			want []string
		}{
			{
				name: "",
				args: args{
					nums: []int{0, 1, 2, 4, 5, 7},
				},
				want: []string{"0->2", "4->5", "7"},
			},
			{
				name: "",
				args: args{
					nums: []int{0, 2, 3, 4, 6, 8, 9},
				},
				want: []string{"0", "2->4", "6", "8->9"},
			},
			{
				name: "",
				args: args{
					nums: []int{},
				},
				want: []string{},
			},
			{
				name: "",
				args: args{
					nums: []int{-1},
				},
				want: []string{"-1"},
			},
			{
				name: "",
				args: args{
					nums: []int{0},
				},
				want: []string{"0"},
			},
		}
	}

	toTest := []func([]int) []string{summaryRanges, summaryRanges2}

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
