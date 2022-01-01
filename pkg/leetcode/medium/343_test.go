package medium

import (
	"reflect"
	"testing"
)

func Test_integerBreak(t *testing.T) {
	type args struct {
		n int
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
					n: 2,
				},
				want: 1,
			},
			{
				name: "",
				args: args{
					n: 10,
				},
				want: 36,
			},
			{
				name: "",
				args: args{
					n: 5,
				},
				want: 6,
			},
			{
				name: "",
				args: args{
					n: 8,
				},
				want: 18,
			},
			{
				name: "",
				args: args{
					n: 20,
				},
				want: 1458,
			},
		}
	}

	toTest := []func(int) int{integerBreak, integerBreak1, integerBreak2}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
