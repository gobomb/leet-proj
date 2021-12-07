package hard

import (
	"reflect"
	"testing"
)

func Test_calculate(t *testing.T) {
	type args struct {
		s string
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
					s: "1 - 1",
				},
				want: 0,
			},
			{
				name: "",
				args: args{
					s: "1 + 1",
				},
				want: 2,
			},

			{
				name: "",
				args: args{
					s: "1 + (1+2)+3",
				},
				want: 7,
			},
			{
				name: "",
				args: args{
					s: "(1+(4+5+2)-3)+(6+8)",
				},
				want: 23,
			},
			{
				name: "",
				args: args{
					s: " 2-1 + 2 ",
				},
				want: 3,
			},
			{
				name: "",
				args: args{
					s: " 2+(-1) + 2 ",
				},
				want: 3,
			},
			{
				name: "",
				args: args{
					s: " -2+(-1) + 2 ",
				},
				want: -1,
			},
		}
	}
	toTest := []func(string) int{calculate, calculate1}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
