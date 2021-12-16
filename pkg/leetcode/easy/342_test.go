package easy

import (
	"reflect"
	"testing"
)

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}

	tests := func() []struct {
		name string
		args args
		want bool
	} {
		return []struct {
			name string
			args args
			want bool
		}{
			{
				name: "",
				args: args{
					n: 16,
				},
				want: true,
			},
			{
				name: "",
				args: args{
					n: 5,
				},
				want: false,
			},
			{
				name: "",
				args: args{
					n: 1,
				},
				want: true,
			},
			{
				name: "",
				args: args{
					n: 2,
				},
				want: false,
			},
			{
				name: "",
				args: args{
					n: 0,
				},
				want: false,
			},
		}
	}

	toTest := []func(int) bool{isPowerOfFour, isPowerOfFour1, isPowerOfFour2}

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
