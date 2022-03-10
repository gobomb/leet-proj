package dynamicprogram

import (
	"reflect"
	"testing"
)

func Test_wiggleMaxLength(t *testing.T) {
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
				name: "",
				args: args{
					nums: []int{1, 7, 4, 9, 2, 5},
				},
				want: 6,
			},
			{
				name: "",
				args: args{
					nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
				},
				want: 7,
			},
			{
				name: "",
				args: args{
					nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				},
				want: 2,
			},
			{
				name: "",
				args: args{
					nums: []int{0, 0},
				},
				want: 1,
			},
			{
				name: "",
				args: args{
					nums: []int{0, 1, 0},
				},
				want: 3,
			},

			{
				name: "",
				args: args{
					nums: []int{0, 0, 1},
				},
				want: 2,
			},
			{
				name: "",
				args: args{
					nums: []int{0, 1, 1},
				},
				want: 2,
			},
			{
				name: "",
				args: args{
					nums: []int{1, 0, 0},
				},
				want: 2,
			},
		}
	}
	toTest := []func([]int) int{wiggleMaxLength, wiggleMaxLength2}

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
