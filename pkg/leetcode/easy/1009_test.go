package easy

import (
	"reflect"
	"testing"
)

func Test_bitwiseComplement(t *testing.T) {
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
					n: 5,
				},
				want: 2,
			},
			{
				name: "",
				args: args{
					n: 7,
				},
				want: 0,
			},
			{
				name: "",
				args: args{
					n: 10,
				},
				want: 5,
			},
			{
				name: "",
				args: args{
					n: 0,
				},
				want: 1,
			},
		}
	}
	toTest := []func(int) int{bitwiseComplement1, bitwiseComplement}

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
