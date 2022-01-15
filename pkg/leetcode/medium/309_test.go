package medium

import (
	"reflect"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
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
					prices: []int{1, 2, 3, 0, 2},
				},
				want: 3,
			},
			{
				name: "",
				args: args{
					prices: []int{1},
				},
				want: 0,
			},
		}

	}
	toTest := []func([]int) int{maxProfit, maxProfit1}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
