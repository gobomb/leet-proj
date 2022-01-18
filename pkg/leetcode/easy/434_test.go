package easy

import (
	"reflect"
	"testing"
)

func Test_countSegments(t *testing.T) {
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
					s: "Hello, my name is John",
				},
				want: 5,
			},
			{
				name: "",
				args: args{
					s: "Hello",
				},
				want: 1,
			},
			{
				name: "",
				args: args{
					s: " Hello",
				},
				want: 1,
			},
			{
				name: "",
				args: args{
					s: "  Hello",
				},
				want: 1,
			},
			{
				name: "",
				args: args{
					s: "He  llo",
				},
				want: 2,
			},
		}
	}
	toTest := []func(string) int{countSegments, countSegments1}

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
