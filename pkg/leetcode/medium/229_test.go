package medium

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := func() []struct {
		name string
		args args
		want []int
	} {
		return []struct {
			name string
			args args
			want []int
		}{
			{
				name: "",
				args: args{
					nums: []int{3, 2, 3},
				},
				want: []int{3},
			},
			{
				name: "",
				args: args{
					nums: []int{1},
				},
				want: []int{1},
			},
			{
				name: "",
				args: args{
					nums: []int{1, 2},
				},
				want: []int{1, 2},
			},
			{
				name: "",
				args: args{
					nums: []int{1, 2, 3},
				},
				want: []int{},
			},
		}
	}

	toTest := []func([]int) []int{majorityElement, majorityElement1}

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
